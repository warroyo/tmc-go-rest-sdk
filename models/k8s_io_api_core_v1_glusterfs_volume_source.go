// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1GlusterfsVolumeSource Represents a Glusterfs mount that lasts the lifetime of a pod.
// Glusterfs volumes do not support ownership management or SELinux relabeling.
//
// swagger:model k8s.io.api.core.v1.GlusterfsVolumeSource
type K8sIoAPICoreV1GlusterfsVolumeSource struct {

	// EndpointsName is the endpoint name that details Glusterfs topology.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Endpoints string `json:"endpoints,omitempty"`

	// Path is the Glusterfs volume path.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path,omitempty"`

	// ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions.
	// Defaults to false.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Validate validates this k8s io api core v1 glusterfs volume source
func (m *K8sIoAPICoreV1GlusterfsVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 glusterfs volume source based on context it is used
func (m *K8sIoAPICoreV1GlusterfsVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1GlusterfsVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1GlusterfsVolumeSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1GlusterfsVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
