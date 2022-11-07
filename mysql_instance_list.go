package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MysqlInstanceList() {
	client, err := Client()
	if err != nil {
		return
	}
	list, err := client.KubedbV1alpha2().MySQLs("demo").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("mysql instance list error %s \n", err.Error())
		return
	}
	for _, v := range list.Items {
		fmt.Printf("mysql instance item %s", v.Name)
	}

}
