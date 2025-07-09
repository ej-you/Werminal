// Package server provides HTTP-server interface.
package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"Werminal/config"
	http "Werminal/internal/app/controller/http/v1"
	"Werminal/internal/app/middleware"
)

var _ Server = (*httpServer)(nil)

// HTTP-server interface.
type Server interface {
	Run()
	WaitForShutdown() error
}

// HTTP-server implementation.
type httpServer struct {
	cfg *config.Config

	fiberApp *fiber.App
	err      chan error // server listen error
}

// New returns new Server instance.
func New(cfg *config.Config) Server {
	return &httpServer{
		cfg: cfg,
		err: make(chan error),
	}
}

// Run starts server.
func (s *httpServer) Run() {
	// app init
	s.fiberApp = fiber.New(fiber.Config{
		AppName:       "Werminal",
		ServerHeader:  "Werminal Server",
		ErrorHandler:  customErrorHandler,
		StrictRouting: false,
	})

	// set up base middlewares
	s.fiberApp.Use(middleware.Logger())
	s.fiberApp.Use(middleware.Recover())

	// create controllers
	terminalController := http.NewTerminalController(
		s.cfg.Server.WS.ReadBufferSize,
		s.cfg.Server.WS.WriteBufferSize,
		s.cfg.Server.WS.ReadTimeout,
	)

	// register endpoints
	apiV1 := s.fiberApp.Group("/api/v1")
	apiV1WS := apiV1.Group("/ws", middleware.WebSocket())

	http.RegisterEndpoints(apiV1WS.Group("/"), terminalController)

	// start app
	go func() {
		if err := s.fiberApp.Listen(":" + s.cfg.Server.Port); err != nil {
			s.err <- fmt.Errorf("listen: %w", err)
		}
	}()
}

// WaitForShutdown waits for OS signal to gracefully shuts down server.
// This method is blocking.
func (s *httpServer) WaitForShutdown() error {
	// skip if server is not running
	if s.fiberApp == nil {
		return nil
	}

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	shutdownDone := make(chan struct{})
	// create gracefully shutdown task
	var err error
	go func() {
		defer close(shutdownDone)
		select {
		case err = <-s.err: // server listen error
			return
		case handledSignal := <-quit:
			logrus.Infof("Got %s signal. Shutdown server", handledSignal.String())
			// shutdown app
			s.fiberApp.ShutdownWithTimeout(s.cfg.Server.ShutdownTimeout) // nolint:errcheck // cannot occurs
		}
	}()

	// wait for shutdown
	<-shutdownDone
	logrus.Info("Server shutdown successfully")
	return err
}
