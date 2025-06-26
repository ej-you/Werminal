package websocket

import (
	"context"
	"fmt"
	"io"
	"sync"

	websocket "github.com/gofiber/websocket/v2"
)

var _ io.ReadWriteCloser = (*Stream)(nil)

// Stream provides write ans read methods for websocket connection.
// It implements io.ReadWriteCloser.
type Stream struct {
	// conn with client
	conn *websocket.Conn
	// buffer with left data (not sent to reader) of the last message
	buf []byte
	// for write sync
	mu sync.Mutex

	// for Close
	once sync.Once
	// send value to it if conn is closed (used for Wait method)
	done chan struct{}
}

// NewStream returns new WebSocketRW instance.
func NewStream(conn *websocket.Conn) *Stream {
	return &Stream{
		conn: conn,
		buf:  make([]byte, 0),
		done: make(chan struct{}),
	}
}

// Read returns data from buffer (if it is not empty) or wait for
// new message from websocket client and return gotten message.
// It implements io.Reader.
func (w *Stream) Read(p []byte) (int, error) {
	if len(w.buf) == 0 {
		// get message from client
		_, byteMessage, err := w.conn.ReadMessage()
		// if connection is closed
		if err != nil {
			w.done <- struct{}{}
			return 0, err
		}
		// save message to buffer
		w.buf = byteMessage
	}
	// copy buffer to reader slice
	written := copy(p, w.buf)
	// cut buffer and left only excess part of message
	w.buf = w.buf[written:]

	return written, nil
}

// Write sends given bytes to websocket client as text message.
// It implements io.Writer.
func (w *Stream) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// send message
	err := w.conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Wait waits for close message from client.
// It immediately closes connection if given context is done.
func (w *Stream) Wait(ctx context.Context) error {
	select {
	// wait for close message from client
	case <-w.done:
	// wait for context
	case <-ctx.Done():
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("close stream: %w", err)
	}
	return nil
}

// Close closes websocket connection.
// It implements io.Closer.
func (w *Stream) Close() error {
	var err error
	// close conn and done chan
	w.once.Do(func() {
		defer close(w.done)
		err = w.conn.Close()
	})
	return err
}
