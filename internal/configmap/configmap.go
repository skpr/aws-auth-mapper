package configmap

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// UpdateDataWithKey updates the binary data of a ConfigMap with a given key.
func UpdateDataWithKey(ctx context.Context, client client.Client, query types.NamespacedName, key string, data []byte) error {
	configmap := &corev1.ConfigMap{}

	if err := client.Get(ctx, query, configmap); err != nil {
		return fmt.Errorf("failed to get ConfigMap: %w", err)
	}

	if len(configmap.Data) == 0 {
		configmap.Data = make(map[string]string)
	}

	configmap.Data[key] = string(data)

	if err := client.Update(ctx, configmap); err != nil {
		return fmt.Errorf("failed to update ConfigMap: %w", err)
	}

	return nil
}
