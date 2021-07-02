package mapuser

import (
	"github.com/hashicorp/terraform/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	iamauthenticatorv1beta1 "github.com/skpr/aws-auth-mapper/apis/iamauthenticator/v1beta1"
)

// Generate the MapRole resource.
func Generate(d *schema.ResourceData) (iamauthenticatorv1beta1.MapUser, error) {
	mapUser := iamauthenticatorv1beta1.MapUser{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.Get(FieldName).(string),
		},
		Spec: iamauthenticatorv1beta1.MapUserSpec{
			UserARN:  d.Get(FieldUserARN).(string),
			Username: d.Get(FieldUsername).(string),
			Groups:   d.Get(FieldGroups).([]string),
		},
	}

	return mapUser, nil
}
