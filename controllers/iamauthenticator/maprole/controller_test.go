package maprole

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
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
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

	node := &iamauthenticatorv1beta1.MapRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node",
		},
		Spec: iamauthenticatorv1beta1.MapRoleSpec{
			RoleARN:  "arn:aws:iam::xxxxxxxxx:role/node",
			Username: "system:node:{{EC2PrivateDNSName}}",
			Groups: []string{
				"system:bootstrappers",
				"system:nodes",
			},
		},
	}

	admin := &iamauthenticatorv1beta1.MapRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "admin",
		},
		Spec: iamauthenticatorv1beta1.MapRoleSpec{
			RoleARN:  "arn:aws:iam::xxxxxxxxx:role/admin",
			Username: "cluster-admin",
			Groups: []string{
				"system:masters",
			},
		},
	}

	reconciler := Reconciler{
		Client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(configmap, node, admin).Build(),
		Log:    ctrl.Log,
		Scheme: scheme,
		ConfigMap: types.NamespacedName{
			Namespace: configmap.ObjectMeta.Namespace,
			Name:      configmap.ObjectMeta.Name,
		},
	}

	_, err = reconciler.Reconcile(context.TODO(), reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name: node.ObjectMeta.Name,
		},
	})
	assert.Nil(t, err)

	err = reconciler.Client.Get(context.TODO(), reconciler.ConfigMap, configmap)
	assert.Nil(t, err)

	expected, err := os.ReadFile("./testdata/data.yaml")
	assert.Nil(t, err)

	assert.NotEmpty(t, configmap.Data)
	assert.Equal(t, string(expected), configmap.Data[FieldName])
}
