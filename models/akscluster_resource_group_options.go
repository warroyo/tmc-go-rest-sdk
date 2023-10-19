// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AksclusterResourceGroupOptions Azure AKS specific resource group options.
//
// swagger:model akscluster.ResourceGroupOptions
type AksclusterResourceGroupOptions struct {

	// Name of the resource group.
	Name string `json:"name,omitempty"`
}

// Validate validates this akscluster resource group options
func (m *AksclusterResourceGroupOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this akscluster resource group options based on context it is used
func (m *AksclusterResourceGroupOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterResourceGroupOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterResourceGroupOptions) UnmarshalBinary(b []byte) error {
	var res AksclusterResourceGroupOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
