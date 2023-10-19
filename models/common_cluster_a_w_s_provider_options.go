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

// CommonClusterAWSProviderOptions AWS cloud provider specific options.
//
// swagger:model common.cluster.AWSProviderOptions
type CommonClusterAWSProviderOptions struct {

	// The information about of our access to the AWS account.
	AccessStatus *CommonClusterAWSAccountAccessStatus `json:"accessStatus,omitempty"`

	// The advanced configuration options for AWS cloud provider.
	AdvancedConfigOptions []*CommonClusterAdvancedConfigOption `json:"advancedConfigOptions"`

	// Default VPC CIDR for AWS provider.
	DefaultVpcCidr string `json:"defaultVpcCidr,omitempty"`

	// List of regional options for the AWS cloud provider.
	RegionalOptions []*CommonClusterAWSRegionalOptions `json:"regionalOptions"`
}

// Validate validates this common cluster a w s provider options
func (m *CommonClusterAWSProviderOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvancedConfigOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionalOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterAWSProviderOptions) validateAccessStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessStatus) { // not required
		return nil
	}

	if m.AccessStatus != nil {
		if err := m.AccessStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessStatus")
			}
			return err
		}
	}

	return nil
}

func (m *CommonClusterAWSProviderOptions) validateAdvancedConfigOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedConfigOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.AdvancedConfigOptions); i++ {
		if swag.IsZero(m.AdvancedConfigOptions[i]) { // not required
			continue
		}

		if m.AdvancedConfigOptions[i] != nil {
			if err := m.AdvancedConfigOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterAWSProviderOptions) validateRegionalOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.RegionalOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.RegionalOptions); i++ {
		if swag.IsZero(m.RegionalOptions[i]) { // not required
			continue
		}

		if m.RegionalOptions[i] != nil {
			if err := m.RegionalOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regionalOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regionalOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common cluster a w s provider options based on the context it is used
func (m *CommonClusterAWSProviderOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdvancedConfigOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegionalOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterAWSProviderOptions) contextValidateAccessStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessStatus != nil {
		if err := m.AccessStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessStatus")
			}
			return err
		}
	}

	return nil
}

func (m *CommonClusterAWSProviderOptions) contextValidateAdvancedConfigOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdvancedConfigOptions); i++ {

		if m.AdvancedConfigOptions[i] != nil {
			if err := m.AdvancedConfigOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterAWSProviderOptions) contextValidateRegionalOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RegionalOptions); i++ {

		if m.RegionalOptions[i] != nil {
			if err := m.RegionalOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regionalOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regionalOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterAWSProviderOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterAWSProviderOptions) UnmarshalBinary(b []byte) error {
	var res CommonClusterAWSProviderOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}