// App binary starts WebSocket server.
package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"Werminal/internal/pkg/terminal"
)

func main() {
	// if err := app.Run(); err != nil {
	// 	logrus.Fatal(err)
	// }

	term, err := terminal.New(45, 180)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := term.Run(os.Stdout, os.Stdin); err != nil {
		logrus.Fatal(err)
	}
	if err := term.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
