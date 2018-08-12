/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tarmak/pkg/wing"
)

var agentFlags = &wing.Flags{}

var agentCmd = &cobra.Command{
	Use: "agent",
	Run: func(cmd *cobra.Command, args []string) {
		w := wing.New(agentFlags)
		w.Must(w.Run(args))
	},
}

func init() {
	instanceName, err := os.Hostname()
	if err != nil {
		instanceName = ""
	}

	agentCmd.Flags().StringVar(&agentFlags.ClusterName, "cluster-name", "myenv-mycluster", "this specifies the cluster name [environment]-[cluster]")
	agentCmd.Flags().StringVar(&agentFlags.ServerURL, "server-url", "https://localhost:9443", "this specifies the URL to the wing server")
	agentCmd.Flags().StringVar(&agentFlags.ManifestURL, "manifest-url", "", "this specifies the URL where the puppet.tar.gz can be found")
	agentCmd.Flags().StringVar(&agentFlags.InstanceName, "instance-name", instanceName, "this specifies the instance's name")

	RootCmd.AddCommand(agentCmd)
}
