/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak"
	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak/utils"
)

// clusterDestroyCmd handles `tarmak clusters destroy`
var clusterDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the current cluster",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		store := &globalFlags.Cluster.Destroy
		if store.DryRun {
			return errors.New("dry run is not yet supported")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		t := tarmak.New(globalFlags)
		defer t.Cleanup()
		utils.WaitOrCancel(
			func(ctx context.Context) error {
				return t.CmdTerraformDestroy(args, ctx)
			},
		)
	},
}

func init() {
	clusterDestroyFlags(clusterDestroyCmd.PersistentFlags())
	clusterCmd.AddCommand(clusterDestroyCmd)
}
