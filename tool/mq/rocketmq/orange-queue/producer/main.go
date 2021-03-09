package main

import (
	"fmt"
	"gitee.com/zhucheer/orange/queue"
	"time"
)

func main() {
	// 注册生产者 填入broker节点,group名称,重试次数信息
	mqProducerClient := queue.RegisterRocketProducerMust([]string{"192.168.2.14:9876"}, "group", 1)

	for i := 0; i < 10; i++ {
		// 向队列发送一条消息 填入消息队列topic和消息体信息
		ret, _ := mqProducerClient.SendMsg("t1", "Hello mq~~")
		fmt.Println("========producer push one message====", ret.MsgId)

		time.Sleep(time.Millisecond)
	}

}
