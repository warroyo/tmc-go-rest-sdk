// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupInsightFullName Full name of the cluster group insight. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model clustergroup.insight.FullName
type ClustergroupInsightFullName struct {

	// Name of the cluster group.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// UID of the insight.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this clustergroup insight full name
func (m *ClustergroupInsightFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup insight full name based on context it is used
func (m *ClustergroupInsightFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupInsightFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupInsightFullName) UnmarshalBinary(b []byte) error {
	var res ClustergroupInsightFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
