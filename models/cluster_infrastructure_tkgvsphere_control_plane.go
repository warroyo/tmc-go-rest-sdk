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

// ClusterInfrastructureTkgvsphereControlPlane VSphere specific control plane configuration for workload cluster object.
//
// swagger:model cluster.infrastructure.tkgvsphere.ControlPlane
type ClusterInfrastructureTkgvsphereControlPlane struct {

	// High Availability or Non High Availability Cluster. HA cluster
	// creates three controlplane machines, and non HA creates just one.
	HighAvailability bool `json:"highAvailability,omitempty"`

	// VM specific configuration.
	VMConfig *CommonClusterTKGVsphereVMConfig `json:"vmConfig,omitempty"`
}

// Validate validates this cluster infrastructure tkgvsphere control plane
func (m *ClusterInfrastructureTkgvsphereControlPlane) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVMConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgvsphereControlPlane) validateVMConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VMConfig) { // not required
		return nil
	}

	if m.VMConfig != nil {
		if err := m.VMConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster infrastructure tkgvsphere control plane based on the context it is used
func (m *ClusterInfrastructureTkgvsphereControlPlane) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVMConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgvsphereControlPlane) contextValidateVMConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VMConfig != nil {
		if err := m.VMConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgvsphereControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgvsphereControlPlane) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgvsphereControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
