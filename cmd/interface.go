package cmd

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var GVK = schema.GroupVersionResource{
	Group:    "druid.apache.org",
	Version:  "v1alpha1",
	Resource: "druids",
}

type Readers interface {
	ListDruids(namespaces string) []string
}
type Client struct {
	dynamic.Interface
}

var Reader Readers = Client{NewClient()}

func NewClient() *Client {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &Client{dynamicClient}
}

func (client Client) ListDruids(namespaces string) []string {

	var druidList *unstructured.UnstructuredList
	var err error
	if namespaces == "" {
		druidList, err = client.Resource(GVK).Namespace(v1.NamespaceAll).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			panic(err)
		}
	} else {
		druidList, err = client.Resource(GVK).Namespace(namespaces).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			panic(err)
		}
	}

	var names []string
	for _, d := range druidList.Items {
		names = append(names, d.GetName())
	}

	return names

}
