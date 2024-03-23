package src

import (
	"backend-test/config"
	"backend-test/src/repositories"
	"backend-test/src/services"
)

type Repository struct {
	ProductRepo *repositories.ProductRepository
}

type Service struct {
	ProductSVC *services.ProductService
}

type Dependency struct {
	Repository *Repository
	Service    *Service
}

func initRepositories() *Repository {
	var r Repository

	r.ProductRepo = repositories.NewProductRepository(config.DB())

	return &r
}

func initServices(repo *Repository) *Service {
	return &Service{
		ProductSVC: services.NewProductService(repo.ProductRepo, config.REDIS()),
	}
}

func Dependencies() *Dependency {
	repository := initRepositories()
	service := initServices(repository)

	return &Dependency{
		Repository: repository,
		Service:    service,
	}
}
