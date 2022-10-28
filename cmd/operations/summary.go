package operations

import (
	"log"
	"strconv"

	"github.com/abrucker235/log-analyzer/types"
	"github.com/spf13/cobra"
)

var SummaryCMD = &cobra.Command{
	Use: "summarize",
	Run: summarize,
}

func init() {
	SummaryCMD.Flags().String("size", "", "size filter")
}

func summarize(cmd *cobra.Command, args []string) {
	filename, _ := cmd.Flags().GetString("file")
	log.Println(filename)
	size, _ := cmd.Flags().GetString("size")

	logs, err := parse(filename)
	if err != nil {
		log.Fatalf("error parsing log file: %v", err)
	}

	if size != "" {
		s, _ := strconv.Atoi(size)
		if err != nil {
			log.Fatalln("size is not an integer")
		}
		for operation, ls := range logs {
			larger, smaller := 0, 0
			for _, l := range ls {
				if l.Size < s {
					smaller++
				}
				if l.Size > s {
					larger++
				}
			}
			log.Printf("%ss larger than %d: %d\n", operation, s, larger)
			log.Printf("%ss smaller than %d: %d\n", operation, s, smaller)
		}
		return
	}

	for operation, ls := range logs {
		min, max, avg, total := summarizeOperation(ls)
		log.Printf("%s occurences: %d min: %d max: %d avg: %d total: %d\n", operation, len(logs[operation]), min, max, avg, total)
	}
}

func summarizeOperation(logs []types.Log) (min int, max int, avg int, total int) {
	for _, l := range logs {
		total += l.Size
		if min == 0 {
			min = l.Size
		}
		if min > l.Size {
			min = l.Size
		}
		if max < l.Size {
			max = l.Size
		}
	}

	avg = total / len(logs)
	return
}
