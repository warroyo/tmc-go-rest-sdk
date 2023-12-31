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

// EksclusterCreateEksClusterResponse Response from creating an EksCluster.
//
// swagger:model ekscluster.CreateEksClusterResponse
type EksclusterCreateEksClusterResponse struct {

	// EksCluster created.
	EksCluster *EksclusterEksCluster `json:"eksCluster,omitempty"`
}

// Validate validates this ekscluster create eks cluster response
func (m *EksclusterCreateEksClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEksCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EksclusterCreateEksClusterResponse) validateEksCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.EksCluster) { // not required
		return nil
	}

	if m.EksCluster != nil {
		if err := m.EksCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eksCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eksCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ekscluster create eks cluster response based on the context it is used
func (m *EksclusterCreateEksClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEksCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EksclusterCreateEksClusterResponse) contextValidateEksCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.EksCluster != nil {
		if err := m.EksCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eksCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eksCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterCreateEksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterCreateEksClusterResponse) UnmarshalBinary(b []byte) error {
	var res EksclusterCreateEksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
