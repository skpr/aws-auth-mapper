package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/skpr/aws-auth-mapper/internal/terraform/provider/aam/iamauthenticator/maprole"
	"github.com/skpr/aws-auth-mapper/internal/terraform/provider/aam/iamauthenticator/mapuser"
	"github.com/skpr/aws-auth-mapper/internal/terraform/provider/config"
)

const (
	// ResourceMapRole provides the MapRole CustomResourceDefintion.
	ResourceMapRole = "aam_iamauthenticator_v1beta1_maprole"
	// ResourceMapUser provides the MapUser CustomResourceDefintion.
	ResourceMapUser = "aam_iamauthenticator_v1beta1_mapuser"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: config.Fields(),
				ResourcesMap: map[string]*schema.Resource{
					ResourceMapRole: maprole.Resource(),
					ResourceMapUser: mapuser.Resource(),
				},
				ConfigureFunc: config.Func,
			}
		},
	})
}
