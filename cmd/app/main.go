// App binary starts WebSocket server.
package main

import (
	"github.com/sirupsen/logrus"

	"Werminal/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}
