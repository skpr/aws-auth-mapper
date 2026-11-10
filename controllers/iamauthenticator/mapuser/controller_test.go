package mapuser

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake" //nolint:staticcheck // Migrating to envtest is out of scope.
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	iamauthenticatorv1beta1 "github.com/skpr/aws-auth-mapper/apis/iamauthenticator/v1beta1"
)

func TestReconcile(t *testing.T) {
	scheme := runtime.NewScheme()

	err := iamauthenticatorv1beta1.AddToScheme(scheme)
	assert.Nil(t, err)

	err = corev1.AddToScheme(scheme)
	assert.Nil(t, err)

	configmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "aws-auth",
			Namespace: "kube-system",
		},
	}

	admin1 := &iamauthenticatorv1beta1.MapUser{
		ObjectMeta: metav1.ObjectMeta{
			Name: "admin1",
		},
		Spec: iamauthenticatorv1beta1.MapUserSpec{
			UserARN:  "arn:aws:iam::xxxxxxxxx:user/admin1",
			Username: "cluster-admin",
			Groups: []string{
				"system:masters",
			},
		},
	}

	admin2 := &iamauthenticatorv1beta1.MapUser{
		ObjectMeta: metav1.ObjectMeta{
			Name: "admin2",
		},
		Spec: iamauthenticatorv1beta1.MapUserSpec{
			UserARN:  "arn:aws:iam::xxxxxxxxx:user/admin2",
			Username: "cluster-admin",
			Groups: []string{
				"system:masters",
			},
		},
	}

	reconciler := Reconciler{
		Client: fake.NewFakeClientWithScheme(scheme, configmap, admin1, admin2),
		Log:    ctrl.Log,
		Scheme: scheme,
		ConfigMap: types.NamespacedName{
			Namespace: configmap.Namespace,
			Name:      configmap.Name,
		},
	}

	_, err = reconciler.Reconcile(reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: admin1.Name,
		},
	})
	assert.Nil(t, err)

	err = reconciler.Get(context.TODO(), reconciler.ConfigMap, configmap)
	assert.Nil(t, err)

	expected, err := os.ReadFile("./testdata/data.yaml")
	assert.Nil(t, err)

	assert.NotEmpty(t, configmap.BinaryData)
	assert.Equal(t, string(expected), string(configmap.BinaryData[FieldName]))
}
