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

// ClusterInfrastructureTkgvsphereDistribution VSphere specific distribution.
//
// swagger:model cluster.infrastructure.tkgvsphere.Distribution
type ClusterInfrastructureTkgvsphereDistribution struct {

	// Arch of the OS used for the cluster.
	OsArch string `json:"osArch,omitempty"`

	// Name of the OS used for the cluster.
	OsName string `json:"osName,omitempty"`

	// Version of the OS used for the cluster.
	OsVersion string `json:"osVersion,omitempty"`

	// Version specifies the version of the Kubernetes cluster.
	Version string `json:"version,omitempty"`

	// Workspace defines a workspace configuration for the vSphere cloud provider.
	Workspace *CommonClusterTKGVsphereWorkspace `json:"workspace,omitempty"`
}

// Validate validates this cluster infrastructure tkgvsphere distribution
func (m *ClusterInfrastructureTkgvsphereDistribution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgvsphereDistribution) validateWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster infrastructure tkgvsphere distribution based on the context it is used
func (m *ClusterInfrastructureTkgvsphereDistribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInfrastructureTkgvsphereDistribution) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspace != nil {
		if err := m.Workspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgvsphereDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgvsphereDistribution) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgvsphereDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
