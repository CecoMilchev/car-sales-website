package repos

import (
	"fmt"

	"github.com/CecoMilchev/car-sales-website/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func (repository *UserRepository) FindAll() []models.User {
	var users []models.User

	repository.database.Find(&users)

	return users
}

func (repository *UserRepository) FindByField(field string, value interface{}) []models.User {
	user := models.User{}
	var users []models.User

	repository.database.
		Where(fmt.Sprintf("%v = ?", field), value).
		Find(&user)

	users = append(users, user)

	return users
}

func (repository *UserRepository) CreateUser(user models.User) []models.User {
	var users []models.User

	repository.database.Create(&user)

	users = append(users, user)

	return users
}

func (repository *UserRepository) UpdateUser(user models.User) []models.User {
	var users []models.User

	repository.database.Save(&user)

	users = append(users, user)

	return users
}

func (repository *UserRepository) DeleteUser(user models.User) []models.User {
	var users []models.User

	repository.database.Delete(&user)

	users = append(users, user)

	return users
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database: database}
}
