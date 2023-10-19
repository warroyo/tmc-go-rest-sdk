// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountCredentialDeleteCredentialResponse Response from deleting a Credential.
//
// swagger:model account.credential.DeleteCredentialResponse
type AccountCredentialDeleteCredentialResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this account credential delete credential response
func (m *AccountCredentialDeleteCredentialResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account credential delete credential response based on context it is used
func (m *AccountCredentialDeleteCredentialResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialDeleteCredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialDeleteCredentialResponse) UnmarshalBinary(b []byte) error {
	var res AccountCredentialDeleteCredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}