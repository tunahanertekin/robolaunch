package kubeops

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateNamespace(t *testing.T) {
	name := "t-test"
	err := CreateNamespace(name, metav1.CreateOptions{DryRun: []string{"All"}})
	if err != nil {
		panic(err)
	}
}

func TestCreateRole(t *testing.T) {
	name := "default" // testing for default namespace that should have for all clusters!
	err := CreateRole(name, metav1.CreateOptions{DryRun: []string{"All"}})
	if err != nil {
		panic(err)
	}
}
