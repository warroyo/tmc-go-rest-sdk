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

// VmwareTanzuCoreV1alpha1PolicySubject Representation of a subject in resource manager.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.Subject
type VmwareTanzuCoreV1alpha1PolicySubject struct {

	// Subject type.
	Kind *VmwareTanzuCoreV1alpha1PolicySubjectKind `json:"kind,omitempty"`

	// Subject name - allow max characters for email - 320.
	Name string `json:"name,omitempty"`
}

// Validate validates this vmware tanzu core v1alpha1 policy subject
func (m *VmwareTanzuCoreV1alpha1PolicySubject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareTanzuCoreV1alpha1PolicySubject) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vmware tanzu core v1alpha1 policy subject based on the context it is used
func (m *VmwareTanzuCoreV1alpha1PolicySubject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareTanzuCoreV1alpha1PolicySubject) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {
		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareTanzuCoreV1alpha1PolicySubject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuCoreV1alpha1PolicySubject) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1PolicySubject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}