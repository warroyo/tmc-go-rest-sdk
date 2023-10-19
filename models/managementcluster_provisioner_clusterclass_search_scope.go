// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementclusterProvisionerClusterclassSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model managementcluster.provisioner.clusterclass.SearchScope
type ManagementclusterProvisionerClusterclassSearchScope struct {

	// Scope search to the specified management_cluster_name; supports globbing; default (*).
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified provisioner_name; supports globbing; default (*).
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this managementcluster provisioner clusterclass search scope
func (m *ManagementclusterProvisionerClusterclassSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this managementcluster provisioner clusterclass search scope based on context it is used
func (m *ManagementclusterProvisionerClusterclassSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerClusterclassSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerClusterclassSearchScope) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerClusterclassSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
