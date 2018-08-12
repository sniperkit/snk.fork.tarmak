/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak"
)

var providerValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate provider(s) used by current cluster",
	Run: func(cmd *cobra.Command, args []string) {
		t := tarmak.New(globalFlags)
		defer t.Cleanup()
		t.Must(t.Environment().Provider().Validate())
	},
}

func init() {
	providerCmd.AddCommand(providerValidateCmd)
}
