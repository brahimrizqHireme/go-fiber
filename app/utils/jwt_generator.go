package utils

import (
	"strconv"
	"time"

	"github.com/brahimrizqHireme/go-fiber/app/configs"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	secretKey string
	expires   string
}

func NewJWT() *JWT {
	config := configs.AppConfig
	return &JWT{
		secretKey: config.JWTSecretKey,
		expires:   config.SecretKeyExpiresMinute,
	}
}

// GenerateNewAccessToken func for generate a new Access token.
func (j *JWT) GenerateNewAccessToken(userID string) (string, error) {

	spew.Dump(time.Now().Zone())
	secret := j.secretKey
	minutesCount, _ := strconv.Atoi(j.expires)
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
