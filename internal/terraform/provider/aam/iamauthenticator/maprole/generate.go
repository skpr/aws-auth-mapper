package maprole

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	iamauthenticatorv1beta1 "github.com/skpr/aws-auth-mapper/apis/iamauthenticator/v1beta1"
)

// Generate the MapRole resource.
func Generate(d *schema.ResourceData) (iamauthenticatorv1beta1.MapRole, error) {
	mapRole := iamauthenticatorv1beta1.MapRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.Get(FieldName).(string),
		},
		Spec: iamauthenticatorv1beta1.MapRoleSpec{
			RoleARN:  d.Get(FieldRoleARN).(string),
			Username: d.Get(FieldUsername).(string),
			Groups:   d.Get(FieldGroups).([]string),
		},
	}

	return mapRole, nil
}
