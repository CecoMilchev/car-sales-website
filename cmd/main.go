package main

import (
	"fmt"
	"os"

	"github.com/CecoMilchev/car-sales-website/internal/models"
	"github.com/CecoMilchev/car-sales-website/internal/repos"
	"github.com/CecoMilchev/car-sales-website/pkg/server"
	"github.com/CecoMilchev/car-sales-website/pkg/services"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	// github.com/denisenkom/go-mssqldb
)

// "github.com/car-sales-website/internal/repos"

func main() {
	config := models.NewConfig()

	login := os.Getenv("GOSQLSERVERLOGIN")
	dsn := fmt.Sprintf("sqlserver://%s@localhost:1433?database=CarSales", login)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	//db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := repos.NewCarRepository(db)

	personService := services.NewCarService(config, personRepository)

	server := server.NewServer(config, personService)

	server.Run()
}
