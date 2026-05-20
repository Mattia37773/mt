/*
Copyright © 2026 Matze
*/
package ui

import (
	"fmt"

	"strings"
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
	fmt.Println(Purple(topBorder))

	for _, line := range lines {
		spacesCount := boxWidth - len(line) - padding

		leftPart := Purple("│") + strings.Repeat(" ", padding)
		middlePart := Purple(line)
		rightPart := strings.Repeat(" ", spacesCount) + Purple("│")

		fmt.Printf("%s%s%s\n", leftPart, middlePart, rightPart)
	}

	bottomBorder := "└" + strings.Repeat("─", boxWidth) + "┘"
	fmt.Println(Purple(bottomBorder))
}
