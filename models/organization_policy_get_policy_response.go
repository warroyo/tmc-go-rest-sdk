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

// OrganizationPolicyGetPolicyResponse Response from getting a Policy.
//
// swagger:model organization.policy.GetPolicyResponse
type OrganizationPolicyGetPolicyResponse struct {

	// Policy returned.
	Policy *OrganizationPolicyPolicy `json:"policy,omitempty"`
}

// Validate validates this organization policy get policy response
func (m *OrganizationPolicyGetPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationPolicyGetPolicyResponse) validatePolicy(formats strfmt.Registry) error {
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

// ContextValidate validate this organization policy get policy response based on the context it is used
func (m *OrganizationPolicyGetPolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationPolicyGetPolicyResponse) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

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
func (m *OrganizationPolicyGetPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationPolicyGetPolicyResponse) UnmarshalBinary(b []byte) error {
	var res OrganizationPolicyGetPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
