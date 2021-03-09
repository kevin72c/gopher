package main

import (
	_ "gopher/box/client"
	_ "gopher/box/http"
	"time"
)

func main() {

	for {
		time.Sleep(1000 * time.Second)
	}

}
