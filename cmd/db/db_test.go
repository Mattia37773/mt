package db

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/tests/base"
	"github.com/stretchr/testify/assert"
)

// change to full test suite

func TestDbShellMysqCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := shellGen(rootCmd.OutOrStdout(), []string{})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "sh", "-c", "mysql -u" + config.ProjectConfig.DB.User + " -p" + config.ProjectConfig.DB.Password + " "}, shellCommand)
}

func TestExportMysqlExportCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := exportGen(rootCmd.OutOrStdout(), []string{})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "sh", "-c", "mysqldump -u" + config.ProjectConfig.DB.User + " -p" + config.ProjectConfig.DB.Password + " " + config.ProjectConfig.DB.Name + " > /tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype + " "}, shellCommand)
}

func TestExportMysqlCopyCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := copyExportGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"docker", "cp", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName + ":/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype, "./"}, shellCommand)
}

func TestExportMysqlRemoveCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := removeExportGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"docker", "exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype}, shellCommand)
}

func TestImportMysqlImportCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := importGen(rootCmd.OutOrStdout(), "snippets.sql")
	assert.Equal(t, []string{"docker", "exec", "-i", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "bash", "-pc", "mysql -u" + config.ProjectConfig.DB.User + " -p" + config.ProjectConfig.DB.Password + " " + config.ProjectConfig.DB.Name + " < /tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype + ""}, shellCommand)
}

func TestImportMysqCopyCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := copyImportGen(rootCmd.OutOrStdout(), "snippy.sql")
	assert.Equal(t, []string{"docker", "cp", "./" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype, config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName + ":/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype + ""}, shellCommand)
}

func TestImportMysqRemoveCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := removeImportGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"docker", "exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype + ""}, shellCommand)
}

func TestDbShellMongoCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := shellGen(rootCmd.OutOrStdout(), []string{})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "sh", "-c", "mongosh --username " + config.ProjectConfig.DB.User + " --password " + config.ProjectConfig.DB.Password + " "}, shellCommand)
}

func TestExportMongoExportCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := exportGen(rootCmd.OutOrStdout(), []string{})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "sh", "-c", "mongodump --db=" + config.ProjectConfig.DB.Name + " --username=" + config.ProjectConfig.DB.User + " --password=" + config.ProjectConfig.DB.Password + " --authenticationDatabase=admin --out=/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype + " --gzip "}, shellCommand)
}

func TestExportMongoCopyCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := copyExportGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"docker", "cp", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName + ":/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype, "./"}, shellCommand)
}

func TestExportMongoRemoveCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := removeExportGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"docker", "exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype}, shellCommand)
}

func TestImportMongoImportCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := importGen(rootCmd.OutOrStdout(), "snippets.gzip")
	assert.Equal(t, []string{"docker", "exec", "-i", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "bash", "-pc", "mongorestore --nsFrom=" + config.ProjectConfig.DB.Name + ".* " + "--nsTo=" + config.ProjectConfig.DB.Name + ".* " + "--dir=/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype + " --gzip --username=" + config.ProjectConfig.DB.User + " --password=" + config.ProjectConfig.DB.Password}, shellCommand)
}

func TestImportMongoCopyCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := copyImportGen(rootCmd.OutOrStdout(), "mongotest.gzip")
	assert.Equal(t, []string{"docker", "cp", "./" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype, config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName + ":/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype}, shellCommand)
}

func TestImportMongoRemoveCommandGenlSuccess(t *testing.T) {
	base.ChangeDirToMongo(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := removeImportGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"docker", "exec", config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/" + config.ProjectConfig.ProjectName + config.GetDbConfig().Filetype}, shellCommand)
}
