/*
Copyright © 2026 Matze
*/
package shell

import (
	"bytes"
	"fmt"
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

func ExecuteCommandOnlyErrors(cmd *exec.Cmd) string {
	var stderrBuf bytes.Buffer
	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderrBuf
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Println()

		if stderrBuf.Len() > 0 {
			return stderrBuf.String()
		}
		// if out.Len() > 0 {
		// 	fmt.Println("Output:", out.String())
		// }

		os.Exit(1)
	}
	return ""
}
