package maprole

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Delete the MapRole resource.
func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(*clientset.Clientset)

	return diag.FromErr(conn.IamauthenticatorV1beta1().MapRoles().Delete(ctx, d.Id(), metav1.DeleteOptions{}))
}
