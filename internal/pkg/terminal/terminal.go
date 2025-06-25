// Package terminal provides interface to create and run pseudo-terminal.
package terminal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/creack/pty"
	"github.com/pkg/errors"
)

// Terminal is a pseudo-terminal.
type Terminal interface {
	Run(output io.Writer, input io.Reader) error
	Wait() error
}

// pterm is the Terminal implementation.
type pterm struct {
	cmd     *exec.Cmd
	size    *pty.Winsize
	ptyFile *os.File
	wg      sync.WaitGroup
}

// New returns new Terminal instance.
// Rows and columns parameters set the terminal size.
func New(rows, columns uint16) (Terminal, error) {
	// parse terminal shell value
	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		return nil, errors.New("terminal shell not specified")
	}
	if len(shell) == 0 {
		return nil, errors.New("invalid terminal shell value")
	}

	return &pterm{
		cmd: exec.Command(shell),
		size: &pty.Winsize{
			Rows: rows,
			Cols: columns,
		},
	}, nil
}

// Run starts pseudo-terminal. It uses given writer
// for term output and reader for term input.
func (t *pterm) Run(output io.Writer, input io.Reader) error {
	var err error
	// start pseudo-terminal with init size
	t.ptyFile, err = pty.StartWithSize(t.cmd, t.size)
	if err != nil {
		return fmt.Errorf("start pty: %w", err)
	}

	// read output from pty and write to output writer
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		io.Copy(output, t.ptyFile)
	}()

	// read input from input reader and write to pty
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		io.Copy(t.ptyFile, input)
	}()

	return nil
}

// Wait waits for started pterm to finish.
// This method is blocking.
func (t *pterm) Wait() error {
	defer t.ptyFile.Close()

	// wait for pty to finish
	err := t.cmd.Wait()
	// wait for streams
	t.wg.Wait()

	return errors.Wrap(err, "wait for pty to finish") // err OR nil
}
