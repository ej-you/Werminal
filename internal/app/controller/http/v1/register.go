package v1

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// RegisterEndpoints registers all endpoints.
func RegisterEndpoints(router fiber.Router, controller *TerminalController) {
	router.Get("/", adaptor.HTTPHandlerFunc(controller.UpgradeWebSocket))
}
