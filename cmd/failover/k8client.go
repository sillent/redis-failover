package main

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func redisCheckEndpoint(senmaster RedisMaster, namespace string, redisMasterService string) {
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
	for _, s := range services.Items {
		log.Println(s.ObjectMeta.GetName())

	}

}
