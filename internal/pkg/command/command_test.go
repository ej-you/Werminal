package command

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	t.Log("Run ping command and read its output")

	t.Log("Init command")
	cmd, err := New("ping", "ya.ru")
	require.NoError(t, err)
	defer cmd.clear()

	reader := strings.NewReader("")

	t.Log("Run command")
	err = cmd.Run(os.Stdout, reader)
	require.NoError(t, err)

	// out := make(chan []byte)
	// defer close(out)
	// t.Log("Read command output")

	// go cmd.Output(out)
	// for line := range out {
	// 	t.Log("output:", string(line))
	// }
}

func TestTelnet(t *testing.T) {
	t.Log("Run telnet command, read its output and open writer to stdin")

	t.Log("Init command")
	cmd, err := New("bash")
	require.NoError(t, err)
	defer cmd.clear()

	t.Log("Run command")
	err = cmd.Run(os.Stdout, os.Stdin)
	require.NoError(t, err)

}
