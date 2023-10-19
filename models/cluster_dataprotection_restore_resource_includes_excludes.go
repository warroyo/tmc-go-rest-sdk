// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterDataprotectionRestoreResourceIncludesExcludes A combination of included and excluded resources.
//
// swagger:model cluster.dataprotection.restore.ResourceIncludesExcludes
type ClusterDataprotectionRestoreResourceIncludesExcludes struct {

	// A list of resources to be excluded.
	ExcludedResources []string `json:"excludedResources"`

	// A list of resources to be included.
	IncludedResources []string `json:"includedResources"`
}

// Validate validates this cluster dataprotection restore resource includes excludes
func (m *ClusterDataprotectionRestoreResourceIncludesExcludes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster dataprotection restore resource includes excludes based on context it is used
func (m *ClusterDataprotectionRestoreResourceIncludesExcludes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreResourceIncludesExcludes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionRestoreResourceIncludesExcludes) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionRestoreResourceIncludesExcludes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
