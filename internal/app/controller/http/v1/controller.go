// Package http/v1 is a first version of HTTP-controller for all entities.
// It provides register for HTTP-routes and controller with handlers for them.
package v1

import (
	"context"

	"github.com/gofiber/fiber/v2"
	gowebsocket "github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"Werminal/internal/pkg/terminal"
	"Werminal/internal/pkg/websocket"
)

const (
	_termRows = 180 // pseudo-terminal rows
	_termCols = 35  // pseudo-terminal columns
)

// TerminalController is a HTTP-controller for task usecase.
type TerminalController struct {
	wsConfig gowebsocket.Config
}

// NewTerminalController returns new TerminalController.
func NewTerminalController(readBufferSize, writeBufferSize int) *TerminalController {
	return &TerminalController{
		wsConfig: gowebsocket.Config{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
		},
	}
}

// // UpgradeWebSocket upgrades HTTP-connection to WebSocket and register client handlers.
// func (c *TerminalController) Terminal() fiber.Handler {
// 	return websocket.New(func(conn *websocket.Conn) {
// 		// create new WebSocket client and start client handlers
// 		client := newClientWS(conn)
// 		go client.HandleRead()
// 		go client.HandleWrite()
// 		// wait for client disconnection
// 		client.Wait()
// 	}, c.wsConfig)
// }

// Terminal handles WebSocket connection with terminal input/output.
// TODO: add query params for terminal rows and columns
func (c *TerminalController) Terminal() fiber.Handler {
	return gowebsocket.New(func(conn *gowebsocket.Conn) {
		// context for stream and terminal
		ctx, cancel := context.WithCancel(context.Background())

		// create new WebSocket stream
		wsStream := websocket.NewStream(conn)
		// wait for ws conn to close
		done := make(chan struct{})
		go func() {
			defer close(done)
			defer cancel()
			if err := wsStream.Wait(ctx); err != nil {
				logrus.Errorf("close ws conn: %v", err)
			}
		}()

		// create new terminal
		term, err := terminal.New(_termRows, _termCols)
		if err != nil {
			logrus.Errorf("create terminal: %v", err)
			return
		}
		// run terminal
		if err := term.Run(wsStream, wsStream); err != nil {
			logrus.Errorf("start terminal: %v", err)
			return
		}

		// wait for terminal to finish
		if err := term.Wait(ctx); err != nil {
			logrus.Error(err)
		}
		// cancel ctx to close ws conn (if it not closed)
		cancel()
		<-done
	}, c.wsConfig)
}
