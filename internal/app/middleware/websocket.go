package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	websocket "github.com/gofiber/websocket/v2"
)

// WebSocket is a middleware to upgrade HTTP-connection to WebSocket.
func WebSocket() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// try to upgrade conn, err if failed
		if !websocket.IsWebSocketUpgrade(ctx) {
			return fiber.ErrUpgradeRequired
		}
		return ctx.Next()
	}
}
