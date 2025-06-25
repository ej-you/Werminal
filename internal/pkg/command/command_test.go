package command

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	t.Log("Run ping command and read its output (it works)")

	t.Log("Init command")
	cmd, err := New("ping", "ya.ru")
	require.NoError(t, err)
	defer cmd.cleanup()

	t.Log("Start command")
	err = cmd.Start(os.Stdout, os.Stdin)
	require.NoError(t, err)

	err = cmd.Wait()
	require.NoError(t, err)
}

func TestBash(t *testing.T) {
	t.Log("Open new bash shell (it does NOT work)")

	t.Log("Init command")
	cmd, err := New("bash")
	require.NoError(t, err)
	defer cmd.cleanup()

	t.Log("Start command")
	err = cmd.Start(os.Stdout, os.Stdin)
	require.NoError(t, err)

	err = cmd.Wait()
	require.NoError(t, err)
}
