package main

import (
	"context"
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "C:\\Users\\prani\\.kube\\config", "Kuberconfig location")   //kubeconfig path 

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("Error : ", err)
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	ctx := context.Background()
	pods , err := clientset.CoreV1().Pods("default").List(ctx,metav1.ListOptions{})
	if err!=nil{
		fmt.Println("Error : ", err)
	}
	fmt.Printf("\n%s\t %s","CreationTime","Pod Name")
	
	for _,pod := range pods.Items {
		fmt.Printf("\n%s\t %s",pod.CreationTimestamp,pod.Name)	
	}
}
