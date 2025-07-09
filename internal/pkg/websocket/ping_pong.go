package websocket

import (
	"context"
	"fmt"
	"time"

	websocket "github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

// SetupPingPong setup ping-pong handler for WebSocket connection.
// Can be stopped by closing given context.
func SetupPingPong(ctx context.Context, conn *websocket.Conn, readTimeout time.Duration) error {

	// set initial read deadline
	err := conn.SetReadDeadline(getDeadline(readTimeout))
	if err != nil {
		return fmt.Errorf("set init read deadline: %w", err)
	}

	// after each gotten pong message from client
	conn.SetPongHandler(func(string) error {
		// update read deadline
		return conn.SetReadDeadline(getDeadline(readTimeout))
	})
	go func() {
		// calc ping period based on read timeout
		pingPeriod := readTimeout * 9 / 10
		// create ticker for sending ping messages
		pingTicker := time.NewTicker(pingPeriod)
		defer pingTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-pingTicker.C:
				err := conn.WriteControl(websocket.PingMessage, nil, getDeadline(readTimeout))
				if err != nil {
					logrus.Warnf("send ping: %v", err)
					return
				}
			}
		}
	}()
	return nil
}

// getDeadline returns deadline from now time (in UTC) with added timeout duration.
func getDeadline(timeout time.Duration) time.Time {
	return time.Now().UTC().Add(timeout)
}
