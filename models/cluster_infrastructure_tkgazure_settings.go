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

// ClusterInfrastructureTkgazureSettings Network and security settings for the cluster.
//
// swagger:model cluster.infrastructure.tkgazure.Settings
type ClusterInfrastructureTkgazureSettings struct {

	// NetworkSettings specifies network-related settings for the cluster.
	Network *ClusterInfrastructureTkgazureNetworkSettings `json:"network,omitempty"`

	// SecuritySettings specifies security-related settings for the cluster.
	Security *ClusterInfrastructureTkgazureSecuritySettings `json:"security,omitempty"`
}

// Validate validates this cluster infrastructure tkgazure settings
func (m *ClusterInfrastructureTkgazureSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgazureSettings) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureSettings) validateSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.Security) { // not required
		return nil
	}

	if m.Security != nil {
		if err := m.Security.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster infrastructure tkgazure settings based on the context it is used
func (m *ClusterInfrastructureTkgazureSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgazureSettings) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureSettings) contextValidateSecurity(ctx context.Context, formats strfmt.Registry) error {

	if m.Security != nil {
		if err := m.Security.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureSettings) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgazureSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
