// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EksclusterSecurityGroupOptions Security group options.
//
// swagger:model ekscluster.SecurityGroupOptions
type EksclusterSecurityGroupOptions struct {

	// Name of the security group.
	Name string `json:"name,omitempty"`

	// ID of the security group.
	SecurityGroupID string `json:"securityGroupId,omitempty"`
}

// Validate validates this ekscluster security group options
func (m *EksclusterSecurityGroupOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ekscluster security group options based on context it is used
func (m *EksclusterSecurityGroupOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterSecurityGroupOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterSecurityGroupOptions) UnmarshalBinary(b []byte) error {
	var res EksclusterSecurityGroupOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}