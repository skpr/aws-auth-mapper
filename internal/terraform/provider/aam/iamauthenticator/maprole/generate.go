package maprole

import (
	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}

	groups := d.Get(FieldGroups).([]interface{})

	for _, group := range groups {
		mapRole.Spec.Groups = append(mapRole.Spec.Groups, group.(string))
	}

	return mapRole, nil
}
