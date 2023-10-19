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

// TanzupackageTapEulaAcceptEulaRequest Request to accept an Eula.
//
// swagger:model tanzupackage.tap.eula.AcceptEulaRequest
type TanzupackageTapEulaAcceptEulaRequest struct {

	// Eula to accept.
	Eula *TanzupackageTapEulaEula `json:"eula,omitempty"`
}

// Validate validates this tanzupackage tap eula accept eula request
func (m *TanzupackageTapEulaAcceptEulaRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEula(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TanzupackageTapEulaAcceptEulaRequest) validateEula(formats strfmt.Registry) error {
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

// ContextValidate validate this tanzupackage tap eula accept eula request based on the context it is used
func (m *TanzupackageTapEulaAcceptEulaRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEula(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TanzupackageTapEulaAcceptEulaRequest) contextValidateEula(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TanzupackageTapEulaAcceptEulaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TanzupackageTapEulaAcceptEulaRequest) UnmarshalBinary(b []byte) error {
	var res TanzupackageTapEulaAcceptEulaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}