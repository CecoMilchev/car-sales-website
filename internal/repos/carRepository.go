package repos

import (
	"github.com/CecoMilchev/car-sales-website/internal/models"
	"gorm.io/gorm"
)

type CarRepository struct {
	database *gorm.DB
}

func (repository *CarRepository) FindAll() []models.Car {
	var cars []models.Car

	repository.database.Find(&cars)

	return cars
}

func (repository *CarRepository) FindByID(id string) []models.Car {
	car := models.Car{}
	var cars []models.Car

	repository.database.First(&car, id)

	cars = append(cars, car)

	return cars
}

func (repository *CarRepository) CreateCar(car models.Car) []models.Car {
	var cars []models.Car

	repository.database.Create(&car)

	cars = append(cars, car)

	return cars
}

func (repository *CarRepository) UpdateCar(car models.Car) []models.Car {
	var cars []models.Car

	repository.database.Save(&car)

	cars = append(cars, car)

	return cars
}

func (repository *CarRepository) DeleteCar(car models.Car) []models.Car {
	var cars []models.Car

	repository.database.Delete(&car)

	cars = append(cars, car)

	return cars
}

func NewCarRepository(database *gorm.DB) *CarRepository {
	return &CarRepository{database: database}
}
