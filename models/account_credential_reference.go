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

// AccountCredentialReference Reference contains information on another TMC resource using this credential.
// Credentials with references cannot be deleted as they are being used by some resource.
//
// swagger:model account.credential.Reference
type AccountCredentialReference struct {

	// UID of the reference.
	ReferenceUID *VmwareTanzuCoreV1alpha1ObjectUID `json:"referenceUid,omitempty"`

	// The RID of the resource referencing the credential.
	ResourceID *VmwareTanzuCoreV1alpha1ObjectReference `json:"resourceId,omitempty"`
}

// Validate validates this account credential reference
func (m *AccountCredentialReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferenceUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialReference) validateReferenceUID(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceUID) { // not required
		return nil
	}

	if m.ReferenceUID != nil {
		if err := m.ReferenceUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referenceUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referenceUid")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialReference) validateResourceID(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceID) { // not required
		return nil
	}

	if m.ResourceID != nil {
		if err := m.ResourceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account credential reference based on the context it is used
func (m *AccountCredentialReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReferenceUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialReference) contextValidateReferenceUID(ctx context.Context, formats strfmt.Registry) error {

	if m.ReferenceUID != nil {
		if err := m.ReferenceUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referenceUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referenceUid")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialReference) contextValidateResourceID(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceID != nil {
		if err := m.ResourceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialReference) UnmarshalBinary(b []byte) error {
	var res AccountCredentialReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
