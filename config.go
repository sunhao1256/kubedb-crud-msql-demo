package main

import (
	"k8s.io/client-go/tools/clientcmd"
	kubedb_cs "kubedb.dev/apimachinery/client/clientset/versioned"
	"os"
)

func Client() (kubedb_cs.Interface, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	client, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	config, err := kubedb_cs.NewForConfig(client)
	if err != nil {
		return nil, err
	}
	return config, nil

}
