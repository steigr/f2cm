package configmap

import (
	"context"
	"fmt"
	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

var (
	client *k8s.Client
	ctx    = context.Background()
)

func loadClient(kubeconfigPath string) (*k8s.Client, error) {
	data, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("read kubeconfig: %v", err)
	}

	// Unmarshal YAML into a Kubernetes config object.
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
	}
	return k8s.NewClient(&config)
}

func makeClient() (client *k8s.Client) {
	var err error

	if len(os.Getenv("KUBECONFIG")) > 0 {
		client, err = loadClient(os.Getenv("KUBECONFIG"))
	} else {
		client, err = k8s.NewInClusterClient()
	}

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func init() {
	client = makeClient()
}
