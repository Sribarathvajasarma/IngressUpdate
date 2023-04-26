package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/go-redis/redis"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	NameSpace := flag.String("ns", "", "Namespace")
	IngressName := flag.String("ingress", "", "Ingress")
	Host := flag.String("host", "", "Host")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\WSO2\\.kube\\config")
	if err != nil {
		panic(err)
	}

	if *NameSpace != "" && *IngressName != "" && *Host != "" {
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		ingressClient := clientset.NetworkingV1().Ingresses(*NameSpace)

		ingress, err := ingressClient.Get(context.TODO(), *IngressName, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}

		ingress.Spec.Rules[0].Host = *Host

		_, err = ingressClient.Update(context.TODO(), ingress, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}

		fmt.Println("Ingress host updated to " + *Host)
		client := redis.NewClient(&redis.Options{
			Addr:     "redis-16912.c12.us-east-1-4.ec2.cloud.redislabs.com:16912",
			Password: "FKSn3h1r4loQu9zY4qFyXSjjyY462Kw6",
			DB:       0,
		})

		err2 := client.Set(*Host, "34.100.199.109", 0).Err()
		// if there has been an error setting the value
		// handle the error
		if err2 != nil {
			fmt.Println(err2)
		}

	} else {
		fmt.Println("Please give inputs NameSpace, Ingress, Host")
	}

}
