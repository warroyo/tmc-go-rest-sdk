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

// DataprotectionProviderBackuplocationUpdateBackupLocationRequest Request to update (overwrite) a BackupLocation.
//
// swagger:model dataprotection.provider.backuplocation.UpdateBackupLocationRequest
type DataprotectionProviderBackuplocationUpdateBackupLocationRequest struct {

	// Update BackupLocation.
	BackupLocation *DataprotectionProviderBackuplocationBackupLocation `json:"backupLocation,omitempty"`
}

// Validate validates this dataprotection provider backuplocation update backup location request
func (m *DataprotectionProviderBackuplocationUpdateBackupLocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataprotectionProviderBackuplocationUpdateBackupLocationRequest) validateBackupLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this dataprotection provider backuplocation update backup location request based on the context it is used
func (m *DataprotectionProviderBackuplocationUpdateBackupLocationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataprotectionProviderBackuplocationUpdateBackupLocationRequest) contextValidateBackupLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DataprotectionProviderBackuplocationUpdateBackupLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationUpdateBackupLocationRequest) UnmarshalBinary(b []byte) error {
	var res DataprotectionProviderBackuplocationUpdateBackupLocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
