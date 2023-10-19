// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingTypeSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model setting.type.SearchScope
type SettingTypeSearchScope struct {

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`
}

// Validate validates this setting type search scope
func (m *SettingTypeSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this setting type search scope based on context it is used
func (m *SettingTypeSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingTypeSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingTypeSearchScope) UnmarshalBinary(b []byte) error {
	var res SettingTypeSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}