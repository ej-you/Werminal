// Package app provides function Run to start full application.
package app

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"Werminal/config"
	"Werminal/internal/pkg/logger"
	// "Werminal/internal/app/server"
)

// Run loads app config and starts HTTP-server.
// This function is blocking.
func Run() error {
	logger.Init()
	// create config
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("create config: %w", err)
	}
	logrus.Info("cfg:", cfg)
	// create and run HTTP-server
	// srv := server.New(cfg)
	// srv.Run()
	// // wait for HTTP-server shutdown
	// if err := srv.WaitForShutdown(); err != nil {
	// 	return fmt.Errorf("http server: %w", err)
	// }
	return nil
}
