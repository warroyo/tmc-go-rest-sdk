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

// TanzupackageTapEulaAcceptEulaResponse Response from accepting an Eula.
//
// swagger:model tanzupackage.tap.eula.AcceptEulaResponse
type TanzupackageTapEulaAcceptEulaResponse struct {

	// Eula accepted.
	Eula *TanzupackageTapEulaEula `json:"eula,omitempty"`
}

// Validate validates this tanzupackage tap eula accept eula response
func (m *TanzupackageTapEulaAcceptEulaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEula(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TanzupackageTapEulaAcceptEulaResponse) validateEula(formats strfmt.Registry) error {
	if swag.IsZero(m.Eula) { // not required
		return nil
	}

	if m.Eula != nil {
		if err := m.Eula.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eula")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eula")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tanzupackage tap eula accept eula response based on the context it is used
func (m *TanzupackageTapEulaAcceptEulaResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEula(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TanzupackageTapEulaAcceptEulaResponse) contextValidateEula(ctx context.Context, formats strfmt.Registry) error {

	if m.Eula != nil {
		if err := m.Eula.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eula")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eula")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TanzupackageTapEulaAcceptEulaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TanzupackageTapEulaAcceptEulaResponse) UnmarshalBinary(b []byte) error {
	var res TanzupackageTapEulaAcceptEulaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
