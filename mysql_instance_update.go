package main

import (
	"context"
	"fmt"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
)

func MysqlInstanceUpdate(instanceName string, replica int32) {

	//query before

	get := MysqlInstanceGet(instanceName)

	if get == nil {
		panic("miss existed instance")
	}

	client, err := Client()
	if err != nil {
		return
	}

	updateInstance := &v1alpha2.MySQL{
		ObjectMeta: v1.ObjectMeta{
			Name:            get.Name,
			ResourceVersion: get.ResourceVersion,
		},
		Spec: v1alpha2.MySQLSpec{
			Version: "5.7.36",
			//update replica
			Replicas: pointer.Int32P(replica),
			Storage: &core.PersistentVolumeClaimSpec{
				StorageClassName: pointer.StringP("standard"),
				AccessModes:      []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
				Resources: core.ResourceRequirements{
					Requests: map[core.ResourceName]resource.Quantity{
						core.ResourceStorage: resource.MustParse("1Gi"),
					},
				},
			},
		},
	}

	instance, err := client.KubedbV1alpha2().MySQLs("demo").Update(context.Background(), updateInstance, v1.UpdateOptions{})
	if err != nil {
		fmt.Printf("mysql instance update error %s \n", err.Error())
		return
	}
	fmt.Printf("mysql instance update %s", instance.Name)
}
