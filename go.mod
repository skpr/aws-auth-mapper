module github.com/skpr/aws-auth-mapper

go 1.13

require (
	github.com/aws/aws-sdk-go v1.38.71
	github.com/go-logr/logr v0.1.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/aws-iam-authenticator v0.5.3
	sigs.k8s.io/controller-runtime v0.5.0
)
