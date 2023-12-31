// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterTKGServiceVsphereVolume TKGServiceVsphereVolume defines a Persistent Volume Claim attached for the workload cluster.
//
// swagger:model common.cluster.TKGServiceVsphereVolume
type CommonClusterTKGServiceVsphereVolume struct {

	// Volume capacity is in gib.
	Capacity float32 `json:"capacity,omitempty"`

	// MountPath is the directory where the volume device is to be mounted.
	MountPath string `json:"mountPath,omitempty"`

	// Volume name.
	Name string `json:"name,omitempty"`

	// Storage class for PVC
	// If omitted, default storage class will be used for the disks.
	StorageClass string `json:"storageClass,omitempty"`
}

// Validate validates this common cluster t k g service vsphere volume
func (m *CommonClusterTKGServiceVsphereVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster t k g service vsphere volume based on context it is used
func (m *CommonClusterTKGServiceVsphereVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGServiceVsphereVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGServiceVsphereVolume) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGServiceVsphereVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
