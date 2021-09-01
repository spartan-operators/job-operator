/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VrTestJobSpec defines the desired state of VrTestJob
type VrTestJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//Image name to be spawn as a job within the cluster
	Image string `json:"image,omitempty"`

	// How many instances should we run for a singe image name
	Retries int32 `json:"retries,omitempty"`
}

// VrTestJobStatus defines the observed state of VrTestJob
type VrTestJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VrTestJob is the Schema for the vrtestjobs API
type VrTestJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VrTestJobSpec   `json:"spec,omitempty"`
	Status VrTestJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VrTestJobList contains a list of VrTestJob
type VrTestJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VrTestJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VrTestJob{}, &VrTestJobList{})
}
