package e2e

// change to full test suite
import (
	"bytes"
	"os"
	"testing"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/stretchr/testify/assert"
)

func TestSingles(t *testing.T) {
	// t.Run("Start", func(t *testing.T) {
	// 	// this is from other package
	// 	testStartCmd(t)
	// 	time.Sleep(30 * time.Second)
	// })

	// t.Run("Composer install", func(t *testing.T) {
	// 	testComposerInstall(t)
	// })

	// t.Run("Composer help", func(t *testing.T) {
	// 	testComposerHelp(t)
	// })

	// t.Run("Run with arguments", func(t *testing.T) {
	// 	testRunWithArguments(t)
	// })

	// t.Run("Run doctrine migrations", func(t *testing.T) {
	// 	testRunDoctrineMigration(t)
	// })

	// t.Run("Run with --help argument", func(t *testing.T) {
	// 	testRunWithHelp(t)
	// })

	// t.Run("Symfony console load fixtures", func(t *testing.T) {
	// 	testConsoleLoadFixtures(t)
	// })

	// t.Run("Destory Everything", func(t *testing.T) {
	// 	testDestroy(t)
	// })

	//*
	//// * Negative tests
	//*

	// t.Run("Shell without container argument", func(t *testing.T) {
	// 	testShellWithoutContainerArgument(t)
	// })

	// t.Run("Shell without flag argument", func(t *testing.T) {
	// 	testShellWithoutFlagArgument(t)
	// })

	//*
	//// * not running tests
	//*

	// t.Run("Console without container", func(t *testing.T) {
	// 	testConsoleContainerNotRunning(t)
	// })

	// t.Run("Craft without container", func(t *testing.T) {
	// 	testCraftContainerNotRunning(t)
	// })

	// t.Run("Composer without container", func(t *testing.T) {
	// 	testComposerContainerNotRunning(t)
	// })

	// t.Run("Run without container", func(t *testing.T) {
	// 	testRunContainerNotRunning(t)
	// })
}

func testComposerInstall(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"composer", "install"})

	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Executing script assets:install public [OK]")

	_, errVendor := os.Stat("./vendor")
	assert.NoError(t, errVendor)
}

func testComposerHelp(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"composer", "--help"})

	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "/usr/local/bin/composer list --raw")
}

func testRunWithArguments(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"run", "ls", "-la"})

	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, ".")
	assert.Contains(t, cleanOutput, "..")
	assert.Contains(t, cleanOutput, ".devcontainer")
	assert.Contains(t, cleanOutput, ".editorconfig")
	assert.Contains(t, cleanOutput, ".env")
	assert.Contains(t, cleanOutput, ".env.dev")
	assert.Contains(t, cleanOutput, ".env.prod")
	assert.Contains(t, cleanOutput, ".gitignore")
	assert.Contains(t, cleanOutput, ".mt.yaml")
	assert.Contains(t, cleanOutput, ".php-cs-fixer.dist.php")
	assert.Contains(t, cleanOutput, ".phpactor.json")
	assert.Contains(t, cleanOutput, ".vscode")
	assert.Contains(t, cleanOutput, "README.md")
	assert.Contains(t, cleanOutput, "assets")
	assert.Contains(t, cleanOutput, "bin")
	assert.Contains(t, cleanOutput, "composer.json")
	assert.Contains(t, cleanOutput, "composer.lock")
	assert.Contains(t, cleanOutput, "config")
	assert.Contains(t, cleanOutput, "docker")
	assert.Contains(t, cleanOutput, "info.md")
	assert.Contains(t, cleanOutput, "migrations")
	assert.Contains(t, cleanOutput, "node_modules")
	assert.Contains(t, cleanOutput, "package-lock.json")
	assert.Contains(t, cleanOutput, "package.json")
	assert.Contains(t, cleanOutput, "public")
	assert.Contains(t, cleanOutput, "src")
	assert.Contains(t, cleanOutput, "symfony.lock")
	assert.Contains(t, cleanOutput, "templates")
	assert.Contains(t, cleanOutput, "var")
	assert.Contains(t, cleanOutput, "vendor")
	assert.Contains(t, cleanOutput, "vite.config.js")
	assert.Contains(t, cleanOutput, "")
}

func testRunDoctrineMigration(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Destory db
	rootCmd.SetArgs([]string{"run", "php", "bin/console", "doctrine:database:drop", "--force"})
	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Dropped database `appDb` for connection named default")

	// create new db
	buf.Reset()
	rootCmd.SetArgs([]string{"run", "php", "bin/console", "doctrine:database:create"})
	errExec = rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput = base.StripANSI(buf.String())

	lastLine = base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "Created database `appDb` for connection named default")

	// apply migration
	buf.Reset()
	rootCmd.SetArgs([]string{"run", "php", "bin/console", "doctrine:migrations:migrate", "-n"})
	errExec = rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput = base.StripANSI(buf.String())

	lastLine = base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "[OK] Successfully migrated to version: DoctrineMigrations\\Version20260620211915")
}

func testRunWithHelp(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"run", "ls", "--help"})
	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "or available locally via: info '(coreutils) ls invocation'")
}

func testConsoleLoadFixtures(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "console", "doctrine:fixtures:load", "-n"})
	errExec := rootCmd.Execute()
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "   > loading App\\DataFixtures\\AppFixtures")
}

func testDestroy(t *testing.T) {
	base.ChangeDirToSymfony(t)
	db := config.GetDbConfig()

	// delete vendor directory
	vendorErr := os.RemoveAll("vendor")
	assert.Nil(t, vendorErr)

	dbDumpErr := os.RemoveAll(config.ProjectConfig.ProjectName + db.Filetype)
	assert.Nil(t, dbDumpErr)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})
}

//*
//// * Negative tests
//*

func testShellWithoutContainerArgument(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"shell"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"Missing argument CONTAINER",
	)
}

func testShellWithoutFlagArgument(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"shell", "fpm", "-u"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"flag needs an argument: 'u' in -u",
	)
}

//*
//// * not running tests
//*

func testConsoleContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "console", "symfony > laravel"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Backend.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testCraftContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "craft", "5 is worse then 4"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Backend.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testComposerContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"composer", "blabla"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Backend.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testRunContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"run", "i hate testing"})
	errExec := rootCmd.Execute()

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Main.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}
