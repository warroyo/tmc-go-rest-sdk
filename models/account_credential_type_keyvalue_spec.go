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

// AccountCredentialTypeKeyvalueSpec KeyValue Type credential stored in Account Manager.
//
// swagger:model account.credential.type.keyvalue.Spec
type AccountCredentialTypeKeyvalueSpec struct {

	// Data contains the secret data. Each key must consist of alphanumeric
	// characters, '-', '_' or '.'.
	Data map[string]strfmt.Base64 `json:"data,omitempty"`

	// Type of Secret.
	// The default value is SECRET_TYPE_OPAQUE.
	Type *AccountCredentialTypeKeyvalueSpecSecretType `json:"type,omitempty"`
}

// Validate validates this account credential type keyvalue spec
func (m *AccountCredentialTypeKeyvalueSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialTypeKeyvalueSpec) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account credential type keyvalue spec based on the context it is used
func (m *AccountCredentialTypeKeyvalueSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialTypeKeyvalueSpec) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialTypeKeyvalueSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialTypeKeyvalueSpec) UnmarshalBinary(b []byte) error {
	var res AccountCredentialTypeKeyvalueSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}