/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package main

import (
	"github.com/sniperkit/snk.fork.tarmak/cmd/wing/cmd"
)

var (
	version string = "dev"
	commit  string = "unknown"
	date    string = ""
)

func main() {
	cmd.Version.Version = version
	cmd.Version.Commit = commit
	cmd.Version.BuildDate = date
	cmd.Execute()
}
