package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationService struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidationService(ctx *fiber.Ctx) *ValidationService {
	trans := ctx.Locals("trans").(ut.Translator)
	validate := ctx.Locals("validate").(*validator.Validate)
	return &ValidationService{
		validate: validate,
		trans:    trans,
	}
}

func (vs *ValidationService) Validate(data interface{}) error {
	err := vs.validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
