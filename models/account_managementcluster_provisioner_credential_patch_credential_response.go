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

// AccountManagementclusterProvisionerCredentialPatchCredentialResponse Response from patching a Credential.
//
// swagger:model account.managementcluster.provisioner.credential.PatchCredentialResponse
type AccountManagementclusterProvisionerCredentialPatchCredentialResponse struct {

	// Credential patched.
	Credential *AccountManagementclusterProvisionerCredentialCredential `json:"credential,omitempty"`
}

// Validate validates this account managementcluster provisioner credential patch credential response
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountManagementclusterProvisionerCredentialPatchCredentialResponse) validateCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account managementcluster provisioner credential patch credential response based on the context it is used
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountManagementclusterProvisionerCredentialPatchCredentialResponse) contextValidateCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.Credential != nil {
		if err := m.Credential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialResponse) UnmarshalBinary(b []byte) error {
	var res AccountManagementclusterProvisionerCredentialPatchCredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
