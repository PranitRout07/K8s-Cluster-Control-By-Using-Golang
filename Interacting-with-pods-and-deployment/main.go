package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	kubeconfig := flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Kubeconfig location")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Error:", err)

		config,err = rest.InClusterConfig()
		if err!=nil{
			fmt.Println("error",err)
		} 

	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error:", err)
		
	}

	ctx := context.Background()
	Pods, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	fmt.Println("Pods are:-")
	for _,i := range Pods.Items{
		fmt.Println(i.Name)
	}

	//get deployments 

	Deployments,err := clientSet.AppsV1().Deployments("default").List(ctx,metav1.ListOptions{})
	fmt.Println("Deployments are:-")
	if err!=nil{
		fmt.Println(err)
	}
	for _,i := range Deployments.Items {
		fmt.Println(i.Name)
	}
}
