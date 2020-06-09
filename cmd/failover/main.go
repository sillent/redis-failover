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
func ChangeService(ns string, hostname string, port int) {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	// clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	svcs, err := clientset.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for {
		// var svcList []byte
		log.Println("Count of SVCs: ", len(svcs.Items))
		for _, it := range svcs.Items {
			// fmt.Println(it.metav1.ListMeta.name)
			log.Println("Services: ", it.Spec.ClusterIP)
		}
		time.Sleep(5 * time.Second)

	}

}

func main() {
	// ChangeService("test-redis", "pod", 3530)
	GetRedisMaster("rfs-redisfailover:26379", "mymaster")

}
