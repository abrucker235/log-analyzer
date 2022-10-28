package users

import (
	"log"

	"github.com/spf13/cobra"
)

var uniqueCMD = &cobra.Command{
	Use: "unique",
	Run: unique,
}

func unique(cmd *cobra.Command, args []string) {
	filename, err := cmd.Flags().GetString("file")
	if err != nil {
		log.Fatal("error getting filename value")
	}

	logs, err := parse(filename)
	if err != nil {
		log.Fatalf("error parsing log file: %v", err)
	}

	log.Printf("%d unique users accessed the server\n", len(logs))
}
