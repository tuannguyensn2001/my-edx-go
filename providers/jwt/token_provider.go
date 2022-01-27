package jwt_provider

import (
	"github.com/golang-jwt/jwt"
	"log"
	"my-edx-go/models"
	"time"
)

type MyCustomClaims struct {
	user models.User `json:"user"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(3600)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	key := []byte("AllYourBase")

	myToken, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return myToken, nil

}

func ValidateToken(token string) {
	res, _ := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	log.Fatal(res)

}
