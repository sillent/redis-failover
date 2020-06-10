package main

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

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
	serviceItems := services.Items
	for _, s := range serviceItems {
		// log.Println(s.ObjectMeta.GetName())
		curName := s.ObjectMeta.GetName()
		// Searching desired service for working with him
		if curName == redisMasterServiceName {

		}

	}

}
