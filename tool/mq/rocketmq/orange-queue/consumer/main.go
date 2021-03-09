package main

import (
	"fmt"

	"git.oschina.net/cloudzone/cloudmq-go-client/cloudmq"
)

func main() {
	conf := &cloudmq.Config{
		//Nameserver:   "10.122.1.201:9876",
		Nameserver:   "192.168.2.14:9876",
		InstanceName: "DEFAULT",
	}
	consumer, err := cloudmq.NewDefaultConsumer("group1", conf)
	if err != nil {
		panic(err)
	}
	consumer.Subscribe("t1", "*")

	var count int
	consumer.RegisterMessageListener(func(msgs []*cloudmq.MessageExt) (int, error) {
		for _, msg := range msgs {
			count++
			fmt.Printf("count=%d|msgId=%s|topic=%s|storeTimestamp=%d|bornTimestamp=%d|storeHost=%s|bornHost=%s"+
				"|msgTag=%s|msgKey=%s|sysFlag=%d|storeSize=%d|queueId=%d|queueOffset=%d|body=%s\n",
				count, msg.MsgId, msg.Topic, msg.StoreTimestamp, msg.BornTimestamp, msg.StoreHost, msg.BornHost,
				msg.Tag(), msg.Key(), msg.SysFlag, msg.StoreSize, msg.QueueId, msg.QueueOffset, msg.Body)
		}
		return cloudmq.Action.CommitMessage, nil
	})
	consumer.Start()

	select {}
}
