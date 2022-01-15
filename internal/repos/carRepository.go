package repos

import (
	//"database/sql"

	"github.com/CecoMilchev/car-sales-website/internal/models"
	"gorm.io/gorm"
)

type CarRepository struct {
	database *gorm.DB
}

func (repository *CarRepository) FindAll() []models.Car {
	var cars []models.Car
	//cars := []*models.Car{}
	repository.database.Find(&cars)
	//defer rows.Close()
	//
	//
	//for rows.Next() {
	//  var (
	//	id   int
	//	name string
	//	age  int
	//  )
	//
	//  rows.Scan(&id, &name, &age)

	//cars = append(cars, &models.Car{
	//	Id:      1,
	//	OwnerId: 1,
	//	Make:    "BMW",
	//	Model:   "M2Comp",
	//	//Category          :"",//string    `json:"category"`
	//	Power:             306,                                                                 //uint      `json:"power"`
	//	Price:             20500,                                                               //float32   `json:"price"`
	//	Mileage:           70000,                                                               //uint      `json:"mileage"`
	//	FuelType:          models.Petrol,                                                       //FuelType  `json:"fuelType"`
	//	Gearbox:           models.Manual,                                                       //Gearbox   `json:"gearbox"`
	//	EmissionsCategory: 5,                                                                   //uint8     `json:"emissionsCategory"`
	//	FirstRegistration: time.Now(),                                                          //time.Time `json:"firstRegistration"`
	//	Color:             "blue",                                                              //string    `json:"color"`
	//	Description:       "Awesome daily driver sports car.",                                  //string    `json:"description"`
	//	Features:          []string{"panoramic roof", "PDC", "LED headlights", "Alloy Wheels"}, //[]string  `json:"features"`
	//})
	//}

	return cars
}

func NewCarRepository(database *gorm.DB) *CarRepository {
	return &CarRepository{database: database}
}

// func NewCarRepository(database *sql.DB) *CarRepository {
// 	return &CarRepository{database: database}
//}
