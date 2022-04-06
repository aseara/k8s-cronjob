/*
Copyright 2022 aseara.
*/

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//+kubebuilder:object:root=true

// ProjectConfig is the Schema for the projectconfigs API
type ProjectConfig struct {
	metav1.TypeMeta `json:",inline"`

	// ControllerManagerConfiguration returns the contfigurations for controllers
	cfg.ControllerManagerConfigurationSpec `json:",inline"`

	ClusterName string `json:"clusterName,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ProjectConfig{})
}
