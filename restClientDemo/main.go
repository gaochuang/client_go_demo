package main

import (
	"context"
	"fmt"
	v12 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	v1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

func getPodList(client *rest.RESTClient) {
	podList := v12.PodList{}
	err := client.
		Get().
		Namespace("default").
		Resource("pods").
		Do(context.TODO()).
		Into(&podList)
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range podList.Items {
		fmt.Println(pod.Name)
	}

}

func main() {
	//1.加载kube/config文件
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}

	///api/v1/namespaces/{namespace}/pods

	//2.配置API路径
	config.APIPath = "api"

	//3.配置分组版本GV
	config.GroupVersion = &v1.SchemeGroupVersion //GV Group: "", Version: "v1"

	//4.配置数据解码的工具
	config.NegotiatedSerializer = scheme.Codecs
	//5.实例化一个restClient
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	getPodList(restClient)
}
