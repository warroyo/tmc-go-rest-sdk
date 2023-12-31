// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterDataprotectionBackuplocationListBackupLocationsResponse Response from listing BackupLocations.
//
// swagger:model cluster.dataprotection.backuplocation.ListBackupLocationsResponse
type ClusterDataprotectionBackuplocationListBackupLocationsResponse struct {

	// List of backuplocations.
	BackupLocations []*ClusterDataprotectionBackuplocationBackupLocation `json:"backupLocations"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this cluster dataprotection backuplocation list backup locations response
func (m *ClusterDataprotectionBackuplocationListBackupLocationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackuplocationListBackupLocationsResponse) validateBackupLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupLocations) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupLocations); i++ {
		if swag.IsZero(m.BackupLocations[i]) { // not required
			continue
		}

		if m.BackupLocations[i] != nil {
			if err := m.BackupLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster dataprotection backuplocation list backup locations response based on the context it is used
func (m *ClusterDataprotectionBackuplocationListBackupLocationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackuplocationListBackupLocationsResponse) contextValidateBackupLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupLocations); i++ {

		if m.BackupLocations[i] != nil {
			if err := m.BackupLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionBackuplocationListBackupLocationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionBackuplocationListBackupLocationsResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionBackuplocationListBackupLocationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
