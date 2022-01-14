package services

import (
	"github.com/CecoMilchev/car-sales-website/internal/models"
	"github.com/CecoMilchev/car-sales-website/internal/repos"
)

type CarService struct {
	config     *models.Config
	repository *repos.CarRepository
}

func (service *CarService) FindAll() []models.Car {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []models.Car{}
}

func NewCarService(config *models.Config, repository *repos.CarRepository) *CarService {
	return &CarService{config: config, repository: repository}
}
