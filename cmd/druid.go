package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

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

	var namespace string
	cmd := &cobra.Command{
		Use:          "list",
		Short:        "Lists druid CR's in all namespaces or a specific namespce",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("this command does not accept arguments")
			}
			return druidCmdList.run(namespace)
		},
	}

	f := cmd.Flags()
	f.StringVar(&namespace, "namespace", "", "list druid CR in a specific namespace")

	return cmd
}

func (sv *druidListCmd) run(namespace string) error {

	for _, l := range reader.ListDruids(namespace) {
		_, err := fmt.Fprintf(sv.out, "%s\n", l)
		if err != nil {
			return err
		}
	}

	return nil
}
