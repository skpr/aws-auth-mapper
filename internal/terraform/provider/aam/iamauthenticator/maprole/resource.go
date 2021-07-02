package maprole

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName is used to identify the name of a Kubernetes object.
	FieldName = "name"
	// FieldRoleARN is used define which IAM Role is mapped to an internal Kubernetes user.
	FieldRoleARN = "role_arn"
	// FieldUsername is used to identify when accessing the Kubernetes cluster.
	FieldUsername = "username"
	// FieldGroup is used to map groups to this IAM Role with interacting with the Kubernetes cluster.
	FieldGroups = "groups"
)

// Resource returns this MapRole definition.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			FieldName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldRoleARN: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldUsername: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldGroups: {
				Elem:     schema.TypeList,
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
