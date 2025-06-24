// Package http/v1 is a first version of HTTP-controller for all entities.
// It provides register for HTTP-routes and controller with handlers for them.
package v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	websocket "github.com/gofiber/websocket/v2"
)

// TerminalController is a HTTP-controller for task usecase.
type TerminalController struct {
	wsConfig websocket.Config
	// taskUC task.Usecase
}

// NewTerminalController returns new TerminalController.
func NewTerminalController(readBufferSize, writeBufferSize int,
	pongWait, pingPeriod time.Duration) *TerminalController {

	return &TerminalController{
		wsConfig: websocket.Config{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
		},
		// taskUC: taskUC,
	}
}

// UpgradeWebSocket upgrades HTTP-connection to WebSocket and register client handlers.
func (c *TerminalController) Terminal() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		// create new WebSocket client and start client handlers
		client := newClientWS(conn)
		go client.HandleRead()
		go client.HandleWrite()
		// wait for client disconnection
		client.Wait()
	}, c.wsConfig)
}
