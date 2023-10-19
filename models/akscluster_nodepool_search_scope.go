// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AksclusterNodepoolSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model akscluster.nodepool.SearchScope
type AksclusterNodepoolSearchScope struct {

	// Scope search to the specified aks_cluster_name; supports globbing; default (*).
	AksClusterName string `json:"aksClusterName,omitempty"`

	// Scope search to the specified credential_name; supports globbing; default (*).
	CredentialName string `json:"credentialName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified resource_group_name; supports globbing; default (*).
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Scope search to the specified subscription_id; supports globbing; default (*).
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// Validate validates this akscluster nodepool search scope
func (m *AksclusterNodepoolSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this akscluster nodepool search scope based on context it is used
func (m *AksclusterNodepoolSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterNodepoolSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterNodepoolSearchScope) UnmarshalBinary(b []byte) error {
	var res AksclusterNodepoolSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}