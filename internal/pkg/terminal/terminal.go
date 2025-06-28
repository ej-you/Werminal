// Package terminal provides interface of pseudo-terminal
// that has run and wait for shutdown methods.
package terminal

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/creack/pty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

const _disableEcho = false // disable terminal echo flag

// Terminal is a pseudo-terminal.
type Terminal interface {
	Run(output io.Writer, input io.Reader) error
	Wait(ctx context.Context) error
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
func New(rows, cols uint16) (Terminal, error) {
	// parse terminal shell value
	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		return nil, errors.New("terminal shell not specified")
	}
	if shell == "" {
		return nil, errors.New("invalid terminal shell value")
	}

	return &pterm{
		cmd: exec.Command(shell),
		size: &pty.Winsize{
			Rows: rows,
			Cols: cols,
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
		return fmt.Errorf("start pterm: %w", err)
	}
	// disable terminal echo if needed
	if _disableEcho {
		if err := t.disableEcho(); err != nil {
			return fmt.Errorf("disable echo: %w", err)
		}
	}

	// read output from pty and write to output writer
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		if _, err := io.Copy(output, t.ptyFile); err != nil {
			logrus.Warnf("copy to output: %v", err)
		}
	}()

	// read input from input reader and write to pty
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		if _, err := io.Copy(t.ptyFile, input); err != nil {
			logrus.Warnf("copy from input: %v", err)
		}
	}()

	return nil
}

// Wait waits for started pterm to finish.
// It immediately kills terminal if given context is done.
func (t *pterm) Wait(ctx context.Context) error {
	if t.ptyFile == nil {
		return nil
	}

	doneErr := make(chan error)
	defer close(doneErr)
	// task with waiting for pty to finish
	go func() {
		defer t.ptyFile.Close()
		doneErr <- t.cmd.Wait()
	}()

	var err error
	select {
	// wait for pty to finish
	case err = <-doneErr:
	// wait for context
	case <-ctx.Done():
		// kill terminal immediately
		if killErr := t.cmd.Process.Kill(); killErr != nil {
			err = fmt.Errorf("kill pterm: %w", killErr)
		}
		// still waiting for pty to finish
		finishErr := <-doneErr
		if finishErr != nil && err != nil {
			err = errors.Wrap(finishErr, err.Error())
		} else if finishErr != nil {
			err = finishErr
		}
	}

	// wait for input/output streams
	t.wg.Wait()
	return errors.Wrap(err, "stop pterm") // err OR nil
}

// disableEcho disables echo for input commands.
func (t *pterm) disableEcho() error {
	fd := int(t.ptyFile.Fd())

	// get current terminal settings
	termios, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return fmt.Errorf("get current pterm settings: %w", err)
	}
	// unset ECHO flag (put away unix.ECHO bit from termios.Lflag)
	termios.Lflag &^= unix.ECHO

	// apply changes
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, termios); err != nil {
		return fmt.Errorf("change pterm settings: %w", err)
	}
	return nil
}
