// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManageEksProvidereksclusterSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model manage.eks.providerekscluster.SearchScope
type ManageEksProvidereksclusterSearchScope struct {

	// Scope search to the specified credential_name; supports globbing; default (*).
	CredentialName string `json:"credentialName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified region; supports globbing; default (*).
	Region string `json:"region,omitempty"`
}

// Validate validates this manage eks providerekscluster search scope
func (m *ManageEksProvidereksclusterSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manage eks providerekscluster search scope based on context it is used
func (m *ManageEksProvidereksclusterSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManageEksProvidereksclusterSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManageEksProvidereksclusterSearchScope) UnmarshalBinary(b []byte) error {
	var res ManageEksProvidereksclusterSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
