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

// ClusterDataprotectionBackuplocationGetBackupLocationResponse Response from getting a BackupLocation.
//
// swagger:model cluster.dataprotection.backuplocation.GetBackupLocationResponse
type ClusterDataprotectionBackuplocationGetBackupLocationResponse struct {

	// BackupLocation returned.
	BackupLocation *ClusterDataprotectionBackuplocationBackupLocation `json:"backupLocation,omitempty"`
}

// Validate validates this cluster dataprotection backuplocation get backup location response
func (m *ClusterDataprotectionBackuplocationGetBackupLocationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackuplocationGetBackupLocationResponse) validateBackupLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupLocation) { // not required
		return nil
	}

	if m.BackupLocation != nil {
		if err := m.BackupLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupLocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster dataprotection backuplocation get backup location response based on the context it is used
func (m *ClusterDataprotectionBackuplocationGetBackupLocationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackuplocationGetBackupLocationResponse) contextValidateBackupLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupLocation != nil {
		if err := m.BackupLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupLocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionBackuplocationGetBackupLocationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionBackuplocationGetBackupLocationResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionBackuplocationGetBackupLocationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
