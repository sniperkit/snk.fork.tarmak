/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak"
)

var clusterDebugTerraformShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Prepares a Terraform container and executes a shell in this container",
	Run: func(cmd *cobra.Command, args []string) {
		t := tarmak.New(globalFlags)
		defer t.Cleanup()
		defer plugin.CleanupClients()
		t.Must(t.CmdTerraformShell(args))
	},
}

func init() {
	clusterDebugTerraformCmd.AddCommand(clusterDebugTerraformShellCmd)
}
