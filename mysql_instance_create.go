package main

import (
	"context"
	"fmt"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
)

func MysqlInstanceCreate(instanceName string) {

	client, err := Client()
	if err != nil {
		return
	}

	mysqlInstance := &v1alpha2.MySQL{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
		},
		Spec: v1alpha2.MySQLSpec{
			Version: "5.7.36",
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

	create, err := client.KubedbV1alpha2().MySQLs("demo").Create(context.Background(), mysqlInstance, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("create mysql error %s \n", err.Error())
		return
	}

	fmt.Printf("create mysql success name %s \n", create.Name)

}
