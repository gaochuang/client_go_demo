package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Clients struct {
	clientSet kubernetes.Interface
}

var config *rest.Config

func NewClient() (clients Clients) {

	var err error
	//加载配置文件
	config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}
	//实例化各种客户端
	clients.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return
}

func (c *Clients) ClientSet() kubernetes.Interface {
	return c.clientSet
}

func GetConfig() *rest.Config {
	return config
}
