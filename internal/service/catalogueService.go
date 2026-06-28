package service

import (
	"ecommerce-app/internal/domain"
	"ecommerce-app/internal/dto"
	"ecommerce-app/internal/helper"
	"ecommerce-app/internal/repository"
	"errors"

	"ecommerce-app/config"
)

type CatalogueService struct {
	Repo   repository.CatalogueRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func(s CatalogueService) CreateCategory(input dto.CreateCategoryRequest) error{
	
	err := s.Repo.CreateCategory(&domain.Category{
	 Name: input.Name,
	 ImageUrl: input.ImageUrl,
	 DisplayOrder: input.DisplayOrder,

	})
	
	return err
}

func (s CatalogueService) GetCategories() ([]*domain.Category, error) {

	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, errors.New("categories dont exist")
	}

	return categories, nil
}

func (s CatalogueService) GetCategory(id int) (*domain.Category, error) {

	category, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}

	return category, nil
}
 
func (s CatalogueService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {

	existCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exists")
	}

	if input.Name != "" {
		existCat.Name = input.Name
	}

	if input.ParentId > 0 {
		existCat.ParentId = input.ParentId
	}

	if input.ImageUrl != "" {
		existCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		existCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(existCat)
	if err != nil {
		return nil, err
	}

	return updatedCat, nil
}

func(s CatalogueService) DeleteCategory(id int) error{
	err := s.Repo.DeleteCategory(id)

	if err != nil{
		return errors.New("category does not exist for deleting")
	}
	return nil
}


