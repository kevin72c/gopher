package client

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gopher/box/util"
	"io"
	"net"
	"time"
)

const ADDR = "192.168.2.14:37777"

var conn net.Conn

func init() {
	var err error
	conn, err = net.Dial("tcp", ADDR)
	if err != nil {
		fmt.Printf("connect fail:%v\n", err)
	} else {
		fmt.Println(" ok")
		handShake()
	}

	go read()
	go heartbeat()
}

func read() {
	for {
		header := make([]byte, 18)
		_, err := io.ReadFull(conn, header)
		if err != nil {
			//conn.Close()
			fmt.Printf("流信息读取失败1:%v\n", err)
			continue
		}

		p := util.Protocol{
			Magic:      binary.BigEndian.Uint32(header[0:4]),
			Src:        header[4],
			Dst:        header[5],
			Command:    binary.BigEndian.Uint16(header[6:8]),
			StatusCode: binary.BigEndian.Uint16(header[8:10]),
			DataType:   binary.BigEndian.Uint16(header[10:12]),
			Reserve:    binary.BigEndian.Uint16(header[12:14]),
			Len:        binary.BigEndian.Uint32(header[14:18]),
		}

		body := make([]byte, p.Len)
		_, err = io.ReadFull(conn, body)
		if err != nil {
			//conn.Close()
			fmt.Printf("流信息读取失败2:%v\n", err)
			continue
		}

		if p.Command != 2 {
			fmt.Printf("response msg: %+v %s \n", p, body)
		}
		handle(p, body)
	}

}

func write(p util.Protocol) {
	fmt.Printf("request msg: %+v \n", p)
	conn.Write(util.Encoding(p))
}

func heartbeat() {
	for {
		time.Sleep(10 * time.Second)
		hexData, _ := hex.DecodeString("7E7E5A5A02000002000000000000000000020000")
		conn.Write(hexData)
	}
}

func handShake() {
	p := util.Protocol{
		Magic:      0x7e7e5a5a,
		Src:        2,
		Dst:        0,
		Command:    0x0001,
		StatusCode: 0x0000,
		DataType:   0,
		Reserve:    0,
	}

	p.Payload = make([]interface{}, 0)
	p.Payload = append(p.Payload, int32(1))           // handshake version
	p.Payload = append(p.Payload, util.GetFirstIp())  //ip
	p.Payload = append(p.Payload, util.GetFirstMac()) // mac
	serialId, _ := hex.DecodeString("ffeeaa00000000000010000000000000")
	p.Payload = append(p.Payload, serialId) // SerialID
	firmwareVersion, _ := hex.DecodeString("00000100000000000000000000000000")
	p.Payload = append(p.Payload, firmwareVersion) // firmware version
	write(p)
}

func RequireConn() {
	p := util.Protocol{
		Magic:      0x7e7e5a5a,
		Src:        2,
		Dst:        0,
		Command:    0x0010,
		StatusCode: 0x0000,
		DataType:   0,
		Reserve:    0,
	}
	write(p)
}

func DisConn() {
	p := util.Protocol{
		Magic:      0x7e7e5a5a,
		Src:        2,
		Dst:        0,
		Command:    0x0011,
		StatusCode: 0x0000,
		DataType:   0,
		Reserve:    0,
	}
	write(p)
}
