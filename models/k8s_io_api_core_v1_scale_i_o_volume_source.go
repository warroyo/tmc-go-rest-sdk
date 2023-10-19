// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1ScaleIOVolumeSource ScaleIOVolumeSource represents a persistent ScaleIO volume
//
// swagger:model k8s.io.api.core.v1.ScaleIOVolumeSource
type K8sIoAPICoreV1ScaleIOVolumeSource struct {

	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// Default is "xfs".
	// +optional
	FsType string `json:"fsType,omitempty"`

	// The host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway,omitempty"`

	// The name of the ScaleIO Protection Domain for the configured storage.
	// +optional
	ProtectionDomain string `json:"protectionDomain,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`

	// SecretRef references to the secret for ScaleIO user and other
	// sensitive information. If this is not provided, Login operation will fail.
	SecretRef *K8sIoAPICoreV1LocalObjectReference `json:"secretRef,omitempty"`

	// Flag to enable/disable SSL communication with Gateway, default false
	// +optional
	SslEnabled bool `json:"sslEnabled,omitempty"`

	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	// Default is ThinProvisioned.
	// +optional
	StorageMode string `json:"storageMode,omitempty"`

	// The ScaleIO Storage Pool associated with the protection domain.
	// +optional
	StoragePool string `json:"storagePool,omitempty"`

	// The name of the storage system as configured in ScaleIO.
	System string `json:"system,omitempty"`

	// The name of a volume already created in the ScaleIO system
	// that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty"`
}

// Validate validates this k8s io api core v1 scale i o volume source
func (m *K8sIoAPICoreV1ScaleIOVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ScaleIOVolumeSource) validateSecretRef(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretRef) { // not required
		return nil
	}

	if m.SecretRef != nil {
		if err := m.SecretRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 scale i o volume source based on the context it is used
func (m *K8sIoAPICoreV1ScaleIOVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ScaleIOVolumeSource) contextValidateSecretRef(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretRef != nil {
		if err := m.SecretRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ScaleIOVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ScaleIOVolumeSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ScaleIOVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}