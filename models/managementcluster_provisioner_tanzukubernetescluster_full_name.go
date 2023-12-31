// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementclusterProvisionerTanzukubernetesclusterFullName Full name of the cluster. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.FullName
type ManagementclusterProvisionerTanzukubernetesclusterFullName struct {

	// Name of the management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this cluster.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Provisioner of the cluster.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster full name
func (m *ManagementclusterProvisionerTanzukubernetesclusterFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this managementcluster provisioner tanzukubernetescluster full name based on context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterFullName) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
