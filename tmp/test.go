package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("D:\\bak\\install\\bin\\64bit\\obs64.exe", "")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("x error", err)
	}

	fmt.Println(out)
}
