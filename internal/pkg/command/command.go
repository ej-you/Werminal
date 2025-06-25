// Package terminal provides interface Cmd to run external command and
// handle its standart streams: STDOUT, STDERR and STDIN.
//
// ! NOTE !
// Now input to STDIN is not working. Cmd works correctly
// only with data output from STDOUT and STDERR.
package command

import (
	"fmt"
	"io"
	"os/exec"
	"sync"
)

// Cmd is an external command.
type Cmd struct {
	// wait group for command standart streams handlers
	wg sync.WaitGroup
	// exec command
	command *exec.Cmd

	// stdout reader
	stdoutPipe io.ReadCloser
	// stderr reader
	stderrPipe io.ReadCloser
	// stdin writer
	stdinPipe io.WriteCloser
}

// New returns new command instance.
func New(name string, arg ...string) (*Cmd, error) {
	// init command
	command := exec.Command(name, arg...)

	// init readers and writers
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("init stdout pipe: %w", err)
	}
	stderrPipe, err := command.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("init stderr pipe: %w", err)
	}
	stdinPipe, err := command.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("init stdin pipe: %w", err)
	}

	return &Cmd{
		command:    command,
		stdoutPipe: stdoutPipe,
		stderrPipe: stderrPipe,
		stdinPipe:  stdinPipe,
	}, nil
}

// Start starts command. It uses given writer for
// command output and reader for command input.
func (c *Cmd) Start(w io.Writer, r io.Reader) error {
	if err := c.command.Start(); err != nil {
		return fmt.Errorf("start command: %w", err)
	}

	// read output from command and write to writer
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		io.Copy(w, io.MultiReader(c.stdoutPipe, c.stderrPipe))
	}()

	// read input from reader and write to command stdin
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		io.Copy(c.stdinPipe, r)
	}()

	return nil
}

// Wait waits for started command to finish.
// This method is blocking.
func (c *Cmd) Wait() error {
	defer c.cleanup()

	err := c.command.Wait()
	c.wg.Wait()

	if err != nil {
		return fmt.Errorf("wait for cmd to finish: %w", err)
	}
	return nil
}

// cleanup cleans up resources used for command execution.
func (c *Cmd) cleanup() {
	if c.stdinPipe != nil {
		c.stdinPipe.Close()
	}
	if c.stdoutPipe != nil {
		c.stdoutPipe.Close()
	}
	if c.stderrPipe != nil {
		c.stderrPipe.Close()
	}
}
