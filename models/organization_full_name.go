// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationFullName Full name of the organization. This includes the org_id.
//
// swagger:model organization.FullName
type OrganizationFullName struct {

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this organization full name
func (m *OrganizationFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization full name based on context it is used
func (m *OrganizationFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationFullName) UnmarshalBinary(b []byte) error {
	var res OrganizationFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
