package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Getting the kubeconfig file path and namespace from the environment variables
	kubeConfig := os.Getenv("KUBECONFIG")
	namespace := os.Getenv("NAMESPACE")

	if kubeConfig == "" {
		log.Fatal("KUBECONFIG is not set")
	}

	if namespace == "" {
		log.Print("NAMESPACE is not set, using default namespace")
		namespace = "default"
	}

	// Building the Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Starting the main CRUD operations

	// Listing pods in the specified namespace
	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods in namespace %s: %v", namespace, err)
	}

	if len(pods.Items) == 0 {
		fmt.Printf("No pods found in namespace: %s\n", namespace)
	} else {
		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}
	}

	// Creating a pod in the specified namespace
	podDefinition := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "main-",
			Namespace:    namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "main",
					Image: "prasadb89/prasad89.github.io",
				},
			},
		},
	}

	createdPod, err := client.CoreV1().Pods(namespace).Create(ctx, podDefinition, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating pod in namespace %s: %v", namespace, err)
	}

	fmt.Printf("Pod created successfully: %s\n", createdPod.Name)

	// Updating the pod in the specified namespace

	// Deleting the pod in the specified namespace
}
