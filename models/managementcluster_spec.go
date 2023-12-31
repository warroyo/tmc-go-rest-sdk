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

// ManagementclusterSpec The ManagementCluster spec.
//
// swagger:model managementcluster.Spec
type ManagementclusterSpec struct {

	// Default cluster group for workload clusters.
	DefaultClusterGroup string `json:"defaultClusterGroup,omitempty"`

	// Optional default workload cluster image registry is the
	// name of the Image Registry Config to be used for workload clusters.
	// If set empty, no image registry config will be used.
	// If set non-empty, defined image registry config will be used.
	// If left unset, management cluster's image registry config will be used.
	DefaultWorkloadClusterImageRegistry string `json:"defaultWorkloadClusterImageRegistry,omitempty"`

	// Optional default workload clusters proxy name is the
	// Proxy Config to be used for workload clusters.
	DefaultWorkloadClusterProxyName string `json:"defaultWorkloadClusterProxyName,omitempty"`

	// Optional image registry is the name of the Image Registry Config
	// to be used for the management cluster.
	ImageRegistry string `json:"imageRegistry,omitempty"`

	// Kubernetes Provider Type of user's choice for registration.
	KubernetesProviderType *CommonClusterKubernetesProviderType `json:"kubernetesProviderType,omitempty"`

	// Optional proxy name is the name of the Proxy Config
	// to be used for the management cluster.
	ProxyName string `json:"proxyName,omitempty"`
}

// Validate validates this managementcluster spec
func (m *ManagementclusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKubernetesProviderType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterSpec) validateKubernetesProviderType(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesProviderType) { // not required
		return nil
	}

	if m.KubernetesProviderType != nil {
		if err := m.KubernetesProviderType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesProviderType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesProviderType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster spec based on the context it is used
func (m *ManagementclusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKubernetesProviderType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterSpec) contextValidateKubernetesProviderType(ctx context.Context, formats strfmt.Registry) error {

	if m.KubernetesProviderType != nil {
		if err := m.KubernetesProviderType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesProviderType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesProviderType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterSpec) UnmarshalBinary(b []byte) error {
	var res ManagementclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
