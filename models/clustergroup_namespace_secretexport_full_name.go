// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupNamespaceSecretexportFullName Full name of the Secret Export.
//
// swagger:model clustergroup.namespace.secretexport.FullName
type ClustergroupNamespaceSecretexportFullName struct {

	// Name of Cluster Group.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Name of the Secret Export (expected to share the same name of the secret).
	Name string `json:"name,omitempty"`

	// Name of Namespace.
	NamespaceName string `json:"namespaceName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this clustergroup namespace secretexport full name
func (m *ClustergroupNamespaceSecretexportFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup namespace secretexport full name based on context it is used
func (m *ClustergroupNamespaceSecretexportFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretexportFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretexportFullName) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceSecretexportFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
