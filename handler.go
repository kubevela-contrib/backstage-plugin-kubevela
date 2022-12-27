package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func entitySyncer(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			log.Println(err)
		}
		return
	}()

	log.Println("API requested", r.RemoteAddr, r.Method, r.RequestURI)

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Only GET method supported.")
		return
	}
	err = r.ParseForm()
	if err != nil {
		return
	}
	var ns = r.Form.Get("ns")
	var appList v1beta1.ApplicationList
	err = k8sClient.List(context.TODO(), &appList, &client.ListOptions{Namespace: ns})
	if err != nil {
		return
	}
	var res []Entity
	for _, app := range appList.Items {

		managedBy := app.Annotations[managedByLocation]
		if managedBy == "" {
			managedBy = "k8s:vela:application:" + app.Namespace + ":" + app.Name
		}
		managedByOrigin := app.Annotations[managedByOriginLocation]
		if managedByOrigin == "" {
			managedByOrigin = managedBy
		}

		res = append(res, Entity{
			ApiVersion: "backstage.io/v1beta1",
			Kind:       "Component",
			Metadata: &EntityMeta{
				Name:      app.Name,
				Namespace: app.Namespace,
				Tags:      []string{"vela-app"},
				Annotations: map[string]string{
					managedByLocation:       managedBy,
					managedByOriginLocation: managedByOrigin,
				},
			},
			Spec: map[string]interface{}{
				"type": "vela-application",

				// TODO: these following items should read from apps.
				"lifecycle": "production",
				"owner":     "vela-users",
				"system":    "vela-apps",
			},
		})
	}
	j, err := json.Marshal(res)
	if err != nil {
		return
	}
	_, err = w.Write(j)
	return
}
