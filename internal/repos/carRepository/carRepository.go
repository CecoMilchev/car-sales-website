package repos

import (
	"database/sql"

	"github.com/car-sales-website/internal/models"
)

type CarRepository struct {
	database *sql.DB
}

func (repository *CarRepository) FindAll() []*models.Car {
	//rows, _ := repository.database.Query(
	//  `SELECT id, name, age FROM people;`
	//)
	//defer rows.Close()
	//
	cars := []*models.Car{}
	//
	//for rows.Next() {
	//  var (
	//	id   int
	//	name string
	//	age  int
	//  )
	//
	//  rows.Scan(&id, &name, &age)

	cars = append(cars, &models.Car{
		Id:    1,
		Make:  "BMW",
		Model: "M2Comp",
	})
	//}

	return cars
}

func NewPersonRepository(database *sql.DB) *CarRepository {
	return &CarRepository{database: database}
}
