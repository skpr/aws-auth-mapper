package maprole

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
)

// Create the MapRole resource.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*clientset.Clientset)

	mapRole, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.IamauthenticatorV1beta1().MapRoles().Create(&mapRole)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(out.ObjectMeta.Name)

	return nil
}
