// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInfrastructureTkgawsDistribution Distribution of the AWS cluster.
//
// swagger:model cluster.infrastructure.tkgaws.Distribution
type ClusterInfrastructureTkgawsDistribution struct {

	// Arch of the OS used for the cluster.
	OsArch string `json:"osArch,omitempty"`

	// Name of the OS used for the cluster.
	OsName string `json:"osName,omitempty"`

	// Version of the OS used for the cluster.
	OsVersion string `json:"osVersion,omitempty"`

	// Name of the account (provisioner credential) in which to create the cluster.
	ProvisionerCredentialName string `json:"provisionerCredentialName,omitempty"`

	// Region of the cluster.
	Region string `json:"region,omitempty"`

	// Version of the cluster.
	Version string `json:"version,omitempty"`
}

// Validate validates this cluster infrastructure tkgaws distribution
func (m *ClusterInfrastructureTkgawsDistribution) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster infrastructure tkgaws distribution based on context it is used
func (m *ClusterInfrastructureTkgawsDistribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgawsDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgawsDistribution) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgawsDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
