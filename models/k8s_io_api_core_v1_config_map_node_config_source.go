// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1ConfigMapNodeConfigSource ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node.
//
// swagger:model k8s.io.api.core.v1.ConfigMapNodeConfigSource
type K8sIoAPICoreV1ConfigMapNodeConfigSource struct {

	// KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure
	// This field is required in all cases.
	KubeletConfigKey string `json:"kubeletConfigKey,omitempty"`

	// Name is the metadata.name of the referenced ConfigMap.
	// This field is required in all cases.
	Name string `json:"name,omitempty"`

	// Namespace is the metadata.namespace of the referenced ConfigMap.
	// This field is required in all cases.
	Namespace string `json:"namespace,omitempty"`

	// ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap.
	// This field is forbidden in Node.Spec, and required in Node.Status.
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID is the metadata.UID of the referenced ConfigMap.
	// This field is forbidden in Node.Spec, and required in Node.Status.
	// +optional
	UID string `json:"uid,omitempty"`
}

// Validate validates this k8s io api core v1 config map node config source
func (m *K8sIoAPICoreV1ConfigMapNodeConfigSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 config map node config source based on context it is used
func (m *K8sIoAPICoreV1ConfigMapNodeConfigSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ConfigMapNodeConfigSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ConfigMapNodeConfigSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ConfigMapNodeConfigSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
