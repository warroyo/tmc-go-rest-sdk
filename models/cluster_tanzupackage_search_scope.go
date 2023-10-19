// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterTanzupackageSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model cluster.tanzupackage.SearchScope
type ClusterTanzupackageSearchScope struct {

	// Scope search to the specified cluster_name; supports globbing; default (*).
	ClusterName string `json:"clusterName,omitempty"`

	// Scope search to the specified management_cluster_name; supports globbing; default (*).
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Scope search to the specified provisioner_name; supports globbing; default (*).
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this cluster tanzupackage search scope
func (m *ClusterTanzupackageSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster tanzupackage search scope based on context it is used
func (m *ClusterTanzupackageSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTanzupackageSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTanzupackageSearchScope) UnmarshalBinary(b []byte) error {
	var res ClusterTanzupackageSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
