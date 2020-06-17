package main

import (
	"context"
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientv1 "k8s.io/client-go/kubernetes/typed/core/v1"
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

// Deleting label from pods which are not a master
func unmarkingPod(pod v1.Pod, podL clientv1.PodInterface) {
	labels := pod.GetLabels()
	log.Println("Labels in pod ", pod.GetName(), ": ", labels)
	delete(labels, ServiceLabelName)
	pod.SetLabels(labels)
	podL.Update(context.TODO(), &pod, metav1.UpdateOptions{})
	log.Println("Labels after deteling :", pod.GetLabels())
}

// Add a main rfailover label to pod which are master right now
func markingPod(pod v1.Pod, podL clientv1.PodInterface) {
	labels := pod.GetLabels()
	log.Println("Labels in pod ", pod.GetName(), ": ", labels)
	labels[ServiceLabelName] = "master"
	pod.SetLabels(labels)
	podL.Update(context.TODO(), &pod, metav1.UpdateOptions{})
	log.Println("Labels after updating :", pod.GetLabels())
}

// func markingPod(pod v1.Pod) v1.Pod
func redisCheckEndpoint(neededSentMaster RedisMaster, namespace string, redisMasterServiceName string, redisstatefulset string) {
	// configuration entity
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Panic("'kubeconfig' ", err)
	}

	//clientset entity
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic("'kubeclientset' ", err)
	}
	// getting pods list entity
	pods := clientset.CoreV1().Pods(namespace)
	podLists, err := pods.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Panic("'kubegetpods' ", err)
	}

	for _, p := range podLists.Items {
		if strings.Contains(p.GetName(), redisstatefulset) {
			if checkPodForIP(neededSentMaster, p) {
				if checkPodLabel(p) {
					/// Pod with master IP and correctly labeling
					log.Printf("Pod %s is master and contain our label.", p.GetName())
				} else {
					// Making label on master POD
					log.Printf("Pod %s is master and not contain master label. Marking...", p.GetName())
					markingPod(p, pods)
				}
			} else {
				if checkPodLabel(p) {
					// Deleting label from non-master POD
					log.Printf("Pod %s not a master and contain our label. Unmarking...", p.GetName())
					unmarkingPod(p, pods)
				}
			}
		}
	}
}
