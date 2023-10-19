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

// ManagementclusterUpdateManagementClusterResponse Response from updating a ManagementCluster.
//
// swagger:model managementcluster.UpdateManagementClusterResponse
type ManagementclusterUpdateManagementClusterResponse struct {

	// ManagementCluster updated.
	ManagementCluster *ManagementclusterManagementCluster `json:"managementCluster,omitempty"`
}

// Validate validates this managementcluster update management cluster response
func (m *ManagementclusterUpdateManagementClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagementCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterUpdateManagementClusterResponse) validateManagementCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagementCluster) { // not required
		return nil
	}

	if m.ManagementCluster != nil {
		if err := m.ManagementCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managementCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managementCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster update management cluster response based on the context it is used
func (m *ManagementclusterUpdateManagementClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManagementCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterUpdateManagementClusterResponse) contextValidateManagementCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagementCluster != nil {
		if err := m.ManagementCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managementCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managementCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterUpdateManagementClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterUpdateManagementClusterResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterUpdateManagementClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}