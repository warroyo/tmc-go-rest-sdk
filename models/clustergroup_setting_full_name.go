// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupSettingFullName Full name of the ClusterGroup on which the Setting is applied.
//
// swagger:model clustergroup.setting.FullName
type ClustergroupSettingFullName struct {

	// Name of the cluster group.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Name of the setting.
	// This should be same as the name of setting type.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this clustergroup setting full name
func (m *ClustergroupSettingFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup setting full name based on context it is used
func (m *ClustergroupSettingFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupSettingFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupSettingFullName) UnmarshalBinary(b []byte) error {
	var res ClustergroupSettingFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
