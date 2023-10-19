// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IamRoleFullName Full name for role.
//
// swagger:model iam.role.FullName
type IamRoleFullName struct {

	// Name of the role.
	Name string `json:"name,omitempty"`

	// Org Id.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this iam role full name
func (m *IamRoleFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this iam role full name based on context it is used
func (m *IamRoleFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamRoleFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamRoleFullName) UnmarshalBinary(b []byte) error {
	var res IamRoleFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}