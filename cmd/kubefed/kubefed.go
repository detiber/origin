package main

import (
	"fmt"
	"os"

	"k8s.io/kubernetes/federation/pkg/kubefed"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus" // for client metric registration
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/util/logs"
	_ "k8s.io/kubernetes/pkg/version/prometheus" // for version metric registration

	"github.com/openshift/origin/pkg/version"
)

var (
	// serverImageName is the name of the default image (without version)
	// used for the federation services (api and controller manager).  It
	// should be set during build via -ldflags.
	serverImageName string

	// defaultEtcImage is the default image (including version) used to run
	// etcd for the federation apiserver.  It should be set during build via
	// -ldflags.
	defaultEtcdImage string
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	defaultServerImage := fmt.Sprintf("%s:%s", serverImageName, version.Get())
	cmd := kubefed.NewKubeFedCommand(cmdutil.NewFactory(nil), os.Stdin, os.Stdout, os.Stderr, defaultServerImage, defaultEtcdImage)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
