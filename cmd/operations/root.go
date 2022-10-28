package operations

import "github.com/spf13/cobra"

var RootCMD = &cobra.Command{
	Use:   "operations",
	Short: "operation based summaries",
}

func init() {
	RootCMD.AddCommand(SummaryCMD)
}
