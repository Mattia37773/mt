/*
Copyright © 2026 Matze
*/
package sys

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattia37773/mt/functions/ui"
)

func Error(err error) {
	errorText := strings.Split(err.Error(), "\n")
	fmt.Println(ui.Red("Error:"), strings.NewReplacer("[", "", "]", "").Replace(errorText[0]))

	if len(errorText) > 1 {
		fmt.Println(errorText[2])
		fmt.Println(errorText[3])
		fmt.Println()
	}

	fmt.Println(ui.GlowPink("Try mt --help for usage."))
	os.Exit(1)
}
