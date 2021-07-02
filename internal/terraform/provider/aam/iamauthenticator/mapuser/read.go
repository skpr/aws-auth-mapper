package mapuser

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Read the MapRole resource.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*clientset.Clientset)

	mapRole, err := conn.IamauthenticatorV1beta1().MapUsers().Get(d.Id(), metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(FieldName, mapRole.ObjectMeta.Name)
	d.Set(FieldUserARN, mapRole.Spec.UserARN)
	d.Set(FieldUsername, mapRole.Spec.Username)
	d.Set(FieldGroups, mapRole.Spec.Groups)

	return nil
}
