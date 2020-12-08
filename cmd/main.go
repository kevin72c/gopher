package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ipconfig")

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func ExecCommand(arg string) (error, string) {
	c := exec.Command("cmd", "/C", arg)
	w := bytes.NewBuffer(nil)
	c.Stderr = w
	message := "执行" + arg + "文件抽取数据成功"
	_, err1 := os.Stat(arg)
	var err error
	//判断文件是否存在
	if err1 != nil {
		err := c.Run()
		if err != nil {
			fmt.Printf("Run returns: %s\n", err)
		}
		//处理中文乱码
		garbledStr := ConvertByte2String(w.Bytes(), GB18030)
		message = err1.Error() + garbledStr
		//文件不存在并且执行报错
		return err, message
	} else {
		err = c.Run()
		if err != nil {
			//处理中文乱码
			garbledStr := ConvertByte2String(w.Bytes(), GB18030)
			//文件存在 但执行bat文件报错
			return err, garbledStr
		}
	}
	//文件存在并且执行bat文件成功
	return err, message
}
