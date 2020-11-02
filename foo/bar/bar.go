package main

import (
	"../../box/util"
	"fmt"
)

func main() {

	fmt.Println(util.GetFirstMac())

	x := util.GetFirstIp()
	fmt.Printf("ips: %q\n", x)
}
