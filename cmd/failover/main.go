package failover

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InCluster()
	if err != nil {
		panic(err.Error())
	}

	// clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Services().List(context.TODO, metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf(pods)

	}

}
