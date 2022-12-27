package main

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	core_oam_dev "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev"
	kubevelaapistandard "github.com/oam-dev/kubevela-core-api/apis/standard.oam.dev/v1alpha1"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
)

var scheme = runtime.NewScheme()

var k8sClient client.Client

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = apiregistrationv1.AddToScheme(scheme)
	_ = crdv1.AddToScheme(scheme)
	_ = core_oam_dev.AddToScheme(scheme)
	_ = kubevelaapistandard.AddToScheme(scheme)
}
