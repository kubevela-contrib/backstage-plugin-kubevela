package main

import (
	"log"
	"net/http"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	http.HandleFunc("/", entitySyncer)
	var err error
	k8sClient, err = client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
	}
	port := ":8080"
	log.Println("Backstage plugin for KubeVela starting at", port)
	http.ListenAndServe(port, nil)
}
