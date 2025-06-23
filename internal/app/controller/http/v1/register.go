package v1

import (
	fiber "github.com/gofiber/fiber/v2"
)

// RegisterEndpoints registers all endpoints.
func RegisterEndpoints(router fiber.Router, controller *TerminalController) {
	router.Get("/", controller.Terminal())
}
