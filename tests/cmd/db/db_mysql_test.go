/*
Copyright © 2026 Matze
*/
package db_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/docker"
	base "github.com/mattia37773/mt/tests/basetest"

	"github.com/stretchr/testify/assert"
)

func TestDbMysql(t *testing.T) {
	t.Run("Start", func(t *testing.T) {
		testStartMysqlStack(t)
		time.Sleep(30 * time.Second)
	})

	t.Run("Db import mysql", func(t *testing.T) {
		testImportMysql(t)
	})

	t.Run("Mysql data exist after import", func(t *testing.T) {
		testMysqlDataExists(t)
	})

	t.Run("Db export Mysql", func(t *testing.T) {
		testExportMysqlDb(t)
	})

	// * negative tests

	t.Run("Mysql import File Doesnt exist", func(t *testing.T) {
		testMysqlImportNoFileGiven(t)
	})

	t.Run("Mysql import without file", func(t *testing.T) {
		testMysqlImportFileDoesntExist(t)
	})

	// * destroy everything

	t.Run("Mysql stack destroy", func(t *testing.T) {
		testDestroyMysql(t)
	})

	// * negative test with destroyed container
	t.Run("Mysql import contianer not running", func(t *testing.T) {
		testMysqlImportContainerNotRunning(t)
	})

	t.Run("Mysql export contianer not running", func(t *testing.T) {
		testMysqlExportContainerNotRunning(t)
	})
	t.Run("Mysql shell contianer not running", func(t *testing.T) {
		testMysqlShellContainerNotRunning(t)
	})
}

func testStartMysqlStack(t *testing.T) {
	base.ChangeDirToMysql(t)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "start"})

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-db",
	}

	for _, container := range expectedContainers {
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

func testExportMysqlDb(t *testing.T) {
	base.ChangeDirToMysql(t)
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

func testImportMysql(t *testing.T) {
	base.ChangeDirToMysql(t)
	db := config.GetDbConfig()
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "backup" + db.Filetype})
	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Imported the Database dump into "+config.ProjectConfig.ProjectName)
}

func testMysqlDataExists(t *testing.T) {
	base.ChangeDirToMysql(t)

	cmd := exec.Command(
		"sh",
		"-c",
		`docker exec `+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+` mysql -u`+config.ProjectConfig.DB.User+` -p`+config.ProjectConfig.DB.Password+` `+config.ProjectConfig.DB.Name+` -N -e "
SELECT
	(SELECT COUNT(*) FROM categories),
	(SELECT COUNT(*) FROM customers),
	(SELECT COUNT(*) FROM products),
	(SELECT COUNT(*) FROM orders),
	(SELECT COUNT(*) FROM order_items),
	(SELECT COUNT(*) FROM roles),
	(SELECT COUNT(*) FROM users);
"`,
	)

	out, err := cmd.CombinedOutput()
	assert.NoError(t, err, string(out))

	// Erwartetes Ergebnis (eine Zeile mit 7 Werten)
	assert.NotEmpty(t, out, "Die Datenbank liefert keine Ausgabe")
}

func testDestroyMysql(t *testing.T) {
	base.ChangeDirToMysql(t)
	db := config.GetDbConfig()

	// delete vendor directory
	vendorErr := os.RemoveAll(config.ProjectConfig.ProjectName + db.Filetype)
	assert.Nil(t, vendorErr)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})
}

// negative tests

func testMysqlImportNoFileGiven(t *testing.T) {
	base.ChangeDirToMysql(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "file.sql"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The to be imported file doesn't exist: file.sql \nPlease choose a other file",
	)
}

func testMysqlImportFileDoesntExist(t *testing.T) {
	base.ChangeDirToMysql(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "test.sql"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The to be imported file doesn't exist: test.sql",
	)
}

func testMysqlImportContainerNotRunning(t *testing.T) {
	base.ChangeDirToMysql(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "backup.sql"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: snippy-db doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testMysqlExportContainerNotRunning(t *testing.T) {
	base.ChangeDirToMysql(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "export"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: snippy-db doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testMysqlShellContainerNotRunning(t *testing.T) {
	base.ChangeDirToMysql(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "shell"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: snippy-db doesn't exist",
		"Did you forget to run mt stack start?",
	)
}
