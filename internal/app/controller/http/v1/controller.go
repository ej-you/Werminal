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

// Terminal handles WebSocket connection with terminal input/output.
// It can parse rows and cols query params to setup custom terminal size.
func (c *TerminalController) Terminal() fiber.Handler {
	return gowebsocket.New(func(conn *gowebsocket.Conn) {
		dataIn, err := parseTerminalIn(conn)
		if err != nil {
			logrus.Errorf("400 bad request: invalid query: %v", err)
			return
		}

		logrus.Infof("New connection to terminal with size %dx%d", dataIn.Rows, dataIn.Cols)
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
				logrus.Warnf("close ws conn: %v", err)
			}
		}()

		// create new terminal
		term, err := terminal.New(dataIn.Rows, dataIn.Cols)
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
			logrus.Warn(err)
		}
		// cancel ctx to close ws conn (if it not closed)
		cancel()
		<-done
		logrus.Info("Disconnect from terminal")
	}, c.wsConfig)
}
