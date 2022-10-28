package main

import (
	"log"

	"github.com/abrucker235/log-analyzer/cmd"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}
