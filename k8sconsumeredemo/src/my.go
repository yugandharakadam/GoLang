package main

import (
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := "C:\\Users\\tk9kxh.PLM\\.kube\\config"
	config, err := clientcmd.BuildConfigFromFlags("https://api.keystosuccess.in", kubeconfig)

	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})

	for _, node := range nodes.Items {
		fmt.Printf("Node name=/%s\n", node.GetName())
	}

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//fmt.Printf(pods.String())
	for _, pod := range pods.Items {
		fmt.Printf("Pod name=/%s\n", pod.GetName())
	}
}