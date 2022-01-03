package kubeops

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateNamespace(t *testing.T) {
	name := "t-test"
	ns, err := CreateNamespace(name, metav1.CreateOptions{DryRun: []string{"All"}})
	if err != nil {
		panic(err)
	}
	assert.Equal(t, name, ns.Name)

}

func TestCreateRole(t *testing.T) {
	name := "default" // testing for default namespace that should have for all clusters!
	role, roleBind, err := CreateRole(name, metav1.CreateOptions{DryRun: []string{"All"}})
	if err != nil {
		panic(err)
	}
	assert.Equal(t, name+"_role", role.Name)
	assert.Equal(t, name, role.Namespace)

	assert.Equal(t, name+"_role", roleBind.Name)
	assert.Equal(t, name, roleBind.Namespace)

}
