package util

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

func GetIpStr(ip []byte) (ipStr string) {
	var b strings.Builder
	ipv4 := (ip[10] == 0 || ip[10] == 0xff) &&
		(ip[11] == 0 || ip[11] == 0xff)
	if ipv4 {
		for _, e := range ip[12:16] {
			fmt.Fprint(&b, int(e), ".")
		}
		ipStr = strings.TrimSuffix(b.String(), ".")
	} else {
		for _, e := range ip {
			fmt.Fprint(&b, int(e), ":")
		}
		ipStr = strings.TrimSuffix(b.String(), ":")
	}
	return ipStr
}
func GetFirstMac() (b []byte) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return b
	}
	b, _ = hex.DecodeString(strings.ReplaceAll(netInterfaces[0].HardwareAddr.String(), ":", ""))
	return b
}

func GetFirstIp() (ip []byte) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ip
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP
			}
		}
	}
	return ip
}

func GetMacAddress() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func GetIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
