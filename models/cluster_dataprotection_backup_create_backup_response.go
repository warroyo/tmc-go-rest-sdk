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

// ClusterDataprotectionBackupCreateBackupResponse Response from creating a Backup.
//
// swagger:model cluster.dataprotection.backup.CreateBackupResponse
type ClusterDataprotectionBackupCreateBackupResponse struct {

	// Backup created.
	Backup *ClusterDataprotectionBackupBackup `json:"backup,omitempty"`
}

// Validate validates this cluster dataprotection backup create backup response
func (m *ClusterDataprotectionBackupCreateBackupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackupCreateBackupResponse) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	if m.Backup != nil {
		if err := m.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster dataprotection backup create backup response based on the context it is used
func (m *ClusterDataprotectionBackupCreateBackupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackupCreateBackupResponse) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {

	if m.Backup != nil {
		if err := m.Backup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionBackupCreateBackupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionBackupCreateBackupResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionBackupCreateBackupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
