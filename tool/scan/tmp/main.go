package tmp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

func main() {
	absPath, _ := filepath.Abs("tool/scan/config.txt")

	fmt.Println(absPath)
	Scanner(absPath, "", "", 1)
}

//首先从命令行中读取线程数和配置文件路径
//从配置文件中解析出ip和port
//配置文件格式为
// [ip]
// 127.0.0.1
// [port]
// 22
// 36000
// 56000
// 3306
//根据开启的线程数对指定ip和端口进行tcp连接
//如果端口开启，把ip:port按照格式返回
func Scanner(configFile string, functionid string, sendInfoFile string, limit int) {

	runtime.GOMAXPROCS(runtime.NumCPU())
	data, err := ioutil.ReadFile(configFile)
	portlist := make([]string, 0, 10)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ip_index := bytes.Index(data, []byte("[ip]"))

	port_index := bytes.Index(data, []byte("[port]"))

	if ip_index < 0 {

		fmt.Println("文件格式有误: missing [ip]")
		return
	}

	if port_index < 0 {

		portlist = append(portlist, "22")
		portlist = append(portlist, "36000")
		portlist = append(portlist, "56000")
		portlist = append(portlist, "3306")

	} else {
		regPort := regexp.MustCompile(`\d+`)
		var ports [][]byte

		if ip_index > port_index {
			ports = regPort.FindAll(data[:ip_index], -1)

		} else {

			ports = regPort.FindAll(data[port_index:], -1)
		}

		for _, v := range ports {

			portlist = append(portlist, string(v))

		}

	}

	regIp := regexp.MustCompile(`((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))`)

	ips := regIp.FindAll(data, -1)

	input := make(chan []byte, len(ips))

	result := make(chan string, len(ips))

	defer close(input)
	defer close(result)

	for _, v := range ips {
		input <- v
	}
	//控制多少并发
	for i := 0; i < limit; i++ {
		//这个时候可以启动扫描函数
		go ScanPort(portlist, input, result)
	}

	for i := 0; i < len(ips); i++ {

		//将扫描的结果输出
		ipResult, ok := <-result

		if !ok {
			break
		}
		fmt.Println(ipResult)
	}

}

func ScanPort(portList []string, intPut chan []byte, result chan string) {

	for {
		task, ok := <-intPut

		if !ok {

			return
		}
		ip := string(task)
		fmt.Println("scaning ", ip, "port", portList)
		portStr := ""
		for i := 0; i < len(portList); i++ {
			_, err := net.DialTimeout("tcp", ip+":"+portList[i], time.Second*3)

			if err != nil {
				continue
			}
			portStr += portList[i] + " "

		}

		if len(portStr) > 0 {
			//说明有打开的端口
			// Info.Println(ip+":"+portStr+"open")
			result <- ip + "----" + "1" + "----" + "open port:" + portStr

		} else {

			result <- ip + "----" + "0" + "----" + "ok"
		}

	}

}
