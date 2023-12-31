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

// ClusterInfrastructureTkgazureClusterNetwork Network information for the cluster.
//
// swagger:model cluster.infrastructure.tkgazure.ClusterNetwork
type ClusterInfrastructureTkgazureClusterNetwork struct {

	// APIServerPort specifies the port address for the cluster (optional).
	// The port value defaults to 6443.
	APIServerPort int32 `json:"apiServerPort,omitempty"`

	// Pod CIDR for Kubernetes pods. Defaults to 100.96.0.0/11.
	Pods *ClusterInfrastructureTkgazureNetworkRanges `json:"pods,omitempty"`

	// Service CIDR for Kubernetes services. Defaults to 100.64.0.0/13.
	Services *ClusterInfrastructureTkgazureNetworkRanges `json:"services,omitempty"`
}

// Validate validates this cluster infrastructure tkgazure cluster network
func (m *ClusterInfrastructureTkgazureClusterNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgazureClusterNetwork) validatePods(formats strfmt.Registry) error {
	if swag.IsZero(m.Pods) { // not required
		return nil
	}

	if m.Pods != nil {
		if err := m.Pods.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pods")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureClusterNetwork) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	if m.Services != nil {
		if err := m.Services.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("services")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster infrastructure tkgazure cluster network based on the context it is used
func (m *ClusterInfrastructureTkgazureClusterNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgazureClusterNetwork) contextValidatePods(ctx context.Context, formats strfmt.Registry) error {

	if m.Pods != nil {
		if err := m.Pods.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pods")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureClusterNetwork) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	if m.Services != nil {
		if err := m.Services.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("services")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureClusterNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureClusterNetwork) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgazureClusterNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
