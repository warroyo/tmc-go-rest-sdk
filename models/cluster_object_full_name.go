// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterObjectFullName Full name of the object. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model cluster.object.FullName
type ClusterObjectFullName struct {

	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of ManagementCluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of Object.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this cluster object full name
func (m *ClusterObjectFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster object full name based on context it is used
func (m *ClusterObjectFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterObjectFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterObjectFullName) UnmarshalBinary(b []byte) error {
	var res ClusterObjectFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
