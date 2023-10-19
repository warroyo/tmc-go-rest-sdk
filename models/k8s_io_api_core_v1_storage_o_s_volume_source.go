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

// K8sIoAPICoreV1StorageOSVolumeSource Represents a StorageOS persistent volume resource.
//
// swagger:model k8s.io.api.core.v1.StorageOSVolumeSource
type K8sIoAPICoreV1StorageOSVolumeSource struct {

	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FsType string `json:"fsType,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`

	// SecretRef specifies the secret to use for obtaining the StorageOS API
	// credentials.  If not specified, default values will be attempted.
	// +optional
	SecretRef *K8sIoAPICoreV1LocalObjectReference `json:"secretRef,omitempty"`

	// VolumeName is the human-readable name of the StorageOS volume.  Volume
	// names are only unique within a namespace.
	VolumeName string `json:"volumeName,omitempty"`

	// VolumeNamespace specifies the scope of the volume within StorageOS.  If no
	// namespace is specified then the Pod's namespace will be used.  This allows the
	// Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
	// Set VolumeName to any name to override the default behaviour.
	// Set to "default" if you are not using namespaces within StorageOS.
	// Namespaces that do not pre-exist within StorageOS will be created.
	// +optional
	VolumeNamespace string `json:"volumeNamespace,omitempty"`
}

// Validate validates this k8s io api core v1 storage o s volume source
func (m *K8sIoAPICoreV1StorageOSVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1StorageOSVolumeSource) validateSecretRef(formats strfmt.Registry) error {
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

// ContextValidate validate this k8s io api core v1 storage o s volume source based on the context it is used
func (m *K8sIoAPICoreV1StorageOSVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1StorageOSVolumeSource) contextValidateSecretRef(ctx context.Context, formats strfmt.Registry) error {

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
func (m *K8sIoAPICoreV1StorageOSVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1StorageOSVolumeSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1StorageOSVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
