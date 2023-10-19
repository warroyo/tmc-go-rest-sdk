// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationInformation Information message provides details about integration.
//
// swagger:model integration.Information
type IntegrationInformation struct {

	// The conditions required for the integration to be enabled.
	//
	// If there are no conditions required, this will be empty.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Configurations holds the static config values of an integration required to be shown on the UI.
	Configurations map[string]interface{} `json:"configurations,omitempty"`

	// Display name of the Integration.
	DisplayName string `json:"displayName,omitempty"`

	// Link to the documentation related to Integration.
	DocumentationLink string `json:"documentationLink,omitempty"`

	// Partner service endpoint of the Integration
	//
	// If there is no specific partner service endpoint, this will be empty.
	Endpoint string `json:"endpoint,omitempty"`
}

// Validate validates this integration information
func (m *IntegrationInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationInformation) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for k := range m.Conditions {

		if err := validate.Required("conditions"+"."+k, "body", m.Conditions[k]); err != nil {
			return err
		}
		if val, ok := m.Conditions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this integration information based on the context it is used
func (m *IntegrationInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationInformation) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		if val, ok := m.Conditions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationInformation) UnmarshalBinary(b []byte) error {
	var res IntegrationInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
