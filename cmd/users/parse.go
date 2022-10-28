package users

import (
	"log"

	"github.com/abrucker235/log-analyzer/types"
	"github.com/abrucker235/log-analyzer/util"
)

func parse(filename string) (map[string]map[string][]types.Log, error) {
	logs, err := util.Parse(filename)
	if err != nil {
		log.Fatalf("error parsing log file: %v", err)
	}

	userLogs := make(map[string]map[string][]types.Log)
	for _, log := range logs {
		operations := userLogs[log.Username]
		if operations == nil {
			operations = make(map[string][]types.Log)
			userLogs[log.Username] = operations
		}
		userLogs[log.Username][log.Operation] = append(operations[log.Operation], log)
	}

	return userLogs, nil
}
