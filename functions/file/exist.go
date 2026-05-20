/*
Copyright © 2026 Matze
*/
package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/mattia37773/mt/functions/ui"
)

func ExistsMany(files []string) {
	for _, file := range files {
		if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
			fmt.Printf(ui.Red("Error: %s file does not exist\n"), file)
			os.Exit(1)
		}
	}
}

func Exists(file string) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Printf(ui.Red("Error: %s file does not exist\n"), file)
		os.Exit(1)
	}
}
