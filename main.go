package main

import (
	"flag"
	"os"

	"github.com/AdheipSingh/druid-kubectl-plugin/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/klog"
)

var version = "undefined"

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("log_file", "logs.log")
	flag.Parse()

	if err := cmd.NewCmdDruidPlugin(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}).Execute(); err != nil {
		// make sure we flush before exiting
		klog.Flush()
		os.Exit(1)
	}
	// make sure we flush before exiting
	klog.Flush()
}
