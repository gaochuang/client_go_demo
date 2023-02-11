package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}
	disCoverClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//发送请求获取GVR数据
	resources, lists, err := disCoverClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}

	for _, resource := range resources {
		fmt.Println(resource.Kind)
	}
	for _, list := range lists {
		kind := schema.ParseGroupKind(list.GroupVersion)
		fmt.Println(kind)
	}
}
