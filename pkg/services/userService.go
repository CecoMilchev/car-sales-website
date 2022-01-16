package services

import (
	"github.com/CecoMilchev/car-sales-website/internal/models"
	"github.com/CecoMilchev/car-sales-website/internal/repos"
)

type UserService struct {
	config     *models.Config
	repository *repos.UserRepository
}

func (service *UserService) FindAll() []models.User {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []models.User{}
}

func (service *UserService) FindByField(field string, value interface{}) []models.User {
	if service.config.Enabled {
		return service.repository.FindByField(field, value)
	}

	return []models.User{{}}
}

func (service *UserService) CreateUser(user models.User) []models.User {
	if service.config.Enabled {
		return service.repository.CreateUser(user)
	}

	return []models.User{{}}
}

func (service *UserService) UpdateUser(user models.User) []models.User {
	if service.config.Enabled {
		return service.repository.UpdateUser(user)
	}

	return []models.User{{}}
}

func (service *UserService) DeleteUser(user models.User) []models.User {
	if service.config.Enabled {
		return service.repository.DeleteUser(user)
	}

	return []models.User{{}}
}

func NewUserService(config *models.Config, repository *repos.UserRepository) *UserService {
	return &UserService{config: config, repository: repository}
}
