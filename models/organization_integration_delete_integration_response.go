// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationIntegrationDeleteIntegrationResponse Response from deleting an Integration.
//
// swagger:model organization.integration.DeleteIntegrationResponse
type OrganizationIntegrationDeleteIntegrationResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this organization integration delete integration response
func (m *OrganizationIntegrationDeleteIntegrationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization integration delete integration response based on context it is used
func (m *OrganizationIntegrationDeleteIntegrationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationIntegrationDeleteIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationIntegrationDeleteIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res OrganizationIntegrationDeleteIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
