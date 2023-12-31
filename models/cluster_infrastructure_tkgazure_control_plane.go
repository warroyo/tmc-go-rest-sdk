// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInfrastructureTkgazureControlPlane Control plane configuration for the Azure cluster.
//
// swagger:model cluster.infrastructure.tkgazure.ControlPlane
type ClusterInfrastructureTkgazureControlPlane struct {

	// List of Availability Zone for control plane nodes.
	AvailabilityZones []string `json:"availabilityZones"`

	// Flag which controls if the cluster needs to be highly available. A highly available cluster has three
	// controlplane machines, and a non highly available cluster has one.
	HighAvailability bool `json:"highAvailability,omitempty"`

	// Control plane vm size.
	VMSize string `json:"vmSize,omitempty"`
}

// Validate validates this cluster infrastructure tkgazure control plane
func (m *ClusterInfrastructureTkgazureControlPlane) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster infrastructure tkgazure control plane based on context it is used
func (m *ClusterInfrastructureTkgazureControlPlane) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureControlPlane) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgazureControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
