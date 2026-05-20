/*
Copyright © 2026 Matze
*/
package main

import (
	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/single"
	_ "github.com/mattia37773/mt/cmd/stack"
	_ "github.com/mattia37773/mt/config"
	_ "github.com/mattia37773/mt/functions/docker"
	_ "github.com/mattia37773/mt/functions/env"
	_ "github.com/mattia37773/mt/functions/file"
	_ "github.com/mattia37773/mt/functions/shell"
	_ "github.com/mattia37773/mt/functions/sys"
	_ "github.com/mattia37773/mt/functions/ui"
	_ "github.com/mattia37773/mt/functions/yaml"
)

func main() {
	cmd.Execute()
}
