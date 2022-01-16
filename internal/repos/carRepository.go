package repos

import (
	//"database/sql"

	"fmt"

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

func (repository *CarRepository) FindByID(id string) models.Car {
	car := models.Car{}

	repository.database.First(&car, id)

	return car
}

func (repository *CarRepository) CreateCar(car models.Car) models.Car {
	repository.database.Create(&car)
	fmt.Print("-----------")
	fmt.Print(&car)
	fmt.Print("-----------")

	return car
}

func NewCarRepository(database *gorm.DB) *CarRepository {
	return &CarRepository{database: database}
}
