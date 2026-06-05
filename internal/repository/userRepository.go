package repository

import (
	"ecommerce-app/internal/domain"
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		log.Printf("error finding user: %v\n", err)
		return domain.User{}, errors.New("user doesn't exist")
	}
	return user, nil
}

func (r userRepository) FindUserByID(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("error finding user: %v\n", err)
		return domain.User{}, errors.New("user doesn't exist")
	}
	return user, nil
}

func (r userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {
	var existingUser domain.User

	err := r.db.Model(&existingUser).Clauses(clause.Returning{}).Where("id = ?", id).Updates(user).Error

	if err != nil {
		log.Printf("error updating user: %v\n", err)
		return domain.User{}, errors.New("failed to update user")
	}

	return existingUser, nil
}
