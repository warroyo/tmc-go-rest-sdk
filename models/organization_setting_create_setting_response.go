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

// OrganizationSettingCreateSettingResponse Response from creating a Setting.
//
// swagger:model organization.setting.CreateSettingResponse
type OrganizationSettingCreateSettingResponse struct {

	// Setting created.
	Setting *OrganizationSettingSetting `json:"setting,omitempty"`
}

// Validate validates this organization setting create setting response
func (m *OrganizationSettingCreateSettingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSetting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationSettingCreateSettingResponse) validateSetting(formats strfmt.Registry) error {
	if swag.IsZero(m.Setting) { // not required
		return nil
	}

	if m.Setting != nil {
		if err := m.Setting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("setting")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this organization setting create setting response based on the context it is used
func (m *OrganizationSettingCreateSettingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSetting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationSettingCreateSettingResponse) contextValidateSetting(ctx context.Context, formats strfmt.Registry) error {

	if m.Setting != nil {
		if err := m.Setting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("setting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationSettingCreateSettingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationSettingCreateSettingResponse) UnmarshalBinary(b []byte) error {
	var res OrganizationSettingCreateSettingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}