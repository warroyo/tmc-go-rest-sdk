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

// AksclusterUpdateAksClusterRequest Request to update (overwrite) an AksCluster.
//
// swagger:model akscluster.UpdateAksClusterRequest
type AksclusterUpdateAksClusterRequest struct {

	// Update AksCluster.
	AksCluster *AksclusterAksCluster `json:"aksCluster,omitempty"`
}

// Validate validates this akscluster update aks cluster request
func (m *AksclusterUpdateAksClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAksCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterUpdateAksClusterRequest) validateAksCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.AksCluster) { // not required
		return nil
	}

	if m.AksCluster != nil {
		if err := m.AksCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aksCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aksCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this akscluster update aks cluster request based on the context it is used
func (m *AksclusterUpdateAksClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAksCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterUpdateAksClusterRequest) contextValidateAksCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.AksCluster != nil {
		if err := m.AksCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aksCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aksCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterUpdateAksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterUpdateAksClusterRequest) UnmarshalBinary(b []byte) error {
	var res AksclusterUpdateAksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
