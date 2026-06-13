package service

import (
	"ecommerce-app/internal/domain"
	"ecommerce-app/internal/dto"
	"ecommerce-app/internal/helper"
	"ecommerce-app/internal/repository"
	"errors"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s *UserService) Signup(input dto.UserSignup) (domain.User, string, error) {

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
// func (s UserService) GetProfile(id uint) (*domain.User, error) {

// 	user, err := s.Repo.FindUserById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
 func (s *UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (s *UserService) GetUserByID(id uint) (domain.User, error) {
	return s.Repo.FindUserByID(id)
}
func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) (string, error) {
	return "", nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}
func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}
func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}
func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}
func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}
func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}
func (s UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	return nil, nil
}
