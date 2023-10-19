// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingEffectiveFullName Full name of the Effective Setting.
//
// swagger:model setting.effective.FullName
type SettingEffectiveFullName struct {

	// Name of the effective setting.
	// This will be the unique identifier (same as meta.UID)
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this setting effective full name
func (m *SettingEffectiveFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this setting effective full name based on context it is used
func (m *SettingEffectiveFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingEffectiveFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingEffectiveFullName) UnmarshalBinary(b []byte) error {
	var res SettingEffectiveFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}