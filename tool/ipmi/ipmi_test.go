package main

import (
	"flag"
	"fmt"
	"log"
	"testing"
)

var c = flag.String("c", "x", "环境参数")

func TestIPMIServer_Query(t *testing.T) {
	flag.Parse()
	fmt.Printf("command:%s\n", *c)

	log.Printf("We are about to restart IPMI with this awesome package !")

	server := IPMIServer{
		Address:  "192.168.180.45",
		User:     "admin",
		Password: "12345678",
	}

	out, err := server.Query("chassis", "power", *c)
	if err != nil {
		log.Fatalf("An error occured")
	}

	log.Printf("result %s", out.String())
}
