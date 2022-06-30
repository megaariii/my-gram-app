package helper

import (
	"errors"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateToken(userId string) (string, error) {
	payload := jwt.MapClaims{}
	payload["id"] = userId
	payload["exp"] = time.Now().Add(time.Hour * 200).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		log.Fatal("Error in Generating key")
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}