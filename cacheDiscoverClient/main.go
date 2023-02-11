package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}

	//实例化客户端，本客户端负责将GVR数据缓存到本地文件夹中
	cacheDiscoverClient, err := disk.NewCachedDiscoveryClientForConfig(config, "./cache/discover", "./cache/http", time.Minute*60)
	if err != nil {
		panic(err.Error())
	}

	_, apiResources, err := cacheDiscoverClient.ServerGroupsAndResources()
	//1.先从缓存文件中找GVR数据，有则返回，没有则需要调用API Server
	//2.调用API Server获取GVR数据
	//3.将获取到GVR数据缓存到本地，返回给客户端

	for _, list := range apiResources {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err.Error())
		}

		for _, resource := range list.APIResources {
			fmt.Printf("name: %v, group: %v, version: %v \n", resource.Name, gv.Group, gv.Version)
		}
	}

}
