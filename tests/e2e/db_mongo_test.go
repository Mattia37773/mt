/*
Copyright © 2026 Matze
*/
package e2e

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/mattia37773/mt/helper/docker"

	"github.com/stretchr/testify/assert"
)

func TestDbMongo(t *testing.T) {
	t.Run("Start", func(t *testing.T) {
		testStartMongoStack(t)
		time.Sleep(30 * time.Second)
	})

	t.Run("Mongo data exist after container creation", func(t *testing.T) {
		testMongoDataExists(t)
	})

	t.Run("Db export Mysql", func(t *testing.T) {
		testExportMongoDb(t)
	})

	t.Run("Mongo import", func(t *testing.T) {
		testMongoImport(t)
	})

	t.Run("Mongo data exist after import", func(t *testing.T) {
		testMongoDataExists(t)
	})

	// negative tests

	t.Run("Mysql import File Doesnt exist", func(t *testing.T) {
		testMongoImportNoFileGiven(t)
	})

	t.Run("Mongo import without file", func(t *testing.T) {
		testMongoImportFileDoesntExist(t)
	})

	// destroy everything
	t.Run("Mongo stack destroy", func(t *testing.T) {
		testDestroyMongo(t)
	})

	// egative test with destroyed container
	t.Run("Mongo import contianer not running", func(t *testing.T) {
		testMongoImportContainerNotRunning(t)
	})

	t.Run("Mongo export contianer not running", func(t *testing.T) {
		testMongoExportContainerNotRunning(t)
	})
	t.Run("Mongo shell contianer not running", func(t *testing.T) {
		testMongoShellContainerNotRunning(t)
	})
}

func testStartMongoStack(t *testing.T) {
	base.ChangeDirToMongo(t)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "start"})

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName,
	}

	for _, container := range expectedContainers {
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

func testExportMongoDb(t *testing.T) {
	base.ChangeDirToMongo(t)
	db := config.GetDbConfig()
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// export db
	rootCmd.SetArgs([]string{"db", "export"})
	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Created Dump Successfully")

	_, err := os.Stat(config.ProjectConfig.ProjectName + db.Filetype)
	assert.NoError(t, err)
}

func testMongoImport(t *testing.T) {
	base.ChangeDirToMongo(t)
	db := config.GetDbConfig()
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", config.ProjectConfig.ProjectName + db.Filetype})
	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Imported the Database dump into "+config.ProjectConfig.ProjectName)
}

func testMongoDataExists(t *testing.T) {
	base.ChangeDirToMongo(t)

	cmd := exec.Command(
		"sh",
		"-c",
		`docker exec `+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+` mongosh \
-u `+config.ProjectConfig.DB.User+` \
-p `+config.ProjectConfig.DB.Password+` \
--authenticationDatabase admin \
--quiet \
--eval '
const db = db.getSiblingDB("appDb");

print(JSON.stringify({
	users: db.users.countDocuments(),
	products: db.products.countDocuments(),
	orders: db.orders.countDocuments(),
	payments: db.payments.countDocuments(),
	sessions: db.sessions.countDocuments(),
	notifications: db.notifications.countDocuments(),
	audit_logs: db.audit_logs.countDocuments()
}));
'`,
	)

	out, err := cmd.CombinedOutput()
	assert.NoError(t, err, string(out))

	assert.NotEmpty(t, out, "The Database returned nothing")

	var result map[string]int

	err = json.Unmarshal(out, &result)
	assert.NoError(t, err, string(out))

	assert.Greater(t, result["users"], 0)
	assert.Greater(t, result["products"], 0)
	assert.Greater(t, result["orders"], 0)
	assert.Greater(t, result["payments"], 0)
	assert.Greater(t, result["sessions"], 0)
	assert.Greater(t, result["notifications"], 0)
	assert.Greater(t, result["audit_logs"], 0)
}

func testDestroyMongo(t *testing.T) {
	base.ChangeDirToMongo(t)
	db := config.GetDbConfig()

	err := os.RemoveAll(config.ProjectConfig.ProjectName + db.Filetype)
	assert.Nil(t, err)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})
}

// // negative tests

func testMongoImportNoFileGiven(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "file.gzip"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The to be imported file doesn't exist: file.gzip \nPlease choose a other file",
	)
}

func testMongoImportFileDoesntExist(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "test.gzip"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The to be imported file doesn't exist: test.gzip",
	)
}

func testMongoImportContainerNotRunning(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", ".env"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testMongoExportContainerNotRunning(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "export"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testMongoShellContainerNotRunning(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "shell"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}
