package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MysqlInstanceDelete(instanceName string) {
	client, err := Client()
	if err != nil {
		return
	}
	err = client.KubedbV1alpha2().MySQLs("demo").Delete(context.Background(), instanceName, v1.DeleteOptions{})
	if err != nil {
		fmt.Printf("mysql instance delete %s error %s \n", instanceName, err.Error())
		return
	}
	fmt.Printf("mysql instance delete success %s", instanceName)
}
