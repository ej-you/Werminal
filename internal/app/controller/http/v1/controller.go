// Package http/v1 is a first version of HTTP-controller for all entities.
// It provides register for HTTP-routes and controller with handlers for them.
package v1

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// TerminalController is a HTTP-controller for task usecase.
type TerminalController struct {
	upgrader   *websocket.Upgrader
	pongWait   time.Duration
	pingPeriod time.Duration
	// taskUC task.Usecase
}

// NewTerminalController returns new TerminalController.
func NewTerminalController(readBufferSize, writeBufferSize int,
	pongWait, pingPeriod time.Duration) *TerminalController {

	return &TerminalController{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
			// allow all origins to upgrade protocol
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
		pongWait:   pongWait,
		pingPeriod: pingPeriod,
		// taskUC: taskUC,
	}
}

// UpgradeWebSocket upgrades HTTP-connection to WebSocket and register client handlers.
func (c *TerminalController) UpgradeWebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade connection to WebSocket
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("upgrade conn to ws: %v", err)
		return
	}
	logrus.Info("open new ws conn")

	// create new WebSocket client and start client handlers
	client := newClientWS(conn, c.pongWait, c.pingPeriod)
	go client.HandleRead()
	go client.HandleWrite()
}
