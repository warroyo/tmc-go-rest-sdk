// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountManagementclusterCredentialFullName Full name of the credential. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model account.managementcluster.credential.FullName
type AccountManagementclusterCredentialFullName struct {

	// Name of the ManagementCluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this credential.
	Name string `json:"name,omitempty"`

	// Organization ID.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this account managementcluster credential full name
func (m *AccountManagementclusterCredentialFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account managementcluster credential full name based on context it is used
func (m *AccountManagementclusterCredentialFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountManagementclusterCredentialFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountManagementclusterCredentialFullName) UnmarshalBinary(b []byte) error {
	var res AccountManagementclusterCredentialFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
