/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak"
	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak/utils"
)

var clusterPlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plan changes on the currently configured cluster",
	Run: func(cmd *cobra.Command, args []string) {
		t := tarmak.New(globalFlags)
		defer t.Cleanup()
		utils.WaitOrCancel(
			func(ctx context.Context) error {
				return t.CmdTerraformPlan(args, ctx)
			},
			2,
		)
	},
}

func init() {
	//clusterPlanFlags(clusterPlanCmd.PersistentFlags())
	clusterCmd.AddCommand(clusterPlanCmd)
}
