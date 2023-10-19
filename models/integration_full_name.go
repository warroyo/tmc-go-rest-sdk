// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegrationFullName Full name of the integration. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model integration.FullName
type IntegrationFullName struct {

	// Name of this Integration.
	Name string `json:"name,omitempty"`
}

// Validate validates this integration full name
func (m *IntegrationFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this integration full name based on context it is used
func (m *IntegrationFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationFullName) UnmarshalBinary(b []byte) error {
	var res IntegrationFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}