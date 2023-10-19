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

// EksclusterVPCOptions AWS specific vpc options.
//
// swagger:model ekscluster.VPCOptions
type EksclusterVPCOptions struct {

	// CIDR of the VPC.
	Cidr string `json:"cidr,omitempty"`

	// Name of the VPC.
	Name string `json:"name,omitempty"`

	// List of security group options.
	SecurityGroupOptions []*EksclusterSecurityGroupOptions `json:"securityGroupOptions"`

	// List of subnets that is supported in the VPC.
	SubnetOptions []*EksclusterSubnetOptions `json:"subnetOptions"`

	// ID of the AWS VPC.
	VpcID string `json:"vpcId,omitempty"`
}

// Validate validates this ekscluster v p c options
func (m *EksclusterVPCOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityGroupOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EksclusterVPCOptions) validateSecurityGroupOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroupOptions); i++ {
		if swag.IsZero(m.SecurityGroupOptions[i]) { // not required
			continue
		}

		if m.SecurityGroupOptions[i] != nil {
			if err := m.SecurityGroupOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroupOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("securityGroupOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EksclusterVPCOptions) validateSubnetOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.SubnetOptions); i++ {
		if swag.IsZero(m.SubnetOptions[i]) { // not required
			continue
		}

		if m.SubnetOptions[i] != nil {
			if err := m.SubnetOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ekscluster v p c options based on the context it is used
func (m *EksclusterVPCOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecurityGroupOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EksclusterVPCOptions) contextValidateSecurityGroupOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecurityGroupOptions); i++ {

		if m.SecurityGroupOptions[i] != nil {
			if err := m.SecurityGroupOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroupOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("securityGroupOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EksclusterVPCOptions) contextValidateSubnetOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubnetOptions); i++ {

		if m.SubnetOptions[i] != nil {
			if err := m.SubnetOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterVPCOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterVPCOptions) UnmarshalBinary(b []byte) error {
	var res EksclusterVPCOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}