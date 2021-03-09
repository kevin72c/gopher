package main

import (
	"fmt"
	"os/exec"
	"strings"
)

//GetCPUID 获取cpuid
func GetCPUID() string {
	var cpuid string
	cmd := exec.Command("wmic", "cpu", "get", "processorid")
	b, e := cmd.CombinedOutput()

	if e == nil {
		cpuid = string(b)
		cpuid = cpuid[12 : len(cpuid)-2]
		cpuid = strings.ReplaceAll(cpuid, "\n", "")
	} else {
		fmt.Printf("%v", e)
	}

	return cpuid
}

//GetBaseBoardID 获取主板的id
func GetBaseBoardID() string {
	var cpuid string
	cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	b, e := cmd.CombinedOutput()

	if e == nil {
		cpuid = string(b)
		cpuid = cpuid[12 : len(cpuid)-2]
		cpuid = strings.ReplaceAll(cpuid, "\n", "")
	} else {
		fmt.Printf("%v", e)
	}

	return cpuid
}

func main() {
	x := GetBaseBoardID()
	fmt.Println(&x)
	x = GetCPUID()
	fmt.Println(x)
}
