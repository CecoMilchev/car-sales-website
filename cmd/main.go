package main

import (
	//"database/sql"

	"github.com/CecoMilchev/car-sales-website/internal/models"
	"github.com/CecoMilchev/car-sales-website/internal/repos"
	"github.com/CecoMilchev/car-sales-website/pkg/server"
	"github.com/CecoMilchev/car-sales-website/pkg/services"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	// github.com/denisenkom/go-mssqldb
)

// "github.com/car-sales-website/internal/repos"

func ConnectDatabase(config *models.Config) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(config.DatabasePath), &gorm.Config{})
}

func main() {
	config := models.NewConfig()
	db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := repos.NewCarRepository(db)

	personService := services.NewCarService(config, personRepository)

	server := server.NewServer(config, personService)

	server.Run()
}
