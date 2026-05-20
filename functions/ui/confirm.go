/*
Copyright © 2026 Matze
*/
package ui

import (
	"github.com/charmbracelet/huh"
)

func Confirm(title string) (bool, error) {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm),
		))
	err := form.Run()
	return confirm, err
}
