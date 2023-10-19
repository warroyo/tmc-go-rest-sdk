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

// AksclusterSpec Spec of the cluster.
//
// swagger:model akscluster.Spec
type AksclusterSpec struct {

	// Name of the cluster in TMC.
	AgentName string `json:"agentName,omitempty"`

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// AKS config for the cluster.
	Config *AksclusterClusterConfig `json:"config,omitempty"`

	// The name of the proxy credential to be used for the TMC agents.
	// If specified, the TMC agents will be installed with the proxy configurations specified in the proxy credential.
	ProxyName string `json:"proxyName,omitempty"`

	// Resource ID of the cluster in Azure.
	ResourceID string `json:"resourceId,omitempty"`
}

// Validate validates this akscluster spec
func (m *AksclusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterSpec) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this akscluster spec based on the context it is used
func (m *AksclusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterSpec) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterSpec) UnmarshalBinary(b []byte) error {
	var res AksclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}