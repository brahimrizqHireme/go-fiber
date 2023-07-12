package utils

import (
	"strings"

	"github.com/brahimrizqHireme/go-fiber/app/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID  string
	Expires int64
}

type JWTParser struct {
	TokenMetadata TokenMetadata
	secretKey     string
}

func NewJWTParser() *JWTParser {
	return &JWTParser{
		TokenMetadata: TokenMetadata{},
		secretKey:     configs.AppConfig.JWTSecretKey,
	}
}

func (j *JWTParser) VerifyAndExtractTokenMetadata(authorization string) (*TokenMetadata, error) {
	tokenClaim, err := j.VerifyToken(authorization)
	if err != nil {
		return nil, err
	}

	TokenMetadata, err := j.ExtractTokenMetadata(*tokenClaim)
	if err != nil {
		return nil, err
	}

	return TokenMetadata, nil
}

func (j *JWTParser) ExtractTokenMetadata(token jwt.Token) (*TokenMetadata, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}

	expires := int64(claims["exp"].(float64))
	userID := claims["userId"].(string)

	return &TokenMetadata{
		UserID:  userID,
		Expires: expires,
	}, nil

}

func extractToken(authorization string) (string, error) {

	// Extract the token from the "Bearer token" format
	authParts := strings.SplitN(authorization, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid access token format")
	}

	return authParts[1], nil
}

func (j *JWTParser) VerifyToken(authorization string) (*jwt.Token, error) {
	tokenString, err := extractToken(authorization)
	if err != nil {
		return nil, err
	}

	tokenClaim, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token signing method")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}

	if !tokenClaim.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token given")
	}

	return tokenClaim, nil
}
