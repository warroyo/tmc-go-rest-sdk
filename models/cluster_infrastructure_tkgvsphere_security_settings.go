// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInfrastructureTkgvsphereSecuritySettings Security related settings for the cluster.
//
// swagger:model cluster.infrastructure.tkgvsphere.SecuritySettings
type ClusterInfrastructureTkgvsphereSecuritySettings struct {

	// SSH key for provisioning and accessing the cluster VMs.
	SSHKey string `json:"sshKey,omitempty"`
}

// Validate validates this cluster infrastructure tkgvsphere security settings
func (m *ClusterInfrastructureTkgvsphereSecuritySettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster infrastructure tkgvsphere security settings based on context it is used
func (m *ClusterInfrastructureTkgvsphereSecuritySettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgvsphereSecuritySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgvsphereSecuritySettings) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgvsphereSecuritySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}