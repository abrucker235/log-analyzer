package users

import (
	"log"
	"time"

	"github.com/abrucker235/log-analyzer/types"
	"github.com/spf13/cobra"
)

var summaryCMD = &cobra.Command{
	Use: "summarize",
	Run: summary,
}

func init() {
	summaryCMD.Flags().String("name", "", "username to analyze")
	summaryCMD.Flags().String("operation", "", "operation filter")
	summaryCMD.Flags().String("date", "", "date filter")
}

func summary(cmd *cobra.Command, args []string) {
	filename, _ := cmd.Flags().GetString("file")
	name, _ := cmd.Flags().GetString("name")
	operation, _ := cmd.Flags().GetString("operation")
	date, _ := cmd.Flags().GetString("date")

	logs, err := parse(filename)
	if err != nil {
		log.Fatalf("error parsing log file: %v", err)
	}

	if name != "" {

		if operation != "" && date != "" {
			d, err := time.Parse(time.UnixDate, date)
			if err != nil {
				log.Fatalf("error parsing date filter: %v", err)
			}
			summarizeUserByOperationAndDate(logs[name][operation], d)
			return
		}

		if operation != "" {
			summarizeUserByOperation(logs[name][operation])
			return
		}

		if date != "" {
			d, err := time.Parse("2006-01-02", date)
			if err != nil {
				log.Fatalf("error parsing date filter: %v", err)
			}
			summarizeUserByDate(logs[name], d)
			return
		}

		summarizeUser(logs[name])
		return
	}

	summarize(logs)
}

func summarize(logs map[string]map[string][]types.Log) {
	for username, ls := range logs {
		log.Printf("user: %s uploads: %d downloads: %d\n", username, len(ls["upload"]), len(ls["download"]))
	}
}

func summarizeUser(logs map[string][]types.Log) {
	uploads := len(logs["upload"])
	downloads := len(logs["download"])
	log.Printf("uploads: %v downloads: %v\n", uploads, downloads)
}

func summarizeUserByDate(logs map[string][]types.Log, date time.Time) {
	for operation, ls := range logs {
		occurences := 0
		for _, l := range ls {
			lyear, lmonth, lday := l.Timestamp.UTC().Date()
			fyear, fmonth, fday := date.UTC().Date()
			if lyear == fyear && lmonth == fmonth && lday == fday {
				occurences++
			}
		}
		log.Printf("%s: %v", operation, occurences)
	}
}

func summarizeUserByOperation(logs []types.Log) {
	log.Println(len(logs))
}

func summarizeUserByOperationAndDate(logs []types.Log, date time.Time) {
	occurences := 0
	for _, l := range logs {
		lyear, lmonth, lday := l.Timestamp.UTC().Date()
		fyear, fmonth, fday := date.UTC().Date()
		if lyear == fyear && lmonth == fmonth && lday == fday {
			occurences++
		}
	}
}
