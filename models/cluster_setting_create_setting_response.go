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

// ClusterSettingCreateSettingResponse Response from creating a Setting.
//
// swagger:model cluster.setting.CreateSettingResponse
type ClusterSettingCreateSettingResponse struct {

	// Setting created.
	Setting *ClusterSettingSetting `json:"setting,omitempty"`
}

// Validate validates this cluster setting create setting response
func (m *ClusterSettingCreateSettingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSetting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingCreateSettingResponse) validateSetting(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster setting create setting response based on the context it is used
func (m *ClusterSettingCreateSettingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSetting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingCreateSettingResponse) contextValidateSetting(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterSettingCreateSettingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettingCreateSettingResponse) UnmarshalBinary(b []byte) error {
	var res ClusterSettingCreateSettingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
