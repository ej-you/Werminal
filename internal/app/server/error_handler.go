package server

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// CustomErrorHandler is a handler for http server errors.
func customErrorHandler(_ *fiber.Ctx, err error) error {
	// log error
	logrus.Error(err.Error())
	return nil
}
