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

// ClusterTestClusterIAMPermissionsRequest TestClusterIAMPermissions request message.
//
// swagger:model cluster.TestClusterIAMPermissionsRequest
type ClusterTestClusterIAMPermissionsRequest struct {

	// Cluster full_name.
	FullName *ClusterFullName `json:"fullName,omitempty"`

	// List of permissions to test.
	Permissions []string `json:"permissions"`
}

// Validate validates this cluster test cluster i a m permissions request
func (m *ClusterTestClusterIAMPermissionsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTestClusterIAMPermissionsRequest) validateFullName(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster test cluster i a m permissions request based on the context it is used
func (m *ClusterTestClusterIAMPermissionsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTestClusterIAMPermissionsRequest) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterTestClusterIAMPermissionsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTestClusterIAMPermissionsRequest) UnmarshalBinary(b []byte) error {
	var res ClusterTestClusterIAMPermissionsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
