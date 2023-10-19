// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonSettingSpec Common Spec for the Setting Instance.
//
// swagger:model common.setting.Spec
type CommonSettingSpec struct {

	// Input with the setting values.
	// These should follow the schema of the provided setting type.
	Values interface{} `json:"values,omitempty"`
}

// Validate validates this common setting spec
func (m *CommonSettingSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common setting spec based on context it is used
func (m *CommonSettingSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonSettingSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonSettingSpec) UnmarshalBinary(b []byte) error {
	var res CommonSettingSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
