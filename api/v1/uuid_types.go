/*

 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UUIDSpec defines the desired state of UUID
type UUIDSpec struct {
	UUID      string `json:"uuid,omitempty"`
	RandomInt int    `json:"random_int,omitempty"`
}

// UUIDStatus defines the observed state of UUID
type UUIDStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// UUID is the Schema for the uuids API
type UUID struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UUIDSpec   `json:"spec,omitempty"`
	Status UUIDStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UUIDList contains a list of UUID
type UUIDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UUID `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UUID{}, &UUIDList{})
}
