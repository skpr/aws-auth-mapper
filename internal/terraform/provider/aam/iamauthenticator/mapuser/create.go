package mapuser

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Create the MapUser resource.
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(*clientset.Clientset)

	mapUser, err := Generate(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to generate"))
	}

	out, err := conn.IamauthenticatorV1beta1().MapUsers().Create(ctx, &mapUser, metav1.CreateOptions{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create"))
	}

	d.SetId(out.ObjectMeta.Name)

	return nil
}
