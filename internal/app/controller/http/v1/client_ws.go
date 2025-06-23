package v1

import (
	"time"

	websocket "github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

// incomingMessage is an incoming message from client.
type incomingMessage []byte

// clientWS is a client for WebSocket connection.
type clientWS struct {
	// conn with client
	conn *websocket.Conn
	// duration for wait client message
	pongWait time.Duration
	// send one ping message per this period
	pingPeriod time.Duration

	// chan for incoming messages
	message chan incomingMessage
	// for graceful shutdown
	done chan struct{}
}

// newClientWS returns new clientWS instance.
func newClientWS(conn *websocket.Conn, pongWait, pingPeriod time.Duration) *clientWS {
	return &clientWS{
		conn:       conn,
		pongWait:   pongWait,
		pingPeriod: pingPeriod,
		message:    make(chan incomingMessage),
		done:       make(chan struct{}),
	}
}

// HandleRead handles new incoming message.
func (c *clientWS) HandleRead() {
	defer func() {
		c.conn.Close()
		close(c.message)
	}()

	for {
		_, byteMessage, err := c.conn.ReadMessage()
		if err != nil {
			// if unexpected close error occurs (CloseGoingAway, CloseAbnormalClosure)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				logrus.Errorf("handle read: unexpected close ws conn: %v", err)
			}
			return // conn closed by client
		}
		// send gotten message to chan for writer handler
		c.message <- byteMessage
	}
}

// HandleWrite waits for new incoming message in client's chan and processes it.
func (c *clientWS) HandleWrite() {
	defer func() {
		c.conn.Close()
		close(c.done)
	}()

	for {
		// new message
		msg, ok := <-c.message
		// close conn if chan is closed
		if !ok {
			if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				logrus.Errorf("handle write: close ws conn: %v", err)
			}
			return
		}
		// process incoming message
		c.processMessage(msg)
	}
}

// Wait waits for client is done.
// This method is blocking.
func (c *clientWS) Wait() {
	<-c.done
}

// processMessage processes incoming message and send answer to client.
func (c *clientWS) processMessage(msg incomingMessage) {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		logrus.Errorf("send message: %v", err)
		return
	}
}
