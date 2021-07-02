package maprole

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Update the MapRole resource.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*clientset.Clientset)

	p, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	mapRole, err := conn.IamauthenticatorV1beta1().MapRoles().Get(p.ObjectMeta.Name, metav1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "failed get existing MapRole")
	}

	mapRole.Spec = p.Spec

	_, err = conn.IamauthenticatorV1beta1().MapRoles().Update(mapRole)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
