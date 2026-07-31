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
)

func main() {
	cmd.Execute()
}
