// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterTKGAzureSubnetOptions Azure specific subnet options.
//
// swagger:model common.cluster.TKGAzureSubnetOptions
type CommonClusterTKGAzureSubnetOptions struct {

	// CIDR of the Azure subnet.
	CidrBlocks []string `json:"cidrBlocks"`

	// Name of the subnet.
	Name string `json:"name,omitempty"`
}

// Validate validates this common cluster t k g azure subnet options
func (m *CommonClusterTKGAzureSubnetOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster t k g azure subnet options based on context it is used
func (m *CommonClusterTKGAzureSubnetOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGAzureSubnetOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGAzureSubnetOptions) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGAzureSubnetOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
