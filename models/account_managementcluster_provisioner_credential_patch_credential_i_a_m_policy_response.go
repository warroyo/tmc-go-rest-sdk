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

// AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse PatchCredentialIAMPolicy response message.
//
// swagger:model account.managementcluster.provisioner.credential.PatchCredentialIAMPolicyResponse
type AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse struct {

	// New policy object.
	Policy *VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// Validate validates this account managementcluster provisioner credential patch credential i a m policy response
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account managementcluster provisioner credential patch credential i a m policy response based on the context it is used
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {
		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res AccountManagementclusterProvisionerCredentialPatchCredentialIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
