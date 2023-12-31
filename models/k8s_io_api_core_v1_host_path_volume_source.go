// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1HostPathVolumeSource Represents a host path mapped into a pod.
// Host path volumes do not support ownership management or SELinux relabeling.
//
// swagger:model k8s.io.api.core.v1.HostPathVolumeSource
type K8sIoAPICoreV1HostPathVolumeSource struct {

	// Path of the directory on the host.
	// If the path is a symlink, it will follow the link to the real path.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Path string `json:"path,omitempty"`

	// Type for HostPath Volume
	// Defaults to ""
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	Type string `json:"type,omitempty"`
}

// Validate validates this k8s io api core v1 host path volume source
func (m *K8sIoAPICoreV1HostPathVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 host path volume source based on context it is used
func (m *K8sIoAPICoreV1HostPathVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1HostPathVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1HostPathVolumeSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1HostPathVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
