package main

import (
	"context"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

/*
func FindActiveMaster() string {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)

}
*/

// ChangeService change K8s service pointing to new Redis Master pod
// New master pod labled special lable
func ChangeService(service string, hostname string, port int) {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	// clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	svcs, err := clientset.CoreV1().Services(service).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for {
		var svcList []byte
		for _, it := range svcs.Items {
			// fmt.Println(it.metav1.ListMeta.name)
			svcList, err = it.Marshal()
			if err != nil {
				panic(err)
			}
			log.Println("svcs: ", svcList)

		}
		time.Sleep(5 * time.Second)
	}

}

func main() {
	ChangeService("", "pod", 3530)

}
