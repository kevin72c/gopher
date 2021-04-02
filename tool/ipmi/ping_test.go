package main

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	x := Ping("192.168.2.141")
	fmt.Println(x)
}
