package services

import (
		"github.com/car-sales-website/internal/models"
		"github.com/car-sales-website/internal/repos"
)

type CarService struct {
	config     *Config
	repository *repos.CarRepository
  }
  
  func (service *CarService) FindAll() []*Person {
	if service.config.Enabled {
	  return service.repository.FindAll()
	}
  
	return []*Person{}
  }
  
  func NewCarService(config *Config, repository *repos.CarRepository)
  *CarService {
	return &CarService{config: config, repository: repository}
  }