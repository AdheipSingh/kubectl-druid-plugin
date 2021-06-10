package cmd

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// GVK for Druid CR
var GVK = schema.GroupVersionResource{
	Group:    "druid.apache.org",
	Version:  "v1alpha1",
	Resource: "druids",
}

type readers interface {
	ListDruids(namespaces string) []string
}
type client struct {
	dynamic.Interface
}

// initalize Reader
var reader readers = client{newClient()}

func (c client) ListDruids(namespaces string) []string {

	var druidList *unstructured.UnstructuredList
	var err error
	//	if namespaces == "all" {
	//	druidList, err = c.Resource(GVK).Namespace(v1.NamespaceAll).List(context.TODO(), v1.ListOptions{})
	//	if err != nil {
	//		panic(err)
	//	}
	//} else {
	druidList, err = c.Resource(GVK).Namespace(namespaces).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	//	}

	var names []string
	for _, d := range druidList.Items {
		names = append(names, d.GetName())
	}

	return names

}
