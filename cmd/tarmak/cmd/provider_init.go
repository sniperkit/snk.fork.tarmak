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

var providerInitCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"initialise", "initialize"},
	Short:   "Initialize a provider",
	Run: func(cmd *cobra.Command, args []string) {
		globalFlags.Initialize = true
		t := tarmak.New(globalFlags)
		defer t.Cleanup()
		t.Must(t.CmdProviderInit())
	},
}

func init() {
	providerCmd.AddCommand(providerInitCmd)
}
