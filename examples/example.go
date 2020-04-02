package main

import (
	"time"

	"github.com/zhangchong5566/manba/pkg/client"
)

func main() {
	//createCluster()
	//createServer()
	createAPI()
}

// 如果你的api server使用了"--discovery"参数启动
func getClientWithDiscovery() (client.Client, error) {
	return client.NewClientWithEtcdDiscovery("/services",
		time.Second*10,
		"127.0.0.1:2379")
}

// 如果你的api server没有使用"--discovery"参数启动
func getClient() (client.Client, error) {
	return client.NewClient(time.Second*10,
		"127.0.0.1:9092")
}
