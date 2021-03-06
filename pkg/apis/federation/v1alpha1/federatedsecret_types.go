/*
Copyright 2018 The Federation v2 Authors.

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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/kubernetes-sigs/federation-v2/pkg/apis/federation"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedSecret
// +k8s:openapi-gen=true
// +resource:path=federatedsecrets,strategy=FederatedSecretStrategy
type FederatedSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedSecretSpec   `json:"spec,omitempty"`
	Status FederatedSecretStatus `json:"status,omitempty"`
}

// FederatedSecretSpec defines the desired state of FederatedSecret
type FederatedSecretSpec struct {
	// Template to derive per-cluster secret from
	Template corev1.Secret `json:"template,omitempty"`
}

// FederatedSecretStatus defines the observed state of FederatedSecret
type FederatedSecretStatus struct {
}

// Validate checks that an instance of FederatedSecret is well formed
func (FederatedSecretStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*federation.FederatedSecret)
	log.Printf("Validating fields for FederatedSecret %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default FederatedSecret field values
func (FederatedSecretSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*FederatedSecret)
	// set default field values here
	log.Printf("Defaulting fields for FederatedSecret %s\n", obj.Name)
}
