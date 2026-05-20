/*
Copyright © 2026 Matze
*/
package sys

import (
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

func VersionStyle(cmd *cobra.Command) {
	cmd.SetVersionTemplate(
		ui.GlowPink("{{.Name}}") +
			" " +
			ui.GlowPurple("version {{.Version}}") +
			"\n",
	)
}
