package v1

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// incomingMessage is an incoming message from client.
type incomingMessage []byte

// clientWS is a client for WebSocket connection.
type clientWS struct {
	// conn with client
	conn *websocket.Conn
	// duration for wait pong message
	pongWait time.Duration
	// send one ping message per this period
	pingPeriod time.Duration

	// chan for incoming messages
	message chan incomingMessage
}

// newClientWS returns new clientWS instance.
func newClientWS(conn *websocket.Conn, pongWait, pingPeriod time.Duration) *clientWS {
	return &clientWS{
		conn:       conn,
		pongWait:   pongWait,
		pingPeriod: pingPeriod,
		message:    make(chan incomingMessage),
	}
}

// HandleRead handles new incoming message.
func (c *clientWS) HandleRead() {
	defer c.conn.Close()
	defer close(c.message)

	// set up timeout of reading messages from client
	err := c.conn.SetReadDeadline(time.Now().UTC().Add(c.pongWait))
	if err != nil {
		logrus.Errorf("handle read: set read deadline: %v", err)
		return
	}
	// set up pong handler
	c.conn.SetPongHandler(func(string) error {
		// update read timeout after getting PONG-message from client
		return c.conn.SetReadDeadline(time.Now().UTC().Add(c.pongWait))
	})

	for {
		_, byteMessage, err := c.conn.ReadMessage()
		if err != nil {
			// if unexpected close error occurs (CloseGoingAway, CloseAbnormalClosure)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				logrus.Infof("unexpected close ws conn: %v", err)
			}
			logrus.Info("close ws conn")
			return
		}
		// send gotten message to chan for writer handler
		logrus.Info("received message")
		c.message <- byteMessage
	}
}

// HandleWrite waits for new incoming message in client's chan and processes it.
func (c *clientWS) HandleWrite() {
	defer c.conn.Close()

	// set up ticker for PING-messages to keep alive connection
	pingTicker := time.NewTicker(c.pingPeriod)
	defer pingTicker.Stop()

	for {
		select {
		// send PING-message
		case <-pingTicker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logrus.Warn("close conn")
				return
			}
		// new message
		case msg, ok := <-c.message:
			// close conn if chan is closed
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					logrus.Errorf("close conn: %v", err)
					return
				}
				logrus.Info("close ws conn")
				return
			}
			// process incoming message
			c.processMessage(msg)
		}
	}
}

// processMessage processes incoming message and send answer to client.
func (c *clientWS) processMessage(msg incomingMessage) {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		logrus.Errorf("send message: %v", err)
		return
	}
	logrus.Info("send message")
}
