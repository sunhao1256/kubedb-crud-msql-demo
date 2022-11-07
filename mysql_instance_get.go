package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
)

func MysqlInstanceGet(instanceName string) *v1alpha2.MySQL {
	client, err := Client()
	if err != nil {
		return nil
	}
	instance, err := client.KubedbV1alpha2().MySQLs("demo").Get(context.Background(), instanceName, v1.GetOptions{})
	if err != nil {
		fmt.Printf("mysql instance get error %s \n", err.Error())
		return nil
	}
	fmt.Printf("mysql instance item %s", instance.Name)

	return instance
}
