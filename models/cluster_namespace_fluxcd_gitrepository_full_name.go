// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceFluxcdGitrepositoryFullName Full name of the Repository.
//
// swagger:model cluster.namespace.fluxcd.gitrepository.FullName
type ClusterNamespaceFluxcdGitrepositoryFullName struct {

	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of the Repository.
	Name string `json:"name,omitempty"`

	// Name of Namespace.
	NamespaceName string `json:"namespaceName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this cluster namespace fluxcd gitrepository full name
func (m *ClusterNamespaceFluxcdGitrepositoryFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster namespace fluxcd gitrepository full name based on context it is used
func (m *ClusterNamespaceFluxcdGitrepositoryFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdGitrepositoryFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdGitrepositoryFullName) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceFluxcdGitrepositoryFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}