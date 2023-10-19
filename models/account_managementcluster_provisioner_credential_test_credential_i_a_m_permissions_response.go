// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse TestCredentialIAMPermissions response message.
//
// swagger:model account.managementcluster.provisioner.credential.TestCredentialIAMPermissionsResponse
type AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse struct {

	// List of allowed permissions.
	Permissions []string `json:"permissions"`

	// Unique ID of the resource. Coupled with FullName (RID), provides a stable unique
	// identifier for the resource.
	UID string `json:"uid,omitempty"`
}

// Validate validates this account managementcluster provisioner credential test credential i a m permissions response
func (m *AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account managementcluster provisioner credential test credential i a m permissions response based on context it is used
func (m *AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse) UnmarshalBinary(b []byte) error {
	var res AccountManagementclusterProvisionerCredentialTestCredentialIAMPermissionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
