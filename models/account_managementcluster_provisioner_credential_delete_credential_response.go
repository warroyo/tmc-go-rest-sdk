// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountManagementclusterProvisionerCredentialDeleteCredentialResponse Response from deleting a Credential.
//
// swagger:model account.managementcluster.provisioner.credential.DeleteCredentialResponse
type AccountManagementclusterProvisionerCredentialDeleteCredentialResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this account managementcluster provisioner credential delete credential response
func (m *AccountManagementclusterProvisionerCredentialDeleteCredentialResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account managementcluster provisioner credential delete credential response based on context it is used
func (m *AccountManagementclusterProvisionerCredentialDeleteCredentialResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialDeleteCredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialDeleteCredentialResponse) UnmarshalBinary(b []byte) error {
	var res AccountManagementclusterProvisionerCredentialDeleteCredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
