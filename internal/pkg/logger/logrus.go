// Package logger provides Init function to setup global logrus logger.
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// textFormatterUTC is the logrus.JSONFormatter wrapper with time in UTC.
type textFormatterUTC struct {
	logrus.TextFormatter
}

// newTextFormatterUTC returns new textFormatterUTC instance.
func newTextFormatterUTC() *textFormatterUTC {
	return &textFormatterUTC{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp: true,
		},
	}
}

// Format implements logrus.Formatter and sets time to UTC.
func (f *textFormatterUTC) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return f.TextFormatter.Format(e)
}

// Init sets up main logger for application.
func Init() {
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(newTextFormatterUTC())
}
