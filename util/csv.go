package util

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/abrucker235/log-analyzer/types"
)

func Parse(filename string) ([]types.Log, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, _ = reader.Read()

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var logs []types.Log
	for _, record := range records {
		timestamp, err := time.Parse(time.UnixDate, record[0])
		if err != nil {
			log.Fatalf("error parsing timestamp: %v", err)
		}

		size, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatalf("error parsing size: %v", err)
		}

		logs = append(logs, types.Log{Timestamp: timestamp, Username: record[1], Operation: record[2], Size: size})
	}

	return logs, nil
}
