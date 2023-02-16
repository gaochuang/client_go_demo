package informer

import (
	clientIn "clinet-go-demo/NewClientFrameWork/pkg/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"time"
)

var informerFactory informers.SharedInformerFactory

func NewSharedInformerFactory(stopCh chan struct{}) (err error) {

	var (
		clients clientIn.Clients
	)

	//1.加载客户端
	clients = clientIn.NewClient()

	//2.实例化 sharedInformerFactory
	informerFactory = informers.NewSharedInformerFactory(clients.ClientSet(), time.Second*60)

	gvrs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "namespaces"},
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "apps", Version: "v1", Resource: "statefulsets"},
	}

	//3.启动informer
	for _, gvr := range gvrs {
		_, err := informerFactory.ForResource(gvr)
		if err != nil {
			panic(err.Error())
		}
	}

	//4.启动所有创建的informer
	informerFactory.Start(stopCh)
	//5.等待所有informer全量数据同步完成
	informerFactory.WaitForCacheSync(stopCh)

	return err
}

func Get() informers.SharedInformerFactory {

	return informerFactory
}

func Setup(stopCh chan struct{}) (err error) {
	err = NewSharedInformerFactory(stopCh)
	if err != nil {
		return err
	}
	return nil
}
