/*
Copyright © 2026 Matze
*/
package base

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mattia37773/mt/config"
)

func ChangeDirToDefault(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "default"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToMongo(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "mongo"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToMysql(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "mysql"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToSymfony(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "symfony"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToDefaultConfig(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "defaultConfig"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

// error tests

func ChangeDirToNotExistingEnvFile(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "error", "notExistingEnvFile"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToNotExistingEnvVar(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "error", "notExistingEnvVar"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToNotExistingComposeile(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "error", "notExistingComposeFile"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToNoProjectName(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "error", "noProjectName"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToNoComposePath(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "error", "noComposePath"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

func ChangeDirToNoConfigFilWithError(t *testing.T) {
	dir := rootDir([]string{"tests", "projects", "error", "configFileWithError"})
	t.Chdir(dir)
	config.ParseConfigFile()
}

// get the project root directory
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
