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

// AksclusterNodepoolTaint The node this Taint is attached to has the "effect" on
// any pod that does not tolerate the Taint.
//
// swagger:model akscluster.nodepool.Taint
type AksclusterNodepoolTaint struct {

	// Current effect state of the nodepool.
	Effect *AksclusterNodepoolTaintEffect `json:"effect,omitempty"`

	// The taint key to be applied to a node.
	Key string `json:"key,omitempty"`

	// The taint value corresponding to the taint key.
	Value string `json:"value,omitempty"`
}

// Validate validates this akscluster nodepool taint
func (m *AksclusterNodepoolTaint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterNodepoolTaint) validateEffect(formats strfmt.Registry) error {
	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	if m.Effect != nil {
		if err := m.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this akscluster nodepool taint based on the context it is used
func (m *AksclusterNodepoolTaint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEffect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterNodepoolTaint) contextValidateEffect(ctx context.Context, formats strfmt.Registry) error {

	if m.Effect != nil {
		if err := m.Effect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterNodepoolTaint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterNodepoolTaint) UnmarshalBinary(b []byte) error {
	var res AksclusterNodepoolTaint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}