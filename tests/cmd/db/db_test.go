package db_test

import (
	"bytes"
	"os"
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

type dBTestConfig struct {
	ProjectPath string // TODO lowercase
}

type testCase struct {
	name   string
	config dBTestConfig
}

func TestDb(t *testing.T) {
	testCases := []testCase{
		{
			name: "Mysql",
			config: dBTestConfig{
				ProjectPath: "mysql",
			},
		},
		{
			name: "MongoDb",
			config: dBTestConfig{
				ProjectPath: "mongo",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" Start Stack", func(t *testing.T) {
			testStartDbStack(t, tc)
			time.Sleep(15 * time.Second)
		})

		t.Run(tc.name+" Import DB", func(t *testing.T) {
			testImport(t, tc)
		})

		t.Run(tc.name+" Export DB", func(t *testing.T) {
			testExport(t, tc)
		})

		//
		// negative tests
		//

		t.Run(tc.name+" Import Without File", func(t *testing.T) {
			testImportWithoutFile(t, tc)
		})

		t.Run(tc.name+" Import With Not Existing File", func(t *testing.T) {
			testImportFileNotExisting(t, tc)
		})

		//
		// destroy
		//

		t.Run(tc.name+" Destroy DB", func(t *testing.T) {
			testDestroy(t, tc)
		})

		//
		// without container
		//
		t.Run(tc.name+" Import Without Running Container", func(t *testing.T) {
			testImportContainerNotRunning(t, tc)
		})

		t.Run(tc.name+" Export Without Running Container", func(t *testing.T) {
			testExportContainerNotRunning(t, tc)
		})

		t.Run(tc.name+" Shell Without Running Container", func(t *testing.T) {
			testShellContainerNotRunning(t, tc)
		})
	}
}

func testStartDbStack(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "start"})

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.DB.ContainerName,
	}

	base.WaitForContainersRunning(t, expectedContainers, 3*time.Minute)

	for _, container := range expectedContainers {
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

func testImport(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	db, _ := config.GetDbConfig()
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "backup" + db.Filetype})
	errExec := base.ExecuteCommand(rootCmd)
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Imported the Database dump into "+config.ProjectConfig.ProjectName)
}

func testExport(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	db, _ := config.GetDbConfig()
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "export"})
	errExec := base.ExecuteCommand(rootCmd)
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Created Dump Successfully")

	_, err := os.Stat(config.ProjectConfig.ProjectName + db.Filetype)
	assert.NoError(t, err)
}

func testDestroy(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	db, _ := config.GetDbConfig()

	err := os.RemoveAll(config.ProjectConfig.ProjectName + db.Filetype)
	assert.Nil(t, err)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})
}

//
// negative tests
//

func testImportWithoutFile(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"flag needs an argument: 'f' in -f",
	)
}

func testImportFileNotExisting(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	db, _ := config.GetDbConfig()
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", "file" + db.Filetype})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The to be imported file doesn't exist: file"+db.Filetype+" \nPlease choose a other file",
	)
}

//
// tests without running container
//

func testImportContainerNotRunning(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "import", "-f", ".mt.yaml"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testExportContainerNotRunning(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "export"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testShellContainerNotRunning(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "shell"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

// func testMongoDataExists(t *testing.T) {
// 	base.ChangeDirToMongo(t)

// 	cmd := exec.Command(
// 		"sh",
// 		"-c",
// 		`docker exec `+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+` mongosh \
// -u `+config.ProjectConfig.DB.User+` \
// -p `+config.ProjectConfig.DB.Password+` \
// --authenticationDatabase admin \
// --quiet \
// --eval '
// const db = db.getSiblingDB("appDb");

// print(JSON.stringify({
// 	users: db.users.countDocuments(),
// 	products: db.products.countDocuments(),
// 	orders: db.orders.countDocuments(),
// 	payments: db.payments.countDocuments(),
// 	sessions: db.sessions.countDocuments(),
// 	notifications: db.notifications.countDocuments(),
// 	audit_logs: db.audit_logs.countDocuments()
// }));
// '`,
// 	)

// 	out, err := cmd.CombinedOutput()
// 	assert.NoError(t, err, string(out))

// 	assert.NotEmpty(t, out, "The Database returned nothing")

// 	var result map[string]int

// 	err = json.Unmarshal(out, &result)
// 	assert.NoError(t, err, string(out))

// 	assert.Greater(t, result["users"], 0)
// 	assert.Greater(t, result["products"], 0)
// 	assert.Greater(t, result["orders"], 0)
// 	assert.Greater(t, result["payments"], 0)
// 	assert.Greater(t, result["sessions"], 0)
// 	assert.Greater(t, result["notifications"], 0)
// 	assert.Greater(t, result["audit_logs"], 0)
// }

// func testMysqlDataExists(t *testing.T) {
// 	base.ChangeDirToMysql(t)

// 	cmd := exec.Command(
// 		"sh",
// 		"-c",
// 		`docker exec `+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.DB.ContainerName+` mysql -u`+config.ProjectConfig.DB.User+` -p`+config.ProjectConfig.DB.Password+` `+config.ProjectConfig.DB.Name+` -N -e "
// SELECT
// 	(SELECT COUNT(*) FROM categories),
// 	(SELECT COUNT(*) FROM customers),
// 	(SELECT COUNT(*) FROM products),
// 	(SELECT COUNT(*) FROM orders),
// 	(SELECT COUNT(*) FROM order_items),
// 	(SELECT COUNT(*) FROM roles),
// 	(SELECT COUNT(*) FROM users);
// "`,
// 	)

// 	out, err := cmd.CombinedOutput()
// 	assert.NoError(t, err, string(out))

// 	assert.NotEmpty(t, out, "Die Datenbank liefert keine Ausgabe")
// }
