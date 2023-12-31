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

// K8sIoAPICoreV1CephFSVolumeSource Represents a Ceph Filesystem mount that lasts the lifetime of a pod
// Cephfs volumes do not support ownership management or SELinux relabeling.
//
// swagger:model k8s.io.api.core.v1.CephFSVolumeSource
type K8sIoAPICoreV1CephFSVolumeSource struct {

	// Required: Monitors is a collection of Ceph monitors
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors"`

	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	// +optional
	Path string `json:"path,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`

	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretFile string `json:"secretFile,omitempty"`

	// Optional: SecretRef is reference to the authentication secret for User, default is empty.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretRef *K8sIoAPICoreV1LocalObjectReference `json:"secretRef,omitempty"`

	// Optional: User is the rados user name, default is admin
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	User string `json:"user,omitempty"`
}

// Validate validates this k8s io api core v1 ceph f s volume source
func (m *K8sIoAPICoreV1CephFSVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1CephFSVolumeSource) validateSecretRef(formats strfmt.Registry) error {
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

// ContextValidate validate this k8s io api core v1 ceph f s volume source based on the context it is used
func (m *K8sIoAPICoreV1CephFSVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1CephFSVolumeSource) contextValidateSecretRef(ctx context.Context, formats strfmt.Registry) error {

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
func (m *K8sIoAPICoreV1CephFSVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1CephFSVolumeSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1CephFSVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
