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

// ManagementclusterUpdateManagementClusterRequest Request to update (overwrite) a ManagementCluster.
//
// swagger:model managementcluster.UpdateManagementClusterRequest
type ManagementclusterUpdateManagementClusterRequest struct {

	// Update ManagementCluster.
	ManagementCluster *ManagementclusterManagementCluster `json:"managementCluster,omitempty"`
}

// Validate validates this managementcluster update management cluster request
func (m *ManagementclusterUpdateManagementClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagementCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterUpdateManagementClusterRequest) validateManagementCluster(formats strfmt.Registry) error {
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

// ContextValidate validate this managementcluster update management cluster request based on the context it is used
func (m *ManagementclusterUpdateManagementClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManagementCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterUpdateManagementClusterRequest) contextValidateManagementCluster(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ManagementclusterUpdateManagementClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterUpdateManagementClusterRequest) UnmarshalBinary(b []byte) error {
	var res ManagementclusterUpdateManagementClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
