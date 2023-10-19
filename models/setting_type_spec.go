// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingTypeSpec Spec of the Setting.
//
// swagger:model setting.type.Spec
type SettingTypeSpec struct {

	// Description is used to describe this type of setting.
	Description string `json:"description,omitempty"`

	// Values schema is used to show the values that can be set by users with this type.
	ValuesSchema *SettingTypeValuesSchema `json:"valuesSchema,omitempty"`
}

// Validate validates this setting type spec
func (m *SettingTypeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValuesSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingTypeSpec) validateValuesSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.ValuesSchema) { // not required
		return nil
	}

	if m.ValuesSchema != nil {
		if err := m.ValuesSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valuesSchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valuesSchema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this setting type spec based on the context it is used
func (m *SettingTypeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValuesSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingTypeSpec) contextValidateValuesSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.ValuesSchema != nil {
		if err := m.ValuesSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valuesSchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valuesSchema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingTypeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingTypeSpec) UnmarshalBinary(b []byte) error {
	var res SettingTypeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}