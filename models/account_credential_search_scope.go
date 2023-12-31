// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountCredentialSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model account.credential.SearchScope
type AccountCredentialSearchScope struct {

	// Filter credentials by a capability; supports globbing; default (*).
	Capability string `json:"capability,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Filter credentials by a provider; supports globbing; default (*).
	Provider string `json:"provider,omitempty"`
}

// Validate validates this account credential search scope
func (m *AccountCredentialSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account credential search scope based on context it is used
func (m *AccountCredentialSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialSearchScope) UnmarshalBinary(b []byte) error {
	var res AccountCredentialSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
