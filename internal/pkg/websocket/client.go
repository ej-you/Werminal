// Package websocket provides client and stream to
// handle WebSocket messages.
package websocket

import (
	websocket "github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

// IncomingMessage is an incoming message from client.
type IncomingMessage []byte

// Client is a client for WebSocket connection.
type Client struct {
	// conn with client
	conn *websocket.Conn

	// chan for incoming messages
	message chan IncomingMessage
	// for graceful shutdown
	done chan struct{}
}

// NewClient returns new Client instance.
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:    conn,
		message: make(chan IncomingMessage),
		done:    make(chan struct{}),
	}
}

// HandleRead handles new incoming message.
func (c *Client) HandleRead() {
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
func (c *Client) HandleWrite() {
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
func (c *Client) Wait() {
	<-c.done
}

// processMessage processes incoming message and send answer to client.
func (c *Client) processMessage(msg IncomingMessage) {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		logrus.Errorf("send message: %v", err)
		return
	}
}
