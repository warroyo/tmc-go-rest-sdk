// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInfrastructureTkgservicevsphereStorageSettings Storage related settings for workload cluster.
//
// swagger:model cluster.infrastructure.tkgservicevsphere.StorageSettings
type ClusterInfrastructureTkgservicevsphereStorageSettings struct {

	// Classes is a list of storage classes from the supervisor namespace to expose within a cluster.
	// If omitted, all storage classes from the supervisor namespace will be exposed within the cluster.
	Classes []string `json:"classes"`

	// DefaultClass is the valid storage class name which is treated as the default storage class within a cluster.
	// If omitted, no default storage class is set.
	DefaultClass string `json:"defaultClass,omitempty"`
}

// Validate validates this cluster infrastructure tkgservicevsphere storage settings
func (m *ClusterInfrastructureTkgservicevsphereStorageSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster infrastructure tkgservicevsphere storage settings based on context it is used
func (m *ClusterInfrastructureTkgservicevsphereStorageSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgservicevsphereStorageSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgservicevsphereStorageSettings) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgservicevsphereStorageSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
