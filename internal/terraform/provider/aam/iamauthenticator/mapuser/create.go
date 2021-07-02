package mapuser

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Create the MapRole resource.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*clientset.Clientset)

	mapUser, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.IamauthenticatorV1beta1().MapUsers().Create(&mapUser)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(out.ObjectMeta.Name)

	return nil
}
