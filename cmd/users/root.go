package users

import "github.com/spf13/cobra"

var RootCMD = &cobra.Command{
	Use:   "users",
	Short: "user based summaries",
}

func init() {
	RootCMD.AddCommand(uniqueCMD)
	RootCMD.AddCommand(summaryCMD)
}
