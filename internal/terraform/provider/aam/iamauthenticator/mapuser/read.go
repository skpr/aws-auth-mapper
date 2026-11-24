package mapuser

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Read the MapUser resource.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(*clientset.Clientset)

	mapUser, err := conn.IamauthenticatorV1beta1().MapUsers().Get(ctx, d.Id(), metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to get"))
	}

	d.Set(FieldName, mapUser.ObjectMeta.Name)
	d.Set(FieldUserARN, mapUser.Spec.UserARN)
	d.Set(FieldUsername, mapUser.Spec.Username)
	d.Set(FieldGroups, mapUser.Spec.Groups)

	return nil
}
