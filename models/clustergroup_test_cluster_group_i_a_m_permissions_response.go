// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupTestClusterGroupIAMPermissionsResponse TestClusterGroupIAMPermissions response message.
//
// swagger:model clustergroup.TestClusterGroupIAMPermissionsResponse
type ClustergroupTestClusterGroupIAMPermissionsResponse struct {

	// List of allowed permissions.
	Permissions []string `json:"permissions"`

	// Unique ID of the resource. Coupled with FullName (RID), provides a stable unique
	// identifier for the resource.
	UID string `json:"uid,omitempty"`
}

// Validate validates this clustergroup test cluster group i a m permissions response
func (m *ClustergroupTestClusterGroupIAMPermissionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup test cluster group i a m permissions response based on context it is used
func (m *ClustergroupTestClusterGroupIAMPermissionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupTestClusterGroupIAMPermissionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupTestClusterGroupIAMPermissionsResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupTestClusterGroupIAMPermissionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
