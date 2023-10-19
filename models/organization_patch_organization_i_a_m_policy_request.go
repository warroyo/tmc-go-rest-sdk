// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationPatchOrganizationIAMPolicyRequest PatchOrganizationIAMPolicy request message.
//
// swagger:model organization.PatchOrganizationIAMPolicyRequest
type OrganizationPatchOrganizationIAMPolicyRequest struct {

	// Policy delta to be applied.
	BindingDeltaList []*VmwareTanzuCoreV1alpha1PolicyBindingDelta `json:"bindingDeltaList"`

	// Organization fullname.
	FullName *OrganizationFullName `json:"fullName,omitempty"`
}

// Validate validates this organization patch organization i a m policy request
func (m *OrganizationPatchOrganizationIAMPolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindingDeltaList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationPatchOrganizationIAMPolicyRequest) validateBindingDeltaList(formats strfmt.Registry) error {
	if swag.IsZero(m.BindingDeltaList) { // not required
		return nil
	}

	for i := 0; i < len(m.BindingDeltaList); i++ {
		if swag.IsZero(m.BindingDeltaList[i]) { // not required
			continue
		}

		if m.BindingDeltaList[i] != nil {
			if err := m.BindingDeltaList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bindingDeltaList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bindingDeltaList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganizationPatchOrganizationIAMPolicyRequest) validateFullName(formats strfmt.Registry) error {
	if swag.IsZero(m.FullName) { // not required
		return nil
	}

	if m.FullName != nil {
		if err := m.FullName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this organization patch organization i a m policy request based on the context it is used
func (m *OrganizationPatchOrganizationIAMPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBindingDeltaList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationPatchOrganizationIAMPolicyRequest) contextValidateBindingDeltaList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BindingDeltaList); i++ {

		if m.BindingDeltaList[i] != nil {
			if err := m.BindingDeltaList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bindingDeltaList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bindingDeltaList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganizationPatchOrganizationIAMPolicyRequest) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

	if m.FullName != nil {
		if err := m.FullName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationPatchOrganizationIAMPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationPatchOrganizationIAMPolicyRequest) UnmarshalBinary(b []byte) error {
	var res OrganizationPatchOrganizationIAMPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
