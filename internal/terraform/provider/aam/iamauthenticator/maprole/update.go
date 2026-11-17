package maprole

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Update the MapRole resource.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(*clientset.Clientset)

	p, err := Generate(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to generate"))
	}

	mapRole, err := conn.IamauthenticatorV1beta1().MapRoles().Get(ctx, p.ObjectMeta.Name, metav1.GetOptions{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed get existing MapRole"))
	}

	mapRole.Spec = p.Spec

	_, err = conn.IamauthenticatorV1beta1().MapRoles().Update(ctx, mapRole, metav1.UpdateOptions{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to update"))
	}

	return nil
}
