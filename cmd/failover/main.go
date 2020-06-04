package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// ChangeService change K8s service pointing to new Redis Master pod
// New master pod labled special lable
func ChangeService(service string, hostname string, port int) (bool, error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		return false, nil
	}

	// clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, err
	}

	svcs, err := clientset.CoreV1().Services(service).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, err
	}
	var svcList []byte
	for _, it := range svcs.Items {
		// fmt.Println(it.metav1.ListMeta.name)
		svcList, err = it.Marshal()
		if err != nil {
			return false, err
		}
		fmt.Println("svclist: ", svcList)

	}

	return true, nil

}

func main() {
	f, err := ChangeService("", "pod", 3530)
	if err != nil {
		panic(err)

	}
	fmt.Println(f)
}
