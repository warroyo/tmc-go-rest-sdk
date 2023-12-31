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

// ClustergroupTestClusterGroupIAMPermissionsRequest TestClusterGroupIAMPermissions request message.
//
// swagger:model clustergroup.TestClusterGroupIAMPermissionsRequest
type ClustergroupTestClusterGroupIAMPermissionsRequest struct {

	// ClusterGroup full_name.
	FullName *ClustergroupFullName `json:"fullName,omitempty"`

	// List of permissions to test.
	Permissions []string `json:"permissions"`
}

// Validate validates this clustergroup test cluster group i a m permissions request
func (m *ClustergroupTestClusterGroupIAMPermissionsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupTestClusterGroupIAMPermissionsRequest) validateFullName(formats strfmt.Registry) error {
	if swag.IsZero(m.FullName) { // not required
		return nil
	}

	if m.FullName != nil {
		if err := m.FullName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup test cluster group i a m permissions request based on the context it is used
func (m *ClustergroupTestClusterGroupIAMPermissionsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupTestClusterGroupIAMPermissionsRequest) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

	if m.FullName != nil {
		if err := m.FullName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupTestClusterGroupIAMPermissionsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupTestClusterGroupIAMPermissionsRequest) UnmarshalBinary(b []byte) error {
	var res ClustergroupTestClusterGroupIAMPermissionsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
