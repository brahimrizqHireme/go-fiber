package middlewares

import (
	"github.com/brahimrizqHireme/go-fiber/app/configs"
	"github.com/brahimrizqHireme/go-fiber/app/repositories"
	"github.com/brahimrizqHireme/go-fiber/app/utils"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

type JwtMiddleware struct {
	jwtParser *utils.JWTParser
	userRepo  *repositories.UserRepository
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{
		jwtParser: utils.NewJWTParser(),
		userRepo:  repositories.NewUserRepository(),
	}
}

func (j *JwtMiddleware) JWTProtected() func(c *fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:    []byte(configs.AppConfig.JWTSecretKey),
		ContextKey:    "jwt",
		SigningMethod: "HS256",
		AuthScheme:    "Bearer",
		ErrorHandler:  jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			authorization := c.Get("Authorization")
			if authorization == "" {
				return fiber.NewError(fiber.StatusUnauthorized, "Missing access token")
			}
			token := c.Locals("jwt").(*jwt.Token)
			TokenMetadata, err := j.jwtParser.ExtractTokenMetadata(*token)
			if err != nil {
				return err
			}

			user, err := j.userRepo.FindByID(TokenMetadata.UserID)
			if err != nil {
				return fiber.NewError(fiber.StatusNotFound, "User not found")
			}

			c.Locals("userId", TokenMetadata.UserID)
			c.Locals("user", user)

			return c.Next()
		},
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		return fiber.NewError(fiber.StatusForbidden, "Missing access token")
	}

	return fiber.NewError(fiber.StatusForbidden, "Invalid access token")
}
