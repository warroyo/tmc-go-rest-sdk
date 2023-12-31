// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupInsightSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model clustergroup.insight.SearchScope
type ClustergroupInsightSearchScope struct {

	// Scope search to the specified cluster_group_name; supports globbing; default (*).
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Filter by cluster name; supports globbing; default (*).
	ClusterName string `json:"clusterName,omitempty"`

	// Filter by insight type; supports globbing; default (*).
	InsightType string `json:"insightType,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Filter by the kind of source resource (eg. integration); supports globbing; default (*).
	SourceKind string `json:"sourceKind,omitempty"`

	// Filter by the name of source resource (eg. tanzu-observability-saas); supports globbing; default (*).
	SourceName string `json:"sourceName,omitempty"`

	// Filter by the uid of source resource; supports globbing; default (*).
	SourceUID string `json:"sourceUid,omitempty"`
}

// Validate validates this clustergroup insight search scope
func (m *ClustergroupInsightSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup insight search scope based on context it is used
func (m *ClustergroupInsightSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupInsightSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupInsightSearchScope) UnmarshalBinary(b []byte) error {
	var res ClustergroupInsightSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
