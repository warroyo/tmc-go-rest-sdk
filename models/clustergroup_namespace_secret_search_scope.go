// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupNamespaceSecretSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model clustergroup.namespace.secret.SearchScope
type ClustergroupNamespaceSecretSearchScope struct {

	// Scope search to the specified cluster_group_name; supports globbing; default (*).
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified namespace_name; supports globbing; default (*).
	NamespaceName string `json:"namespaceName,omitempty"`

	// Filter by secret type; supports globbing; default (*).
	SecretType string `json:"secretType,omitempty"`
}

// Validate validates this clustergroup namespace secret search scope
func (m *ClustergroupNamespaceSecretSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup namespace secret search scope based on context it is used
func (m *ClustergroupNamespaceSecretSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretSearchScope) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceSecretSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
