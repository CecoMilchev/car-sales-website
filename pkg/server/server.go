package server

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/CecoMilchev/car-sales-website/internal/models"
	"github.com/CecoMilchev/car-sales-website/pkg/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     *models.Config
	carService *services.CarService
}

func setupRouter(s *Server) *gin.Engine {
	r := gin.Default()

	r.GET("/cars", findCars(s))
	r.GET("/cars/:id", findCar(s))
	r.POST("/cars", createCar(s))

	return r
}

func (s *Server) Run() {
	r := setupRouter(s)
	r.Run(s.config.Port)
}

func findCars(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": s.carService.FindAll()})
	}

	return gin.HandlerFunc(fn)
}

func findCar(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": s.carService.FindByID(c.Param("id"))})
	}

	return gin.HandlerFunc(fn)
}

func createCar(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input models.Car
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print("-----------")
		fmt.Print(&input)
		fmt.Print("-----------")

		//book := models.Car{Title: input.Title, Author: input.Author}
		//models.DB.Create(&book)

		c.JSON(http.StatusOK, gin.H{"data": s.carService.CreateCar(input)})
	}

	return gin.HandlerFunc(fn)
}

func NewServer(config *models.Config, service *services.CarService) *Server {
	return &Server{
		config:     config,
		carService: service,
	}
}
