package eks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skpr/aws-auth-mapper/internal/aws/sts/mock"
)

func TestGenerateToken(t *testing.T) {
	stsClient := mock.NewPresignClient()
	generator := NewSTSTokenGenerator(stsClient)
	token, err := generator.GenerateToken(context.TODO(), "foo")
	assert.NoError(t, err)
	assert.Equal(t, "k8s-aws-v1.aHR0cDovL2V4YW1wbGUvY29t", token)
}
