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

// ClusterDataprotectionRestoreGetRestoreResponse Response from getting a Restore.
//
// swagger:model cluster.dataprotection.restore.GetRestoreResponse
type ClusterDataprotectionRestoreGetRestoreResponse struct {

	// Restore returned.
	Restore *ClusterDataprotectionRestoreRestore `json:"restore,omitempty"`
}

// Validate validates this cluster dataprotection restore get restore response
func (m *ClusterDataprotectionRestoreGetRestoreResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionRestoreGetRestoreResponse) validateRestore(formats strfmt.Registry) error {
	if swag.IsZero(m.Restore) { // not required
		return nil
	}

	if m.Restore != nil {
		if err := m.Restore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restore")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster dataprotection restore get restore response based on the context it is used
func (m *ClusterDataprotectionRestoreGetRestoreResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRestore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionRestoreGetRestoreResponse) contextValidateRestore(ctx context.Context, formats strfmt.Registry) error {

	if m.Restore != nil {
		if err := m.Restore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreGetRestoreResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreGetRestoreResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionRestoreGetRestoreResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}