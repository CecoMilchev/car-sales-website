package main

import (
	"github.com/CecoMilchev/car-sales-website/internal/models"
	"github.com/CecoMilchev/car-sales-website/internal/repos"
	"github.com/CecoMilchev/car-sales-website/pkg/server"
	"github.com/CecoMilchev/car-sales-website/pkg/services"
)

// "github.com/car-sales-website/internal/repos"

func main() {
	config := models.NewConfig()

	//db, err := ConnectDatabase(config)

	//if err != nil {
	//	panic(err)
	//}

	personRepository := repos.NewCarRepository()

	personService := services.NewCarService(config, personRepository)

	server := server.NewServer(config, personService)

	server.Run()
}
