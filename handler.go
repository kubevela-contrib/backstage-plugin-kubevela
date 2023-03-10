package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	BackstageComponent = "backstage-component"
	BackstageLocation  = "backstage-location"
)

func SyncEntity(w http.ResponseWriter, r *http.Request) {
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
		ann := app.Annotations
		if ann == nil {
			ann = map[string]string{}
		}

		if ann[AnnOwner] == "" {
			ann[AnnOwner] = "kubevela"
		}
		if ann[AnnDomain] == "" {
			ann[AnnDomain] = "kubevela"
		}
		if ann[AnnSystem] == "" {
			ann[AnnSystem] = app.Name
			res = append(res, ConvertBackstageSystem(ann, app))
		}

		for _, comp := range app.Spec.Components {
			res = append(res, ConvertBackstageComponent(ann, app, comp))
			l := ConvertBackstageLocation(app, comp)
			if l.Kind != "" {
				res = append(res, l)
			}
		}
	}
	j, err := json.Marshal(res)
	if err != nil {
		return
	}
	_, err = w.Write(j)
	return
}

func fillNecessaryAnnotation(ann map[string]string, app v1beta1.Application) map[string]string {
	if ann[managedByLocation] == "" {
		ann[managedByLocation] = "vela-app:" + app.Namespace + ":" + app.Name
	}
	if ann[managedByOriginLocation] == "" {
		ann[managedByOriginLocation] = ann[managedByLocation]
	}
	return ann
}

func getTagsFromAnn(ann map[string]string) []string {
	rawtags := ann[AnnTags]
	tags := strings.Split(rawtags, ",")
	if len(tags) == 1 && tags[0] == "" {
		tags = []string{}
	}
	for i, v := range tags {
		tags[i] = strings.TrimSpace(v)
	}
	return tags
}

// ConvertBackstageSystem will handle OAM Application as Backstage System
func ConvertBackstageSystem(ann map[string]string, app v1beta1.Application) Entity {
	tags := getTagsFromAnn(ann)
	tags = append(tags, "vela-app", app.Namespace)
	return Entity{
		ApiVersion: "backstage.io/v1alpha1",
		Kind:       "System",
		Metadata: &EntityMeta{
			Name:      app.Name,
			Namespace: app.Namespace,

			Tags:        tags,
			Description: ann[AnnDescription],
			Title:       ann[AnnTitle],

			Annotations: fillNecessaryAnnotation(ann, app),
			Labels:      app.Labels,

			//TODO: handle links for system
		},
		Spec: map[string]interface{}{
			"owner":  ann[AnnOwner],
			"domain": ann[AnnDomain],
		},
	}
}

func fillDefaultSpec(bt *VelaBackstageTrait, appAnn map[string]string, app v1beta1.Application, comp common.ApplicationComponent) {
	if bt.System == "" {
		bt.System = appAnn[AnnSystem]
	}
	if bt.Type == "" {
		bt.Type = comp.Type
	}
	if bt.LifeCycle == "" {
		bt.LifeCycle = "default"
	}
	if bt.Owner == "" {
		bt.Owner = appAnn[AnnOwner]
	}
	if len(bt.Annotations) == 0 {
		bt.Annotations = appAnn
	}
	if len(bt.Labels) == 0 {
		bt.Labels = app.Labels
	}
	if len(bt.Tags) == 0 {
		bt.Tags = getTagsFromAnn(appAnn)
	}
}

// ConvertBackstageComponent will handle OAM Application Component as Backstage Component
func ConvertBackstageComponent(appAnn map[string]string, app v1beta1.Application, comp common.ApplicationComponent) Entity {
	var bt VelaBackstageTrait
	for _, tr := range comp.Traits {
		if tr.Type == "backstage" || tr.Type == BackstageComponent {
			data, err := json.Marshal(tr.Properties)
			if err != nil {
				log.Println("marshal backstage trait failed", err)
				break
			}
			err = json.Unmarshal(data, &bt)
			if err != nil {
				log.Println("unmarshal backstage trait failed", err)
				break
			}
			break
		}
	}
	fillDefaultSpec(&bt, appAnn, app, comp)
	tags := append(bt.Tags, "vela-component", app.Namespace)
	relations := []EntityRelation{}
	for _, dep := range comp.DependsOn {
		relations = append(relations, EntityRelation{
			Type: "dependsOn",
			// Naming rule [<kind>:][<namespace>/]<name>
			TargetRef: "Component:" + app.Namespace + "/" + dep,
		})
	}

	return Entity{
		ApiVersion: "backstage.io/v1beta1",
		Kind:       "Component",
		Metadata: &EntityMeta{
			Name:      comp.Name,
			Namespace: app.Namespace,

			Tags:        tags,
			Description: bt.Description,
			Title:       bt.Title,

			Annotations: fillNecessaryAnnotation(bt.Annotations, app),
			Labels:      bt.Labels,
			Links:       bt.Links,
		},
		Spec: map[string]interface{}{
			"type":      bt.Type,
			"lifecycle": bt.LifeCycle,
			"owner":     bt.Owner,
			"system":    bt.System,
		},
	}
}

// ConvertBackstageLocation will handle OAM Trait as Backstage Location Entity
func ConvertBackstageLocation(app v1beta1.Application, comp common.ApplicationComponent) Entity {
	var bt VelaBackstageTrait
	var found bool
	for _, tr := range comp.Traits {
		if tr.Type == BackstageLocation {
			data, err := json.Marshal(tr.Properties)
			if err != nil {
				log.Println("marshal backstage trait failed", err)
				break
			}
			err = json.Unmarshal(data, &bt)
			if err != nil {
				log.Println("unmarshal backstage trait failed", err)
				break
			}
			found = true
			break
		}
	}
	if !found {
		return Entity{}
	}
	if bt.Annotations == nil {
		bt.Annotations = map[string]string{}
	}
	return Entity{
		ApiVersion: "backstage.io/v1alpha1",
		Kind:       "Location",
		Metadata: &EntityMeta{
			Name:        comp.Name,
			Namespace:   app.Namespace,
			Annotations: fillNecessaryAnnotation(bt.Annotations, app),
		},
		Spec: map[string]interface{}{
			"type":    bt.Type,
			"targets": bt.Targets,
		},
	}
}
