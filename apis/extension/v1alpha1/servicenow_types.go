/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Servicenow struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicenowSpec   `json:"spec,omitempty"`
	Status            ServicenowStatus `json:"status,omitempty"`
}

type ServicenowSpec struct {
	State *ServicenowSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServicenowSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ServicenowSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EndpointURL      *string  `json:"-" sensitive:"true" tf:"endpoint_url"`
	ExtensionObjects []string `json:"extensionObjects" tf:"extension_objects"`
	ExtensionSchema  *string  `json:"extensionSchema" tf:"extension_schema"`
	// +optional
	HtmlURL *string `json:"htmlURL,omitempty" tf:"html_url"`
	// +optional
	Name         *string `json:"name,omitempty" tf:"name"`
	Referer      *string `json:"referer" tf:"referer"`
	SnowPassword *string `json:"-" sensitive:"true" tf:"snow_password"`
	SnowUser     *string `json:"snowUser" tf:"snow_user"`
	SyncOptions  *string `json:"syncOptions" tf:"sync_options"`
	Target       *string `json:"target" tf:"target"`
	TaskType     *string `json:"taskType" tf:"task_type"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServicenowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicenowList is a list of Servicenows
type ServicenowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Servicenow CRD objects
	Items []Servicenow `json:"items,omitempty"`
}