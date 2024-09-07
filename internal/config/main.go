package config

import (
	"fmt"
	"os"
	"time"
)

func Init() (time.Time, string) {
	now := time.Now()
	url := os.Args[1]

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}

	return now, url
}
