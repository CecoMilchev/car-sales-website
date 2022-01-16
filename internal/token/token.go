package token

import (
	"fmt"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user_id int) (string, error) {
	fmt.Print("-----------\n")
	fmt.Print(os.Getenv("GOTOKENHOURLIFESPAN"))
	fmt.Print("-----------\n")

	token_lifespan, err := strconv.Atoi(os.Getenv("GOTOKENHOURLIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("GOAPISECRET")))
}
