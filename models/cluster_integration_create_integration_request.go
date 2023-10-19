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

// ClusterIntegrationCreateIntegrationRequest Request to create an Integration.
//
// swagger:model cluster.integration.CreateIntegrationRequest
type ClusterIntegrationCreateIntegrationRequest struct {

	// Integration to create.
	Integration *ClusterIntegrationIntegration `json:"integration,omitempty"`
}

// Validate validates this cluster integration create integration request
func (m *ClusterIntegrationCreateIntegrationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterIntegrationCreateIntegrationRequest) validateIntegration(formats strfmt.Registry) error {
	if swag.IsZero(m.Integration) { // not required
		return nil
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster integration create integration request based on the context it is used
func (m *ClusterIntegrationCreateIntegrationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterIntegrationCreateIntegrationRequest) contextValidateIntegration(ctx context.Context, formats strfmt.Registry) error {

	if m.Integration != nil {
		if err := m.Integration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterIntegrationCreateIntegrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterIntegrationCreateIntegrationRequest) UnmarshalBinary(b []byte) error {
	var res ClusterIntegrationCreateIntegrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
