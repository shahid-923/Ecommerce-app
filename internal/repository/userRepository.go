package repository

import (
	"ecommerce-app/internal/domain"
    "log"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
  CreateUser(user domain.User) (domain.User, error)
  FindUser(email string) (domain.User, error)
  FindUserByID(id uint) (domain.User, error)
  UpdateUser(id uint, user domain.User) (domain.User, error) 

}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
   db *gorm.DB
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil{
		log.Printf("error creating user: %v\n", err)
		return domain.User{}, errors.New("failed to create user")
	}

   return user, nil
}

func (r userRepository) FindUser(email string)(domain.User, error){
	return domain.User{}, nil
}

func (r userRepository) FindUserByID(id uint )(domain.User, error){
	return domain.User{}, nil
}

func (r userRepository) UpdateUser(id uint , user domain.User)(domain.User, error){
	return domain.User{}, nil
}
