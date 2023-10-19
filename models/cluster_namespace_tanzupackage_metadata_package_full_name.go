// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceTanzupackageMetadataPackageFullName Full name of the Package.
//
// swagger:model cluster.namespace.tanzupackage.metadata.package.FullName
type ClusterNamespaceTanzupackageMetadataPackageFullName struct {

	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of the Package metadata.
	MetadataName string `json:"metadataName,omitempty"`

	// Name of the Package. It represents version of the Package metadata.
	Name string `json:"name,omitempty"`

	// Name of Namespace.
	NamespaceName string `json:"namespaceName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this cluster namespace tanzupackage metadata package full name
func (m *ClusterNamespaceTanzupackageMetadataPackageFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster namespace tanzupackage metadata package full name based on context it is used
func (m *ClusterNamespaceTanzupackageMetadataPackageFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageMetadataPackageFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageMetadataPackageFullName) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageMetadataPackageFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
