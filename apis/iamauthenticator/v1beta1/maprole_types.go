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

// MapRoleSpec defines the desired state of MapRole
type MapRoleSpec struct {
	// ARN of the role interacting with this cluster.
	RoleARN string `json:"rolearn" yaml:"rolearn"`
	// Username to assign this role while interacting with the Kubernetes cluster.
	Username string `json:"username" yaml:"username"`
	// Groups which are assigned to this role while interacting with the Kubernetes cluster.
	Groups []string `json:"groups" yaml:"groups"`
}

// MapRoleStatus defines the observed state of MapRole
type MapRoleStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +genclient
// +genclient:nonNamespaced

// MapRole is the Schema for the maproles API
type MapRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MapRoleSpec   `json:"spec,omitempty"`
	Status MapRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MapRoleList contains a list of MapRole
type MapRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MapRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MapRole{}, &MapRoleList{})
}
