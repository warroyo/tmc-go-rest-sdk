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

// ClusterClusterReattachRequest The request type for reattaching a cluster.
//
// swagger:model cluster.ClusterReattachRequest
type ClusterClusterReattachRequest struct {

	// Cluster full_name.
	FullName *ClusterFullName `json:"fullName,omitempty"`

	// Optional image registry is the name of the Image Registry Config
	// to be used for the cluster.
	// If set empty, no image registry config will be used.
	// If set non-empty, defined image registry config will be used.
	// If left unset, existing image registry config will be used.
	ImageRegistry string `json:"imageRegistry,omitempty"`

	// Optional proxy name is the name of the Proxy Config
	// to be used for the cluster.
	// If set empty, existing proxy config will be used.
	// If set non-empty, defined proxy config will be used.
	// If left unset, no proxy config will be used.
	ProxyName string `json:"proxyName,omitempty"`
}

// Validate validates this cluster cluster reattach request
func (m *ClusterClusterReattachRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterClusterReattachRequest) validateFullName(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster cluster reattach request based on the context it is used
func (m *ClusterClusterReattachRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterClusterReattachRequest) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterClusterReattachRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterClusterReattachRequest) UnmarshalBinary(b []byte) error {
	var res ClusterClusterReattachRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}