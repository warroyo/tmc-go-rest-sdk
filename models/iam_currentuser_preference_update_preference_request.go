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

// IamCurrentuserPreferenceUpdatePreferenceRequest Request to update (overwrite) a Preference.
//
// swagger:model iam.currentuser.preference.UpdatePreferenceRequest
type IamCurrentuserPreferenceUpdatePreferenceRequest struct {

	// Update Preference.
	Preference *IamCurrentuserPreferencePreference `json:"preference,omitempty"`
}

// Validate validates this iam currentuser preference update preference request
func (m *IamCurrentuserPreferenceUpdatePreferenceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamCurrentuserPreferenceUpdatePreferenceRequest) validatePreference(formats strfmt.Registry) error {
	if swag.IsZero(m.Preference) { // not required
		return nil
	}

	if m.Preference != nil {
		if err := m.Preference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iam currentuser preference update preference request based on the context it is used
func (m *IamCurrentuserPreferenceUpdatePreferenceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamCurrentuserPreferenceUpdatePreferenceRequest) contextValidatePreference(ctx context.Context, formats strfmt.Registry) error {

	if m.Preference != nil {
		if err := m.Preference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamCurrentuserPreferenceUpdatePreferenceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamCurrentuserPreferenceUpdatePreferenceRequest) UnmarshalBinary(b []byte) error {
	var res IamCurrentuserPreferenceUpdatePreferenceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
