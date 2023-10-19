// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingEffectiveListEffectiveResponse Response from listing Effective.
//
// swagger:model setting.effective.ListEffectiveResponse
type SettingEffectiveListEffectiveResponse struct {

	// List of effective.
	Effective []*SettingEffectiveEffective `json:"effective"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this setting effective list effective response
func (m *SettingEffectiveListEffectiveResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffective(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingEffectiveListEffectiveResponse) validateEffective(formats strfmt.Registry) error {
	if swag.IsZero(m.Effective) { // not required
		return nil
	}

	for i := 0; i < len(m.Effective); i++ {
		if swag.IsZero(m.Effective[i]) { // not required
			continue
		}

		if m.Effective[i] != nil {
			if err := m.Effective[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("effective" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("effective" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this setting effective list effective response based on the context it is used
func (m *SettingEffectiveListEffectiveResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEffective(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingEffectiveListEffectiveResponse) contextValidateEffective(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Effective); i++ {

		if m.Effective[i] != nil {
			if err := m.Effective[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("effective" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("effective" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingEffectiveListEffectiveResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingEffectiveListEffectiveResponse) UnmarshalBinary(b []byte) error {
	var res SettingEffectiveListEffectiveResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}