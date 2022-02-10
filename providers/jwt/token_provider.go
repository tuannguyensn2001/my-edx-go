package jwt_provider

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"my-edx-go/models"
	"time"
)

type MyCustomClaims struct {
	User models.User `json:"user"`
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

func ValidateToken(token string) (*models.User, error) {
	res, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		return nil, err
	}

	if !res.Valid {
		return nil, err
	}

	claims, ok := res.Claims.(*MyCustomClaims)

	if !ok {
		return nil, errors.New("Payload not valid")
	}

	return &claims.User, nil

}
