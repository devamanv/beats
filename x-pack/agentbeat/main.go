// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"os"
	_ "time/tzdata" // for timezone handling

	agentbeatcmd "github.com/elastic/beats/v7/x-pack/agentbeat/cmd"
)

func main() {
	rootCmd := agentbeatcmd.AgentBeat()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
