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

// EksclusterAmiTypeOptions AWS specific ami types options.
//
// swagger:model ekscluster.AmiTypeOptions
type EksclusterAmiTypeOptions struct {

	// Supported AMI type.
	AmiType string `json:"amiType,omitempty"`

	// List of instance types.
	InstanceTypeOptions []*EksclusterInstanceTypeOptions `json:"instanceTypeOptions"`
}

// Validate validates this ekscluster ami type options
func (m *EksclusterAmiTypeOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceTypeOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EksclusterAmiTypeOptions) validateInstanceTypeOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceTypeOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceTypeOptions); i++ {
		if swag.IsZero(m.InstanceTypeOptions[i]) { // not required
			continue
		}

		if m.InstanceTypeOptions[i] != nil {
			if err := m.InstanceTypeOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceTypeOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceTypeOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ekscluster ami type options based on the context it is used
func (m *EksclusterAmiTypeOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceTypeOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EksclusterAmiTypeOptions) contextValidateInstanceTypeOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceTypeOptions); i++ {

		if m.InstanceTypeOptions[i] != nil {
			if err := m.InstanceTypeOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceTypeOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceTypeOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterAmiTypeOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterAmiTypeOptions) UnmarshalBinary(b []byte) error {
	var res EksclusterAmiTypeOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
