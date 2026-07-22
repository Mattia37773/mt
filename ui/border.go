/*
Copyright © 2026 Matze
*/
package ui

import (
	"fmt"

	"strings"

	"github.com/mattia37773/mt/ui/text"
)

func Border(lines []string) {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	padding := 2
	boxWidth := maxLen + padding*2

	topBorder := "┌" + strings.Repeat("─", boxWidth) + "┐"
	fmt.Println(text.Purple(topBorder))

	for _, line := range lines {
		spacesCount := boxWidth - len(line) - padding

		leftPart := text.Purple("│") + strings.Repeat(" ", padding)
		middlePart := text.Purple(line)
		rightPart := strings.Repeat(" ", spacesCount) + text.Purple("│")

		fmt.Printf("%s%s%s\n", leftPart, middlePart, rightPart)
	}

	bottomBorder := "└" + strings.Repeat("─", boxWidth) + "┘"
	fmt.Println(text.Purple(bottomBorder))
}
