// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInfrastructureTkgawsVPC VPC configuration for the cluster.
//
// swagger:model cluster.infrastructure.tkgaws.VPC
type ClusterInfrastructureTkgawsVPC struct {

	// CIDR for AWS VPC.
	// A valid example is 10.0.0.0/16.
	// For the allowed ranges, please refer to AWS documentation..
	CidrBlock string `json:"cidrBlock,omitempty"`

	// AWS VPC ID. The rest of the fields are ignored if this field is specified..
	ID string `json:"id,omitempty"`
}

// Validate validates this cluster infrastructure tkgaws v p c
func (m *ClusterInfrastructureTkgawsVPC) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster infrastructure tkgaws v p c based on context it is used
func (m *ClusterInfrastructureTkgawsVPC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgawsVPC) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgawsVPC) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgawsVPC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}