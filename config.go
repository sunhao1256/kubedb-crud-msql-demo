package main

import (
	"k8s.io/client-go/rest"
	kubedb_cs "kubedb.dev/apimachinery/client/clientset/versioned"
	"net"
	"os"
)

func Client() (kubedb_cs.Interface, error) {
	//kubeconfig := os.Getenv("KUBECONFIG")

	masterUrl := os.Getenv("KUBE_MASTER_HOST")
	masterPort := os.Getenv("KUBE_MASTER_PORT")
	tokenPath := os.Getenv("KUBE_TOKEN_PATH")

	if len(masterUrl) == 0 {
		masterUrl = "127.0.0.1"
	}
	if len(masterPort) == 0 {
		masterPort = "56923"
	}
	if len(tokenPath) == 0 {
		panic("ENV KUBE_MASTER_URL or KUBE_TOKEN_PATH required ")
	}

	token, err := os.ReadFile(tokenPath)

	if err != nil {
		panic(err)
	}

	kubeconfig := &rest.Config{
		Host:            "https://" + net.JoinHostPort(masterUrl, masterPort),
		BearerToken:     string(token),
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}

	//client, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	//if err != nil {
	//	return nil, err
	//}
	config, err := kubedb_cs.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	return config, nil

}
