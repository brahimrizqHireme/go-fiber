package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/brahimrizqHireme/go-fiber/app/errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler interface for handling errors
// type ErrorHandler interface {
// 	HandleError(c *fiber.Ctx, err error) error
// }

// // DefaultErrorHandler is an implementation of ErrorHandler
// type DefaultErrorHandler struct{}

func HandleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *fiber.Error:
		logErrorMsg(e.Code, c, err)
		return c.Status(e.Code).JSON(fiber.Map{
			"message": e.Message,
			"error":   true,
		})
	case validator.ValidationErrors:
		trans := c.Locals("trans").(ut.Translator)
		errors := formatValidationErrors(trans, e)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data body",
			"errors":  errors,
			"error":   true,
		})
	case *errors.AppError:
		logErrorMsg(e.StatusCode, c, err)
		return c.Status(e.StatusCode).JSON(fiber.Map{
			"message": e.Error(),
			"error":   true,
		})
	default:
		logErrorMsg(500, c, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   true,
		})
	}

}
func logErrorMsg(code int, c *fiber.Ctx, err error) {
	log.Printf(fmt.Sprintf("error: '%s' - code: %d,  %s %s %s - %s \n", err.Error(), code, time.Now().Format(time.RFC3339), c.Method(), c.Path(), c.IP()))
}

// // ErrorHandlerMiddleware handles errors in the application

//	func ErrorHandlerMiddleware(handler ErrorHandler) fiber.Handler {
//		return func(c *fiber.Ctx) error {
//			err := c.Next()
//			if err != nil {
//				return handler.HandleError(c, err)
//			}
//			return nil
//		}
//	}

// func TranslateValidationError(trans ut.Translator, err error) string {
// 	// Translate the validation error to English
// 	validationErr := err.(validator.ValidationErrors)[0]
// 	fieldName := validationErr.StructField()
// 	fieldErr := validationErr.Translate(trans)
// 	return fmt.Sprintf("%s %s", fieldName, fieldErr)
// }

func formatValidationErrors(trans ut.Translator, err error) map[string]string {
	errors := make(map[string]string)
	validationErr := err.(validator.ValidationErrors)

	for _, err := range validationErr {
		field := err.Field()
		errors[field] = err.Translate(trans)
	}
	return errors
}
