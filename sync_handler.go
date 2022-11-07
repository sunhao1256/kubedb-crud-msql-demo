package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	kubedb_et "kubedb.dev/apimachinery/client/informers/externalversions"
)

func SyncHandler() {
	client, err := Client()
	if err != nil {
		return
	}
	factory := kubedb_et.NewSharedInformerFactory(client, 0)

	stop := make(chan struct{})

	mysqlInformer := factory.Kubedb().V1alpha2().MySQLs().Informer()
	mysqlInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mysql, ok := obj.(*v1alpha2.MySQL)
			if !ok {
				fmt.Println("mysql add error")
			}
			fmt.Printf("mysql add name %s \n", mysql.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {

			oldMysql, ok := oldObj.(*v1alpha2.MySQL)
			newMysql, ok := newObj.(*v1alpha2.MySQL)
			if !ok {
				fmt.Println("mysql update error")
			}
			fmt.Printf("mysql update old name %s new name  %s \n", oldMysql.Name, newMysql.Name)

		},

		DeleteFunc: func(obj interface{}) {
			mysql, ok := obj.(*v1alpha2.MySQL)
			if !ok {
				fmt.Println("mysql delete error")
			}
			fmt.Printf("mysql delete name %s \n", mysql.Name)
		},
	})

	go factory.Start(stop)
	fmt.Printf("start syncing")
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stop, mysqlInformer.HasSynced); !ok {
		fmt.Errorf("failed to wait for caches to sync")
		return
	} else {
		ls, err := factory.Kubedb().V1alpha2().MySQLs().Lister().MySQLs("demo").List(labels.Everything())
		if err != nil {
			panic(err)
		}
		for _, v := range ls {

			fmt.Printf("mysql instance %s\n", v.Name)
		}
	}

	<-stop

}
