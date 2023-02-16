package main

import (
	"clinet-go-demo/NewClientFrameWork/pkg/informer"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
)

func main() {
	stopCh := make(chan struct{})
	err := informer.Setup(stopCh)
	if err != nil {
		panic(err.Error())
	}
	//启动一个web服务
	//2. 实例化Gin
	r := gin.Default()
	//2.写路由
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    200000,
			"message": "ping",
		})
	})

	r.GET("/pod/list", func(context *gin.Context) {
		items, err := informer.Get().Core().V1().Pods().Lister().List(labels.Everything())
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":    400,
				"message": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code":    5000,
				"message": "successful",
				"data":    items,
			})
		}
	})

	r.GET("/:resource/:group/:version", func(context *gin.Context) {
		resource := context.Param("resource")
		group := context.Param("group")
		version := context.Param("version")
		//组合GVR
		gvr := schema.GroupVersionResource{
			Resource: resource,
			Group:    group,
			Version:  version,
		}
		//通过GVR获取到对应资源对象的informer
		i, err := informer.Get().ForResource(gvr)
		if err != nil {
			panic(err.Error())
		}

		items, err := i.Lister().List(labels.Everything())
		if err != nil {
			panic(err.Error())
		}

		context.JSON(http.StatusOK, gin.H{
			"code":    2100,
			"message": "successful",
			"data":    items,
		})
	})

	_ = r.Run("192.168.31.100:8888")
}
