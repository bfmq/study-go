package main

import (
	clientset "github.com/operator-crd/pkg/generated/clientset/versioned"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	if err != nil {
		inClusterConfig, err := rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
		config = inClusterConfig
	}

	clientset, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	clientset.CrdV1
}
