package service

import (
	"ecommerce-app/internal/domain"
	"ecommerce-app/internal/dto"
	"ecommerce-app/internal/helper"
	"ecommerce-app/internal/repository"
	"errors"
	"time"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s *UserService) Register(input dto.UserSignup) (domain.User, string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return domain.User{}, "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		return domain.User{}, "", err
	}

	token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
	if err != nil {
		return domain.User{}, "", err
	}

	return user, token, nil
}

func (s *UserService) Login(email string, password string) (string, error) {
	user, err := s.Repo.FindUser(email)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s *UserService) GetProfile(id uint) (domain.User, error) {
	return s.Repo.FindUserByID(id)
}

func (s *UserService) GetUserByID(id uint) (domain.User, error) {
	return s.Repo.FindUserByID(id)
}

func (s *UserService) isVerifiedUser(id uint) (bool) {

	currentUser, err := s.Repo.FindUserByID(id)
	return err == nil && currentUser.Verified     // if true else return false
    
}
func (s *UserService) GetVerificationCode(e domain.User) (int, error) {
	
	// if user already verified
	if s.isVerifiedUser(e.ID) {
    return 0, errors.New("user already verified")
    }

	// generate verification code 
    code, err := s.Auth.GenerateCode()
	if err != nil{
		return 0, err
	}

	// update user 
	user := domain.User{
	  Expiry: time.Now().Add(30 * time.Minute),	
	  Code: code,
	}
    
	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil{
	  return 0, errors.New("failed to update user")	
	}
    // send the sms to user phone number


	// return the verification code
	return code, nil	
}

func (s *UserService) VerifyCode(id uint, code int) error {
	if s.isVerifiedUser(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserByID(id)
	if err != nil {
		return err
	}

	log.Printf("DB code: %d, Input code: %d", user.Code, code)
	log.Printf("Expiry: %v, Now: %v", user.Expiry, time.Now())

	if user.Code != code {
		return errors.New("invalid verification code")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updatedUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updatedUser)
	if err != nil {
		return errors.New("unable to update user")
	}

	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s *UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s *UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s *UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s *UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	return nil, nil
}