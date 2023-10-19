// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupDeleteClusterGroupResponse Response from deleting a ClusterGroup.
//
// swagger:model clustergroup.DeleteClusterGroupResponse
type ClustergroupDeleteClusterGroupResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this clustergroup delete cluster group response
func (m *ClustergroupDeleteClusterGroupResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup delete cluster group response based on context it is used
func (m *ClustergroupDeleteClusterGroupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupDeleteClusterGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupDeleteClusterGroupResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupDeleteClusterGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}