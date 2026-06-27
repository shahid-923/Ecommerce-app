package service

import(
	
	"ecommerce-app/internal/helper"
	"ecommerce-app/internal/repository"

	"ecommerce-app/config"
)

type CatalogueService struct {
	Repo   repository.CatalogueRepository
	Auth   helper.Auth
	Config config.AppConfig
}