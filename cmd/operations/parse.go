package operations

import (
	"log"

	"github.com/abrucker235/log-analyzer/types"
	"github.com/abrucker235/log-analyzer/util"
)

func parse(filename string) (map[string][]types.Log, error) {
	logs, err := util.Parse(filename)
	if err != nil {
		log.Fatalf("error parsing log file: %v", err)
	}

	operationLogs := make(map[string][]types.Log)
	for _, log := range logs {
		operationLogs[log.Operation] = append(operationLogs[log.Operation], log)
	}

	return operationLogs, nil
}
