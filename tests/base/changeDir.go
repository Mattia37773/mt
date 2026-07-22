package base

import (
	"os"
	"path/filepath"
	"testing"
)

func ChangeDirToDefault(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "default"})
	t.Chdir(dir)
}

func ChangeDirToSymfony(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "symfony"})
	t.Chdir(dir)
}

func rootDir(projectFilepaths []string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			targetDir := filepath.Join(append([]string{dir}, projectFilepaths...)...)
			return targetDir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			panic("Project root couldn't been found")
		}
		dir = parent
	}
}
