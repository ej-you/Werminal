// Package terminal provides interface Cmd to run external command and
// handle its standart streams: STDOUT, STDERR and STDIN.
package command

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
)

// Cmd is an external command.
type Cmd struct {
	// context for command
	ctx context.Context
	// cancel func for context for command
	ctxCancel context.CancelFunc
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

// Run runs command and wait for it to finish. This method is blocking.
// It uses given writer for command output and reader for command input.
func (c *Cmd) Run(w io.Writer, r io.Reader) error {
	if err := c.command.Start(); err != nil {
		return fmt.Errorf("start command: %w", err)
	}

	// chan for reader/writer exit signal
	done := make(chan error, 2)
	defer close(done)

	// read output from command and write to writer
	go func() {
		fmt.Println("BBBBBBB")
		var err error
		scanner := bufio.NewScanner(io.MultiReader(c.stdoutPipe, c.stderrPipe))
		for scanner.Scan() {
			fmt.Println("EEEEEE")
			_, err = w.Write(scanner.Bytes())
			fmt.Println("FFFFFF")
			if err != nil {
				fmt.Println("HHHHHH")
				fmt.Println("err:", err)
				done <- fmt.Errorf("write to writer: %w", err)
				return
			}
			fmt.Println("GGGGGG")
		}
		done <- fmt.Errorf("scan stdout and stderr: %w", scanner.Err())
	}()

	// read input from reader and write to command stdin
	go func() {
		fmt.Println("CCCCCCC")
		var err error
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Println("DDDD")
			_, err = c.stdinPipe.Write(append(scanner.Bytes(), '\n'))
			if err != nil {
				done <- fmt.Errorf("write to stdin: %w", err)
				return
			}
		}
		done <- fmt.Errorf("scan reader: %w", scanner.Err())
	}()

	fmt.Println("AAAAAAAA")
	err := <-done
	c.clear()
	fmt.Println("done err:", err)
	// wait for the command to finish
	if err := c.command.Wait(); err != nil {
		return fmt.Errorf("wait for cmd to finish: %w", err)
	}
	fmt.Println("IIIIIII")
	return err
}

// clear cleans up resources for command execution.
func (c *Cmd) clear() {
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
