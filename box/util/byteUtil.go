package util

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

func infoToBytes(n []interface{}) (b []byte) {
	for _, v := range n {
		if v1, ok := v.(string); ok {
			b = append(b, str2bytes(v1)...)
		} else {
			b = append(b, DataToBytes(v)...)
		}

	}
	return b
}

//整形转换成字节
func DataToBytes(n interface{}) []byte {
	var bytesBuffer bytes.Buffer
	binary.Write(&bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type Protocol struct {
	Magic      uint32
	Src        byte
	Dst        byte
	Command    uint16
	StatusCode uint16
	DataType   uint16
	Reserve    uint16
	Len        uint32
	Payload    []interface{}
}

func Encoding(p Protocol) (b []byte) {
	b = append(b, DataToBytes(p.Magic)...)
	b = append(b, p.Src)
	b = append(b, p.Dst)
	b = append(b, DataToBytes(p.Command)...)
	b = append(b, DataToBytes(p.StatusCode)...)
	b = append(b, DataToBytes(p.DataType)...)
	b = append(b, DataToBytes(p.Reserve)...)
	arr := infoToBytes(p.Payload)
	b = append(b, DataToBytes(int32(len(arr)))...)
	b = append(b, arr...)
	return b
}
