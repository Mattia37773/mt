package single

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/stretchr/testify/assert"
)

func TestNotExistingEnvFile(t *testing.T) {
	base.ChangeDirToNotExistingEnvFile(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "start"})
	errExec := rootCmd.Execute()
	assert.Error(t, errExec)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The defined env file in .mt.yaml doesnt exist: .env",
	)
}

func TestNotExistingEnvVar(t *testing.T) {
	base.ChangeDirToNotExistingEnvVar(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "start"})
	errExec := rootCmd.Execute()
	assert.Error(t, errExec)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"required environment variable 'DB_PASSWORD' is not set",
	)
}

func TestNotExistingComposeFile(t *testing.T) {
	base.ChangeDirToNotExistingComposeile(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "start"})
	errExec := rootCmd.Execute()
	assert.Error(t, errExec)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The docker compose file doesnt exist: "+config.ProjectConfig.Paths.DockerCompose,
	)
}

// func TestNoProjectName(t *testing.T) {
// 	base.ChangeDirToNoProjectName(t)
// 	rootCmd := cmd.RootCmd

// 	buf := new(bytes.Buffer)
// 	rootCmd.SetOut(buf)
// 	rootCmd.SetErr(buf)

// 	rootCmd.SetArgs([]string{"stack", "start"})
// 	errExec := rootCmd.Execute()
// 	assert.NoError(t, errExec)

// 	assert.Error(t, errExec)
// 	assert.Contains(t, errExec.Error(),
// 		"The docker compose file doesnt exist: "+config.ProjectConfig.Paths.DockerCompose,
// 	)
// }

// func TestNoComposePath(t *testing.T) {
// 	base.ChangeDirToNoComposePath(t)
// 	rootCmd := cmd.RootCmd

// 	buf := new(bytes.Buffer)
// 	rootCmd.SetOut(buf)
// 	rootCmd.SetErr(buf)

// 	rootCmd.SetArgs([]string{"stack", "start"})
// 	errExec := rootCmd.Execute()
// 	assert.NoError(t, errExec)

// 	assert.Error(t, errExec)
// 	assert.Contains(t, errExec.Error(),
// 		"The docker compose file doesnt exist: "+config.ProjectConfig.Paths.DockerCompose,
// 	)
// }

func TestNotExistingConfigWithError(t *testing.T) {
	base.ChangeDirToNoConfigFilWithError(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "start"})
	errExec := rootCmd.Execute()
	assert.Error(t, errExec)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"An error occurred while parsing '.mt.yaml':",
	)
}
