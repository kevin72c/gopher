package main

import (
	_ "./client"
	_ "./http"
	"time"
)

func main() {

	for {
		time.Sleep(1000 * time.Second)
	}

}
