package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 1; i++ {
			time.Sleep(100 * time.Millisecond)
			c1 <- i
			close(c1)
		}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	defer close(c1)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received1", msg1)
		case msg2 := <-c2:
			fmt.Println("received2", msg2)
		}
	}
}
