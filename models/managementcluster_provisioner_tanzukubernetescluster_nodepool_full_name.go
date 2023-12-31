// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName Full name of the nodepool. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.nodepool.FullName
type ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName struct {

	// Name of the management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this nodepool.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Provisioner of the cluster.
	ProvisionerName string `json:"provisionerName,omitempty"`

	// Name of the cluster.
	TanzuKubernetesClusterName string `json:"tanzuKubernetesClusterName,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster nodepool full name
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this managementcluster provisioner tanzukubernetescluster nodepool full name based on context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
