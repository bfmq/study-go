package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	ctx = context.Background()
)

func main() {
	// config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	// if err != nil {
	// 	panic(err)
	// }
	// config.GroupVersion = &v1.SchemeGroupVersion
	// config.NegotiatedSerializer = scheme.Codecs
	// config.APIPath = "/api"

	// restClient, err := rest.RESTClientFor(config)
	// if err != nil {
	// 	panic(err)
	// }

	// pod := v1.Pod{}
	// restClient.Get().Namespace("devops").Resource("pods").Name("devops-vue-bbdbc8bb7-5gttv").Do(context.TODO()).Into(&pod)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pod.Name)

	config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pod, err := clientset.CoreV1().Pods("devops").Get(ctx, "devops-vue-bbdbc8bb7-5gttv", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod.Name)
}
