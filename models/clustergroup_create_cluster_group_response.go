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

// ClustergroupCreateClusterGroupResponse Response from creating a ClusterGroup.
//
// swagger:model clustergroup.CreateClusterGroupResponse
type ClustergroupCreateClusterGroupResponse struct {

	// ClusterGroup created.
	ClusterGroup *ClustergroupClusterGroup `json:"clusterGroup,omitempty"`
}

// Validate validates this clustergroup create cluster group response
func (m *ClustergroupCreateClusterGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupCreateClusterGroupResponse) validateClusterGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterGroup) { // not required
		return nil
	}

	if m.ClusterGroup != nil {
		if err := m.ClusterGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup create cluster group response based on the context it is used
func (m *ClustergroupCreateClusterGroupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupCreateClusterGroupResponse) contextValidateClusterGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterGroup != nil {
		if err := m.ClusterGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupCreateClusterGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupCreateClusterGroupResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupCreateClusterGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
