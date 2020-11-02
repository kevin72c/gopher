// client client for golang
// https://golangr.com
package client

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const ADDR = "192.168.2.14:30000"

var conn net.Conn

func init() {

	var err error
	fmt.Println("client init")
	// connect to server
	conn, err = net.Dial("tcp", ADDR)

	if err != nil {
		fmt.Println("连接服务端失败:", err)
		return
	}

	go read()

	go heartbeat()
}

func read() {
	for {

		header := make([]byte, 18)
		_, err := io.ReadFull(conn, header)
		if err != nil {
			conn.Close()
			log.Println(err)
			break
		}
		len := binary.BigEndian.Uint32(header[14:18])

		body := make([]byte, len)
		_, err = io.ReadFull(conn, body)
		if err != nil {
			conn.Close()
			log.Println(err)
			break
		}

		fmt.Println("response: ", header, body)
	}
}

func Send(msg string) {
	// 将 16进制的字符串 转换 byte
	//hexData, _ := hex.DecodeString("7E7E5A5A020000010000000000000000003A00000001aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa10dddddddddddddddddddddddddddddddd00000100000000000000000000000000")
	hexData, _ := hex.DecodeString(msg)
	conn.Write(hexData)
}

func heartbeat() {
	for {
		hexData, _ := hex.DecodeString("7E7E5A5A02000002000000000000000000020000")
		conn.Write(hexData)
		time.Sleep(10 * time.Second)
	}
}
