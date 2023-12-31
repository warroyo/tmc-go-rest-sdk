// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupPolicySearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model clustergroup.policy.SearchScope
type ClustergroupPolicySearchScope struct {

	// Scope search to the specified cluster_group_name; supports globbing; default (*).
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`
}

// Validate validates this clustergroup policy search scope
func (m *ClustergroupPolicySearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup policy search scope based on context it is used
func (m *ClustergroupPolicySearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupPolicySearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupPolicySearchScope) UnmarshalBinary(b []byte) error {
	var res ClustergroupPolicySearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
