package server

import (
	goerrors "errors"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// CustomErrorHandler is a handler for http server errors.
func customErrorHandler(_ *fiber.Ctx, err error) error {
	msg := err.Error()
	// if resource was not found
	var fiberErr *fiber.Error
	if goerrors.As(err, &fiberErr) && strings.HasPrefix(fiberErr.Message, "Cannot GET") {
		msg = "resource not found"
	}
	// log error
	logrus.Error(msg)
	return nil
}
