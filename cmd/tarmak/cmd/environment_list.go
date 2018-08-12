/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak"
	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak/utils"
)

var environmentListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a list of environments",
	Run: func(cmd *cobra.Command, args []string) {
		t := tarmak.New(globalFlags)
		defer t.Cleanup()
		varMaps := make([]map[string]string, 0)
		for _, env := range t.Environments() {
			varMaps = append(varMaps, env.Parameters())
		}
		utils.ListParameters(os.Stdout, []string{"name", "provider", "location"}, varMaps)
	},
}

func init() {
	environmentCmd.AddCommand(environmentListCmd)
}
