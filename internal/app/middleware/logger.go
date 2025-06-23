// Package middleware provides all middlewares for HTTP-server.
package middleware

import (
	"os"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// JSON-format for logs
const _logFormat = `INFO[${time}] ${status} | ${method} | ${path} | ${latency} | (err) ${error}"`

// Logger is a middleware for logging all request-response chains.
func Logger() fiber.Handler {
	return logger.New(logger.Config{
		TimeFormat:    time.RFC3339,
		TimeZone:      "UTC",
		Format:        _logFormat + "\n",
		Output:        os.Stderr,
		DisableColors: true,
	})
}
