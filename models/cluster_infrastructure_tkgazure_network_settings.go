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

// ClusterInfrastructureTkgazureNetworkSettings Cluster network and provider network information for the cluster.
//
// swagger:model cluster.infrastructure.tkgazure.NetworkSettings
type ClusterInfrastructureTkgazureNetworkSettings struct {

	// Kubernetes network information for the cluster.
	Cluster *ClusterInfrastructureTkgazureClusterNetwork `json:"cluster,omitempty"`

	// PrivateCluster specifies Private cluster configuration. If empty, TMC creates a public cluster.
	PrivateCluster *ClusterInfrastructureTkgazurePrivateCluster `json:"privateCluster,omitempty"`

	// Provider specific network information for the cluster.
	Provider *ClusterInfrastructureTkgazureProviderNetwork `json:"provider,omitempty"`
}

// Validate validates this cluster infrastructure tkgazure network settings
func (m *ClusterInfrastructureTkgazureNetworkSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgazureNetworkSettings) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureNetworkSettings) validatePrivateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateCluster) { // not required
		return nil
	}

	if m.PrivateCluster != nil {
		if err := m.PrivateCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privateCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privateCluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureNetworkSettings) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster infrastructure tkgazure network settings based on the context it is used
func (m *ClusterInfrastructureTkgazureNetworkSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgazureNetworkSettings) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureNetworkSettings) contextValidatePrivateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivateCluster != nil {
		if err := m.PrivateCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privateCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privateCluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInfrastructureTkgazureNetworkSettings) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.Provider != nil {
		if err := m.Provider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureNetworkSettings) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgazureNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
