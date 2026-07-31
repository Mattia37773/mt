package basecmd

import (
	"fmt"
	"io"

	"github.com/mattia37773/mt/helper/docker"
	"github.com/mattia37773/mt/ui/text"
)

func CheckContainerExits(out io.Writer, name string) error {
	exists, err := docker.ContainerExists(name)
	if err != nil {
		fmt.Fprintf(out, text.Red("Error: Somehting is wrong with the contaner %s \n"), name)
		return fmt.Errorf("Somehting is wrong with the contaner %s \n, %s", name, err)
	}
	if exists == false {
		return fmt.Errorf("The Container: %s doesn't exist\nDid you forget to run mt stack start?", name)
	}
	return nil
}
