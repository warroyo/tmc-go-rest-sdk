// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataprotectionProviderBackuplocationSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model dataprotection.provider.backuplocation.SearchScope
type DataprotectionProviderBackuplocationSearchScope struct {

	// Filter backup locations by an assigned group name; supports globbing; default (*).
	AssignedGroupName string `json:"assignedGroupName,omitempty"`

	// Filter backup locations by a credential name; supports globbing; default (*).
	CredentialName string `json:"credentialName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified provider_name; supports globbing; default (*).
	ProviderName string `json:"providerName,omitempty"`
}

// Validate validates this dataprotection provider backuplocation search scope
func (m *DataprotectionProviderBackuplocationSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dataprotection provider backuplocation search scope based on context it is used
func (m *DataprotectionProviderBackuplocationSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationSearchScope) UnmarshalBinary(b []byte) error {
	var res DataprotectionProviderBackuplocationSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}