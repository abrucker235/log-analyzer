package cmd

import (
	"github.com/abrucker235/log-analyzer/cmd/operations"
	"github.com/abrucker235/log-analyzer/cmd/users"
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Use: "log-analyzer",
}

func init() {
	RootCMD.PersistentFlags().String("file", "", "Log file to analyze")
	RootCMD.AddCommand(users.RootCMD)
	RootCMD.AddCommand(operations.RootCMD)
}
