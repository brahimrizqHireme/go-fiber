package middlewares

import (
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddlewares(app *fiber.App) {
	app.Use(SetDefaultLanguage("en"))
	app.Use(recoverMiddleware())
	app.Use(loggerMiddleware())
	app.Use(corsMiddleware())
	app.Use(compressMiddleware())
	app.Use(contentSecurityPolicyMiddleware())
	app.Use(secureHeadersMiddleware())
	// app.Use(secureMiddleware())
	// app.Use(csrfMiddleware())
	// app.Use(pprofMiddleware())
	//app.Use(contentTypeMiddleware())
	app.Use(rateLimitMiddleware())

	// app.Use(ErrorHandlerMiddleware(&DefaultErrorHandler{}))
}
func SetDefaultLanguage(defaultLang string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Create a new validator instance
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator(defaultLang)

		// Register the English translation for the validator
		validate := validator.New()

		// Register the translator with the validator
		enTranslations.RegisterDefaultTranslations(validate, trans)

		// Set the default language to "en"
		ctx.Locals("trans", trans)
		ctx.Locals("validate", validate)

		// Continue to the next middleware or route handler
		return ctx.Next()
	}
}

func recoverMiddleware() fiber.Handler {
	return recover.New()
}

func loggerMiddleware() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} ${method} ${path} - ${ip} - ${status}\n",
		TimeFormat: "2006-01-02 15:04:05",
	})
}

func corsMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	})
}

func contentTypeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")
		if contentType != "application/json" {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid Content-Type. Only application/json is accepted.")
		}
		return c.Next()
	}
}

func compressMiddleware() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
}

func contentSecurityPolicyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	}
}

func rateLimitMiddleware() fiber.Handler {
	// Create rate limiting middleware with desired options
	rateLimiter := limiter.New(limiter.Config{
		Max:        100,           // Maximum number of requests within the duration
		Expiration: 1 * time.Hour, // Duration for which the limit counts
	})

	return rateLimiter
}

func secureHeadersMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Referrer-Policy", "no-referrer")
		c.Set("Content-Type", "application/json")
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	}
}

func secureMiddleware() fiber.Handler {
	config := helmet.Config{
		XSSProtection:             "1; mode=block",
		ContentTypeNosniff:        "nosniff",
		XFrameOptions:             "SAMEORIGIN",
		HSTSMaxAge:                31536000,
		HSTSExcludeSubdomains:     false,
		ContentSecurityPolicy:     "default-src 'self'",
		CSPReportOnly:             false,
		HSTSPreloadEnabled:        false,
		ReferrerPolicy:            "same-origin",
		PermissionPolicy:          "",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "same-origin",
		OriginAgentCluster:        "?1",
		XDNSPrefetchControl:       "off",
		XDownloadOptions:          "noopen",
		XPermittedCrossDomain:     "none",
	}

	return helmet.New(config)
}

func csrfMiddleware() fiber.Handler {

	config := csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_token",
		CookieSameSite: "Strict",
	}

	return csrf.New(config)
}

func pprofMiddleware() fiber.Handler {
	config := pprof.Config{
		Prefix: "/debug",
	}
	return pprof.New(config)
}
