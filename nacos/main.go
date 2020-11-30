package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	var endpoint = "http://localhost:8848/"
	var namespaceId = "xxx"

	// 推荐使用 RAM 用户的 accessKey、secretKey
	//var accessKey = "${accessKey}"
	//var secretKey = "${secretKey}"

	clientConfig := constant.ClientConfig{
		//
		Endpoint:    endpoint,
		NamespaceId: namespaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs: 5 * 1000,
		Username:  "nacos",
		Password:  "xxx123",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "server-room-ctrl-system"
	var group = "xxx"

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	fmt.Println("Get config：" + content)

}
