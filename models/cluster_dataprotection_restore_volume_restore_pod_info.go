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

// ClusterDataprotectionRestoreVolumeRestorePodInfo VolumeRestorePodInfo contains additional metadata about the pod where a restored volume was mounted.
//
// swagger:model cluster.dataprotection.restore.VolumeRestorePodInfo
type ClusterDataprotectionRestoreVolumeRestorePodInfo struct {

	// The name of the pod where the volume was mounted.
	PodName string `json:"podName,omitempty"`

	// The namespace of the pod where the volume was mounted.
	PodNamespace string `json:"podNamespace,omitempty"`

	// The name of the volume as depicted in the pod manifest.
	PodVolumeName string `json:"podVolumeName,omitempty"`

	// The phase of the pod volume restore.
	RestorePhase *ClusterDataprotectionRestoreVolumeRestorePodInfoPodVolumeRestorePhase `json:"restorePhase,omitempty"`
}

// Validate validates this cluster dataprotection restore volume restore pod info
func (m *ClusterDataprotectionRestoreVolumeRestorePodInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestorePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionRestoreVolumeRestorePodInfo) validateRestorePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.RestorePhase) { // not required
		return nil
	}

	if m.RestorePhase != nil {
		if err := m.RestorePhase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restorePhase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restorePhase")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster dataprotection restore volume restore pod info based on the context it is used
func (m *ClusterDataprotectionRestoreVolumeRestorePodInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRestorePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionRestoreVolumeRestorePodInfo) contextValidateRestorePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.RestorePhase != nil {
		if err := m.RestorePhase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restorePhase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restorePhase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreVolumeRestorePodInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreVolumeRestorePodInfo) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionRestoreVolumeRestorePodInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}