/*
Copyright © 2026 Matze
*/
package ui

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

func FilePicker(startDir string) (string, error) {
	currentDir := startDir

	isArchive := func(name string) bool {
		ext := strings.ToLower(filepath.Ext(name))
		return ext == ".zip" || ext == ".gz" || ext == ".gzip"
	}

	for {
		entries, err := os.ReadDir(currentDir)
		if err != nil {
			return "", err
		}

		var options []huh.Option[string]

		if currentDir != "/" && currentDir != startDir {
			options = append(options, huh.NewOption(Lime(".."), ".."))
		}

		for _, entry := range entries {
			name := entry.Name()
			fullPath := filepath.Join(currentDir, name)

			if entry.IsDir() && !isArchive(name) {
				options = append(options, huh.NewOption(GlowPink(name), fullPath))
			} else {
				options = append(options, huh.NewOption(Purple(name), fullPath))
			}
		}

		var choice string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Choose a file").
					Options(options...).
					Value(&choice),
			),
		)

		if err := form.Run(); err != nil {
			return "", err
		}

		if choice == ".." {
			currentDir = filepath.Dir(currentDir)
			continue
		}

		info, err := os.Stat(choice)
		if err != nil {
			return "", err
		}

		if info.IsDir() && !isArchive(choice) {
			currentDir = choice
			continue
		}

		return choice, nil
	}
}
