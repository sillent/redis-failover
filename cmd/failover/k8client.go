package main

import (
	"context"
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func inList(svclblnm string, subs map[string]string) bool {
	for k := range subs {
		if k == svclblnm {
			return true
		}
	}
	return false
}
func checkPodForIP(neededSentMaster RedisMaster, pod v1.Pod) bool {
	if pod.Status.PodIP == neededSentMaster.IP {
		return true
	}
	return false
}
func checkPodLabel(pod v1.Pod) bool {
	if inList(ServiceLabelName, pod.GetLabels()) {
		return true
	}
	return false
}
func redisCheckEndpoint(neededSentMaster RedisMaster, namespace string, redisMasterServiceName string, redisstatefulset string) {
	// configuration entity
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println("'kubeconfig' ", err)
		return
	}

	//clientset entity
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("'kubeclientset' ", err)
		return
	}
	// getting pods list entity
	services, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("'kubegetpods' ", err)
		return
	}

	for _, p := range services.Items {
		if strings.Contains(p.GetName(), redisstatefulset) {
			if checkPodForIP(neededSentMaster, p) {
				if checkPodLabel(p) {
					/// Pod with master IP and correctly labeling
					log.Println("Pod is master and contain our label")
				} else {
					// Making label on master POD
					markingPod(p)
				}
			} else {
				if checkPodLabel(p) {
					// Deleting label from non-master POD
					log.Println("Pod not a master and contain our label")
					unmarkingPod(p)
				}
			}
		}

	}

}
