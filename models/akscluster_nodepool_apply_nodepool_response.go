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

// AksclusterNodepoolApplyNodepoolResponse Response from applying a Nodepool.
//
// swagger:model akscluster.nodepool.ApplyNodepoolResponse
type AksclusterNodepoolApplyNodepoolResponse struct {

	// Nodepool applied.
	Nodepool *AksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// Validate validates this akscluster nodepool apply nodepool response
func (m *AksclusterNodepoolApplyNodepoolResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodepool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterNodepoolApplyNodepoolResponse) validateNodepool(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodepool) { // not required
		return nil
	}

	if m.Nodepool != nil {
		if err := m.Nodepool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodepool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodepool")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this akscluster nodepool apply nodepool response based on the context it is used
func (m *AksclusterNodepoolApplyNodepoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodepool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterNodepoolApplyNodepoolResponse) contextValidateNodepool(ctx context.Context, formats strfmt.Registry) error {

	if m.Nodepool != nil {
		if err := m.Nodepool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodepool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodepool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterNodepoolApplyNodepoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterNodepoolApplyNodepoolResponse) UnmarshalBinary(b []byte) error {
	var res AksclusterNodepoolApplyNodepoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
