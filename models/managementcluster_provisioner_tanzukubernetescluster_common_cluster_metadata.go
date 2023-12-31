// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata The labels and annotations configurations.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.common.cluster.Metadata
type ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata struct {

	// The annotations configuration.
	Annotations map[string]string `json:"annotations,omitempty"`

	// The labels configuration.
	Labels map[string]string `json:"labels,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster common cluster metadata
func (m *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this managementcluster provisioner tanzukubernetescluster common cluster metadata based on context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
