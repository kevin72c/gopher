package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
	"strings"
)

func main() {
	var site string = `https://oapi.dingtalk.com/robot/send?access_token=32e4e1101ac76d4ce4bdb6b4e760c657e16996cb34397fc26509efab9513c32d`

	var text_content0 string = `{"msgtype": "text", 
    "text": {
        "content": "我就是我, 是不一样的烟火"
     }
}`

	// var text_content1 string =`
	// {
	//     "msgtype": "markdown",
	//     "markdown": {
	//         "title": "测试markdown",
	//         "text": " 整理知识，学习笔记 \n  发布日记，杂文，所见所想 \n  ![cmd-markdown-logo  ](https: //www.zybuluo.com/static/img/logo.png)\n"
	//     }
	// }
	// ` //没搞定换行问题

	//  var text_content2 string = `
	// {
	//     "msgtype": "link",
	//     "link": {
	//         "text": "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。
	// 而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
	//         "title": "时代的火车向前开",
	//         "picUrl": "",
	//         "messageUrl": "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
	//     }
	// }
	//  `
	client := &http.Client{}

	req, err := http.NewRequest("POST", site, strings.NewReader(text_content0))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
