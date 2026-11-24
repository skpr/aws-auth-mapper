package config

import (
	"bytes"
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/skpr/aws-auth-mapper/internal/clientset"
	"github.com/skpr/aws-auth-mapper/internal/eks"
)

// Func which configures the Kubernetes provider.
// https://github.com/terraform-providers/terraform-provider-kubernetes/blob/master/kubernetes/provider.go
func Func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if v, ok := d.GetOk(FieldKubeConfig); ok {
		cfg, err := clientcmd.BuildConfigFromFlags("", v.(string))
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return newClientset(cfg)
	}

	cfg := &rest.Config{}

	// Overriding with static configuration
	cfg.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform")

	if v, ok := d.GetOk(FieldHost); ok {
		cfg.Host = v.(string)
	}

	if v, ok := d.GetOk(FieldInsecure); ok {
		cfg.Insecure = v.(bool)
	}

	if v, ok := d.GetOk(FieldClusterCACertificate); ok {
		cfg.CAData = bytes.NewBufferString(v.(string)).Bytes()
	}

	if v, ok := d.GetOk(FieldEKSCluster); ok {
		var (
			region  string
			profile string
			cluster = v.(string)
		)

		if v, ok := d.GetOk(FieldAWSProfile); ok {
			profile = v.(string)
		}

		if v, ok := d.GetOk(FieldAWSRegion); ok {
			region = v.(string)
		}

		token, err := eksToken(ctx, region, cluster, profile)
		if err != nil {
			return nil, diag.FromErr(errors.Wrap(err, "failed to EKS bearer token"))
		}

		cfg.BearerToken = token
	}

	return newClientset(cfg)
}

// Helper function to generate an EKS Kubernetes client.
func eksToken(ctx context.Context, region, cluster, profile string) (string, error) {
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(region),
		awsconfig.WithSharedConfigProfile(profile),
	)
	if err != nil {
		return "", fmt.Errorf("failed to get aws config: %w", err)
	}

	var (
		stsClient        = sts.NewFromConfig(cfg)
		stsPresignClient = sts.NewPresignClient(stsClient)
	)

	gen := eks.NewSTSTokenGenerator(stsPresignClient)
	if err != nil {
		return "", fmt.Errorf("failed to create token generator: %w", err)
	}

	token, err := gen.GenerateToken(ctx, cluster)
	if err != nil {
		return "", fmt.Errorf("failed to get sts token: %w", err)
	}

	return token, nil
}

// Helper function to wrap clientset errors as Terraform diagnostics.
func newClientset(cfg *rest.Config) (*clientset.Clientset, diag.Diagnostics) {
	cs, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return cs, nil
}
