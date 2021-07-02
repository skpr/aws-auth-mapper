package mapuser

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Delete the MapRole resource.
func Delete(d *schema.ResourceData, m interface{}) error {
	conn := m.(*clientset.Clientset)

	return conn.IamauthenticatorV1beta1().MapUsers().Delete(d.Id(), &metav1.DeleteOptions{})
}
