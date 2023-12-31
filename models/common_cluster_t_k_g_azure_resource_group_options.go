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

// CommonClusterTKGAzureResourceGroupOptions Resource group options for TKG Azure.
//
// swagger:model common.cluster.TKGAzureResourceGroupOptions
type CommonClusterTKGAzureResourceGroupOptions struct {

	// Name of the resource group.
	Name string `json:"name,omitempty"`

	// List of VNets.
	VnetOptions []*CommonClusterTKGAzureVNetOptions `json:"vnetOptions"`
}

// Validate validates this common cluster t k g azure resource group options
func (m *CommonClusterTKGAzureResourceGroupOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVnetOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterTKGAzureResourceGroupOptions) validateVnetOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.VnetOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.VnetOptions); i++ {
		if swag.IsZero(m.VnetOptions[i]) { // not required
			continue
		}

		if m.VnetOptions[i] != nil {
			if err := m.VnetOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vnetOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vnetOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common cluster t k g azure resource group options based on the context it is used
func (m *CommonClusterTKGAzureResourceGroupOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVnetOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterTKGAzureResourceGroupOptions) contextValidateVnetOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VnetOptions); i++ {

		if m.VnetOptions[i] != nil {
			if err := m.VnetOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vnetOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vnetOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGAzureResourceGroupOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGAzureResourceGroupOptions) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGAzureResourceGroupOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
