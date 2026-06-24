package repository

import (
	"errors"
	"log"

	"ecommerce-app/internal/domain"

	"gorm.io/gorm"
)

type CatalogueRepository interface {
	CreateCategory(c *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryById(id int) (*domain.Category, error)
	EditCategory(c *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
}

type catalogueRepository struct {
	db *gorm.DB
}

func NewCatalogueRepository(db *gorm.DB) CatalogueRepository {
	return &catalogueRepository{
		db: db,
	}
}

func (c catalogueRepository) CreateCategory(e *domain.Category) error {

	err := c.db.Create(e).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("creating category failed")
	}

	return nil
}

func (c catalogueRepository) FindCategories() ([]*domain.Category, error) {

	var categories []*domain.Category

	err := c.db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c catalogueRepository) FindCategoryById(id int) (*domain.Category, error) {

	var category domain.Category

	err := c.db.First(&category, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("category does not exist")
	}

	return &category, nil
}

func (c catalogueRepository) EditCategory(e *domain.Category) (*domain.Category, error) {

	err := c.db.Save(e).Error 

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("failed to update category")
	}

	return e, nil
}

func (c catalogueRepository) DeleteCategory(id int) error {

	err := c.db.Delete(&domain.Category{}, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("failed to delete category")
	}

	return nil
}