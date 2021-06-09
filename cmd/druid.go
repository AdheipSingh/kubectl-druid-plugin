package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/AdheipSingh/druid-kubectl-plugin/utils"
	do "github.com/druid-io/druid-operator/apis/druid/v1alpha1"
	"github.com/spf13/cobra"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

type druidListCmd struct {
	out io.Writer
}

func DruidClusterList(streams genericclioptions.IOStreams) *cobra.Command {
	druidCmdList := &druidListCmd{
		out: streams.Out,
	}

	cmd := &cobra.Command{
		Use:          "list",
		Short:        "Lists Druid Clusters in all namespaces",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("this command does not accept arguments")
			}
			getList()
			return druidCmdList.run()
		},
	}

	return cmd
}

func (sv *druidListCmd) run() error {
	_, err := fmt.Fprintf(sv.out, "Hello from Kubernetes server with version %s!\n", getList())
	if err != nil {
		return err
	}

	return nil
}

func getList() string {
	a := do.DruidList{}
	fmt.Println(a.APIVersion)

	clientset := utils.GetClientSet()

	deploymentRes := schema.GroupVersionResource{Group: "druid.apache.org", Version: "v1alpha1", Resource: "druids"}

	druidList, err := clientset.Resource(deploymentRes).Namespace(v1.NamespaceAll).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, d := range druidList.Items {
		fmt.Printf("%s\n", d.GetName())
	}

	return ""

}

var rootCmd = &cobra.Command{
	Use:          "druid",
	Long:         "kubectl druid plugin",
	SilenceUsage: true,
}

func NewCmdDruidPlugin(streams genericclioptions.IOStreams) *cobra.Command {
	rootCmd.AddCommand(DruidClusterList(streams))
	return rootCmd
}
