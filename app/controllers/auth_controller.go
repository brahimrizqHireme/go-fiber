package controllers

import (
	"github.com/brahimrizqHireme/go-fiber/app/models"
	services "github.com/brahimrizqHireme/go-fiber/app/services"
	"github.com/brahimrizqHireme/go-fiber/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: *services.NewAuthService(),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user models.User
	user.ID = uuid.New()
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	validationService := utils.NewValidationService(c)
	if err := validationService.Validate(&user); err != nil {
		return err
	}

	err = h.authService.RegisterUser(user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var credentials models.Credentials
	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	token, err := h.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) ProtectedEndpoint(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Protected endpoint accessed successfully",
	})
}
