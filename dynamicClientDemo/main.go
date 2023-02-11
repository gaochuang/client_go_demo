package main

import (
	"context"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//配置需要的GVR
	gvr := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

	//发送请求，得到请求结果,unStructList是非结构化数据
	//Resource 返回一个基于GVR的客户端，动态资源客户端
	unStructData, err := dynamicClient.Resource(gvr).Namespace("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//unStructList转换为结构化数据
	podList := &v12.PodList{}
	runtime.DefaultUnstructuredConverter.FromUnstructured(
		unStructData.UnstructuredContent(),
		podList,
	)

	for _, pod := range podList.Items {
		fmt.Println(pod.Name)
	}

}
