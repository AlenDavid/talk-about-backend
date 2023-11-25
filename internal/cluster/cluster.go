package cluster

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func configFromFile() (*rest.Config, error) {
	var kube *string
	if home := homedir.HomeDir(); home != "" {
		kube = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kube = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	return clientcmd.BuildConfigFromFlags("", *kube)
}

func Connect() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return configFromFile()
	}

	return config, err
}
