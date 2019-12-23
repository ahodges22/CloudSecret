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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CloudSecretSpec defines the desired state of CloudSecret
type CloudSecretSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Data map[string]string `json:"data,omiteempty"`
}

// CloudSecretStatus defines the observed state of CloudSecret
type CloudSecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// CloudSecret is the Schema for the cloudsecrets API
type CloudSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudSecretSpec   `json:"spec,omitempty"`
	Status CloudSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudSecretList contains a list of CloudSecret
type CloudSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSecret{}, &CloudSecretList{})
}