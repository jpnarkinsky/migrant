package adapter

import (
	"fmt"
	"os"
	"path/filepath"

	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sAdapter struct {
	config  *restclient.Config
	options map[string]string
}

// Create returns a new Kubernetes Adapter
func k8sCreate(options map[string]string) (Adapter, error) {
	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return k8sAdapter{
		config,
		options,
	}, nil
}

func (a k8sAdapter) Apply(_ string) error { return fmt.Errorf("not implemented") }
func (a k8sAdapter) Mark() error          { return fmt.Errorf("not implemented") }
func (a k8sAdapter) Save(_ string) error  { return fmt.Errorf("not implemented") }
func (a k8sAdapter) Status() error        { return fmt.Errorf("not implemented") }
