package shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/mattia37773/mt/functions/ui"
)

func ExecuteCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
	}

	sigChan := make(chan os.Signal, 1)

	go func() {
		<-sigChan
		if cmd.Process != nil {
			cmd.Process.Signal(os.Interrupt)
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Printf(ui.Red("Error: %s\n"), err)
		os.Exit(1)
	}

}

func ExecuteCommandReturn(cmd *exec.Cmd) (bool, error) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
	}

	sigChan := make(chan os.Signal, 1)

	go func() {
		<-sigChan
		if cmd.Process != nil {
			cmd.Process.Signal(os.Interrupt)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return false, err
	}

	return true, nil
}

func ExecuteCommandOnlyErrors(cmd *exec.Cmd) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println()

		if stderr.Len() > 0 {
			fmt.Println("Error:", stderr.String())
		}
		if out.Len() > 0 {
			fmt.Println("Output:", out.String())
		}

		os.Exit(1)
	}
}

func ExecuteCommandOnlyErrorsReturn(cmd *exec.Cmd) error {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println()

		if stderr.Len() > 0 {
			fmt.Println("Error:", stderr.String())
		}
		if out.Len() > 0 {
			fmt.Println("Output:", out.String())
		}

		return err
	}

	return nil
}
