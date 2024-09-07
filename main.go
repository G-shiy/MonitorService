package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	now, url := config()
	get, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	elapsedTime := time.Since(now).Milliseconds()
	statusCodes := get.StatusCode

	fmt.Println("Time: ", elapsedTime, "ms", "Status Code: ", statusCodes)

	defer get.Body.Close()

}

func config() (time.Time, string) {
	now := time.Now()
	url := os.Args[1]

	return now, url
}
