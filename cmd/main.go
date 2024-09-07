package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/G-shiy/MonitorService/internal/config"
)

func main() {

	now, url := config.Init()
	get, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	elapsedTime := time.Since(now).Milliseconds()
	statusCodes := get.StatusCode

	fmt.Println("Time: ", elapsedTime, "ms", "Status Code: ", statusCodes)

	defer get.Body.Close()
}
