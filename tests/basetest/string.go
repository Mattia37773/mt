/*
Copyright © 2026 Matze
*/
package base

import (
	"strings"
)

func GetLineFromStringReversed(text string, lineNumber int) string {
	trimmed := strings.TrimSpace(text)
	lines := strings.Split(trimmed, "\n")
	line := lines[len(lines)-lineNumber]

	return line
}
