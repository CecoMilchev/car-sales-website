package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/CecoMilchev/car-sales-website/internal/token"
)

type User struct {
	ID        int    `json:"id" header:"ID"`
	Password  string `json:"password" gorm:"column:password"`
	UserName  string `json:"userName" gorm:"column:username" header:"Username"`
	FirstName string `json:"firstName" gorm:"column:firstname" header:"First Name"`
	LastName  string `json:"lastName" gorm:"column:lastname" header:"Last Name"`
	Email     string `json:"email" header:"Email"`
	IsPrivate bool   `json:"isPrivate" gorm:"'is_private' boolean"`
	RoleID    int    `json:"roleID" gorm:"'role_id'"` //[role_id]
	//IsActive  bool   `gorm:"'Active' boolean"`
	//Cars      []Car  `gorm:"many2many:cars_users;"`
}

func (u *User) BeforeSave(*gorm.DB) error {
	//fmt.Print("-----------")

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(string(u.Password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(u User, username string, password string) (string, error) {
	var err error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
