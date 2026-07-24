package basecmd

import (
	"fmt"
	"io"
	"os"

	"github.com/mattia37773/mt/helper/docker"
	"github.com/mattia37773/mt/ui/text"
)

func CheckContainerExits(out io.Writer, name string) {
	exists, err := docker.ContainerExists(name)
	if err != nil {
		fmt.Fprintf(out, text.Red("Error: Somehting is wrong with the contaner %s \n"), name)
		fmt.Fprintf(out, "%s \n", err)
		os.Exit(1)
	}
	if exists == false {
		fmt.Fprintf(out, text.Red("Error: The Container: %s doesn't exist\n"), name)
		fmt.Fprintf(out, "Did you forget to run mt stack start? \n")
		os.Exit(1)
	}
}
