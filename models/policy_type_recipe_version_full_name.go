// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTypeRecipeVersionFullName Full name of the policy recipe version. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model policy.type.recipe.version.FullName
type PolicyTypeRecipeVersionFullName struct {

	// Name of policy recipe version.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of policy recipe.
	RecipeName string `json:"recipeName,omitempty"`

	// Name of policy type.
	TypeName string `json:"typeName,omitempty"`
}

// Validate validates this policy type recipe version full name
func (m *PolicyTypeRecipeVersionFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy type recipe version full name based on context it is used
func (m *PolicyTypeRecipeVersionFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTypeRecipeVersionFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTypeRecipeVersionFullName) UnmarshalBinary(b []byte) error {
	var res PolicyTypeRecipeVersionFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}