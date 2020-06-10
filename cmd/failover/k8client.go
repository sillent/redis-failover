package main

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func inList(search string, where []v1.Service) bool {
	for _, s := range where {
		curN := s.ObjectMeta.GetName()
		if curN == search {
			return true
		}
	}
	return false
}
func redisCheckEndpoint(neededSentMaster RedisMaster, namespace string, redisMasterServiceName string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println("'config' ", err)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("'clientset' ", err)
		return
	}
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("'getservice' ", err)
		return
	}
	// serviceItems := services.Items
	fmt.Println(services)

}
