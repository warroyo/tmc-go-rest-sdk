// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTypeRecipeSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model policy.type.recipe.SearchScope
type PolicyTypeRecipeSearchScope struct {

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified type_name; supports globbing; default (*).
	TypeName string `json:"typeName,omitempty"`
}

// Validate validates this policy type recipe search scope
func (m *PolicyTypeRecipeSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy type recipe search scope based on context it is used
func (m *PolicyTypeRecipeSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTypeRecipeSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTypeRecipeSearchScope) UnmarshalBinary(b []byte) error {
	var res PolicyTypeRecipeSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
