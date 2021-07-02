/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MapUserSpec defines the desired state of MapUser
type MapUserSpec struct {
	// ARN of the role interacting with this cluster.
	UserARN string `json:"userarn" yaml:"userarn"`
	// Username to assign this role while interacting with the Kubernetes cluster.
	Username string `json:"username" yaml:"username"`
	// Groups which are assigned to this role while interacting with the Kubernetes cluster.
	Groups []string `json:"groups" yaml:"groups"`
}

// MapUserStatus defines the observed state of MapUser
type MapUserStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +genclient
// +genclient:nonNamespaced

// MapUser is the Schema for the mapusers API
type MapUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MapUserSpec   `json:"spec,omitempty"`
	Status MapUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MapUserList contains a list of MapUser
type MapUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MapUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MapUser{}, &MapUserList{})
}
