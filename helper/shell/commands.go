/*
Copyright © 2026 Matze
*/
package shell

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func ExecuteCommand(cmd *exec.Cmd) error {
	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}
	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)

	go func() {
		<-sigChan
		if cmd.Process != nil {
			_ = cmd.Process.Signal(os.Interrupt)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func ExecuteCommandReturn(cmd *exec.Cmd) (string, error) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		return "", err
	}

	sigChan := make(chan os.Signal, 1)

	go func() {
		<-sigChan
		if cmd.Process != nil {
			_ = cmd.Process.Signal(os.Interrupt)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("command execution failed: %w (stderr: %s)", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func ExecuteCommandOnlyErrors(execArgs []string) error {
	cmd := exec.Command(execArgs[0], execArgs[1:]...)
	var stderrBuf bytes.Buffer

	cmd.Stdout = io.Discard
	cmd.Stderr = &stderrBuf

	err := ExecuteCommand(cmd)
	if err != nil {
		if stderrBuf.Len() > 0 {
			return errors.New(stderrBuf.String())
		}
		return err
	}

	return nil
}
