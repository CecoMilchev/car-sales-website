package server

import (
	//"encoding/json"
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
	r.GET("/cars", people(s))

	return r
}

func (s *Server) Run() {
	r := setupRouter(s)
	r.Run(s.config.Port)
}

func people(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": s.carService.FindAll()})
	}

	return gin.HandlerFunc(fn)
}

func NewServer(config *models.Config, service *services.CarService) *Server {
	return &Server{
		config:     config,
		carService: service,
	}
}
