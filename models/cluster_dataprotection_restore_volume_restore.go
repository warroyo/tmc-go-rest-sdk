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

// ClusterDataprotectionRestoreVolumeRestore VolumeRestore contains metadata about a particular volume restored.
//
// swagger:model cluster.dataprotection.restore.VolumeRestore
type ClusterDataprotectionRestoreVolumeRestore struct {

	// The method used to perform the volume restore.
	Method *ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod `json:"method,omitempty"`

	// Additional metadata about the pod where the volume was mounted.
	// Only present for file system restores.
	PodInfo *ClusterDataprotectionRestoreVolumeRestorePodInfo `json:"podInfo,omitempty"`

	// The name of the persistent volume.
	PvName string `json:"pvName,omitempty"`

	// The name of the persistent volume claim.
	PvcName string `json:"pvcName,omitempty"`

	// The namespace of the persistent volume claim.
	PvcNamespace string `json:"pvcNamespace,omitempty"`

	// The name of the storage class used by the persistent volume.
	ScName string `json:"scName,omitempty"`

	// The complete size of the snapshot in bytes.
	Size string `json:"size,omitempty"`
}

// Validate validates this cluster dataprotection restore volume restore
func (m *ClusterDataprotectionRestoreVolumeRestore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionRestoreVolumeRestore) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if m.Method != nil {
		if err := m.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionRestoreVolumeRestore) validatePodInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PodInfo) { // not required
		return nil
	}

	if m.PodInfo != nil {
		if err := m.PodInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster dataprotection restore volume restore based on the context it is used
func (m *ClusterDataprotectionRestoreVolumeRestore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionRestoreVolumeRestore) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.Method != nil {
		if err := m.Method.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionRestoreVolumeRestore) contextValidatePodInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PodInfo != nil {
		if err := m.PodInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreVolumeRestore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreVolumeRestore) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionRestoreVolumeRestore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
