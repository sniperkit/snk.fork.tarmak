/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package cmd

import (
	"flag"
	"os"
	"runtime"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/util/logs"

	"github.com/sniperkit/snk.fork.tarmak/pkg/wing/server"
)

func init() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	stopCh := genericapiserver.SetupSignalHandler()
	serverCmd := server.NewCommandStartWingServer(os.Stdout, os.Stderr, stopCh)
	serverCmd.Use = "server"
	serverCmd.Flags().AddGoFlagSet(flag.CommandLine)
	RootCmd.AddCommand(serverCmd)
}
