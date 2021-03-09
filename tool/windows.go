package tool

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"github.com/astaxie/beego/logs"
	"net"
	"strings"
	"syscall"
	"unsafe"
)

type cpuInfo struct {
	Name          string
	NumberOfCores uint32
	ThreadCount   uint32
}

func GetCPUInfo() {

	var cpuinfo []cpuInfo

	err := wmi.Query("Select * from Win32_Processor", &cpuinfo)
	if err != nil {
		return
	}
	fmt.Printf("Cpu info =%v", cpuinfo)
}

type operatingSystem struct {
	Name    string
	Version string
}

func GetOSInfo() {
	var operatingSystem []operatingSystem
	err := wmi.Query("Select * from Win32_OperatingSystem", &operatingSystem)
	if err != nil {
		return
	}
	fmt.Printf("OS info =%v", operatingSystem)
}

var kernel = syscall.NewLazyDLL("Kernel32.dll")

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func GetMemoryInfo() {

	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return
	}
	fmt.Printf("total=:%v", memInfo.ullTotalPhys)
	fmt.Printf("free=:%v", memInfo.ullAvailPhys)
}

type Network struct {
	Name       string
	IP         string
	MACAddress string
}

type intfInfo struct {
	Name       string
	MacAddress string
	Ipv4       []string
}

func GetNetworkInfo() error {
	intf, err := net.Interfaces()
	if err != nil {
		logs.Error("get network info failed: %v", err)
		return err
	}
	var is = make([]intfInfo, len(intf))
	for i, v := range intf {
		ips, err := v.Addrs()
		if err != nil {
			logs.Error("get network addr failed: %v", err)
			return err
		}
		//此处过滤loopback（本地回环）和isatap（isatap隧道）
		if !strings.Contains(v.Name, "Loopback") && !strings.Contains(v.Name, "isatap") {
			var network Network
			is[i].Name = v.Name
			is[i].MacAddress = v.HardwareAddr.String()
			for _, ip := range ips {
				if strings.Contains(ip.String(), ".") {
					is[i].Ipv4 = append(is[i].Ipv4, ip.String())
				}
			}
			network.Name = is[i].Name
			network.MACAddress = is[i].MacAddress
			if len(is[i].Ipv4) > 0 {
				network.IP = is[i].Ipv4[0]
			}

			fmt.Printf("network:=%v", network)
		}

	}

	return nil
}

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func GetStorageInfo() {
	var storageInfo []storageInfo
	var localStorages []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageInfo)
	if err != nil {
		return
	}

	for _, storage := range storageInfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size,
			Free:       storage.FreeSpace,
		}
		localStorages = append(localStorages, info)
	}
	fmt.Printf("localStorages:=%v", localStorages)
}

type gpuInfo struct {
	Name string
}

func GetGPUInfo() {

	var gpuInfo []gpuInfo
	err := wmi.Query("Select * from Win32_VideoController", &gpuInfo)
	if err != nil {
		return
	}
	fmt.Printf("GPU:=%v", gpuInfo[0].Name)
}
