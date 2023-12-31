// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IamIdentityrolemappingFullName FullName.
//
// swagger:model iam.identityrolemapping.FullName
type IamIdentityrolemappingFullName struct {

	// Subject name.
	Name string `json:"name,omitempty"`

	// Org Id.
	OrgID string `json:"orgId,omitempty"`

	// Subject type.
	SubjectType string `json:"subjectType,omitempty"`
}

// Validate validates this iam identityrolemapping full name
func (m *IamIdentityrolemappingFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this iam identityrolemapping full name based on context it is used
func (m *IamIdentityrolemappingFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamIdentityrolemappingFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamIdentityrolemappingFullName) UnmarshalBinary(b []byte) error {
	var res IamIdentityrolemappingFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
