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
	config      *models.Config
	userService *services.UserService
	carService  *services.CarService
}

func setupRouter(s *Server) *gin.Engine {
	r := gin.Default()

	r.GET("/cars", findCars(s))
	r.GET("/cars/:id", findCar(s))
	r.POST("/cars", createCar(s))
	r.PUT("/cars/:id", updateCar(s))
	r.DELETE("/cars/:id", deleteCar(s))

	r.POST("/register", registerUser(s))
	r.POST("/login", login(s))

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

func updateCar(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input models.Car
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// if err := db.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		//   c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		//   return
		// }

		//db.Model(&car).Updates(input)

		c.JSON(http.StatusOK, gin.H{"data": s.carService.UpdateCar(input)})
	}

	return gin.HandlerFunc(fn)
}

func deleteCar(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input models.Car
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": s.carService.DeleteCar(input)})
	}

	return gin.HandlerFunc(fn)
}

func registerUser(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		s.userService.CreateUser(input)
		c.JSON(http.StatusOK, gin.H{"message": "Successful registration!"})
	}

	return gin.HandlerFunc(fn)
}

func login(s *Server) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input models.User
		u := models.User{}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u.UserName = input.UserName
		u.Password = input.Password

		token, err := models.LoginCheck(s.userService.FindByField("username", u.UserName)[0], u.UserName, u.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
			return
		}

		//s.userService.CreateUser(input)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}

	return gin.HandlerFunc(fn)
}

func NewServer(config *models.Config, userService *services.UserService, carService *services.CarService) *Server {
	return &Server{
		config:      config,
		userService: userService,
		carService:  carService,
	}
}
