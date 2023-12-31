// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterDataprotectionSnapshotlocationFullName Full name of the namespace. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model cluster.dataprotection.snapshotlocation.FullName
type ClusterDataprotectionSnapshotlocationFullName struct {

	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this Snapshot Location.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this cluster dataprotection snapshotlocation full name
func (m *ClusterDataprotectionSnapshotlocationFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster dataprotection snapshotlocation full name based on context it is used
func (m *ClusterDataprotectionSnapshotlocationFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionSnapshotlocationFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionSnapshotlocationFullName) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionSnapshotlocationFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
