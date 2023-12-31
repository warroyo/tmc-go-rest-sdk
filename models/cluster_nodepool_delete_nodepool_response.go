// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNodepoolDeleteNodepoolResponse Response from deleting a Nodepool.
//
// swagger:model cluster.nodepool.DeleteNodepoolResponse
type ClusterNodepoolDeleteNodepoolResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this cluster nodepool delete nodepool response
func (m *ClusterNodepoolDeleteNodepoolResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster nodepool delete nodepool response based on context it is used
func (m *ClusterNodepoolDeleteNodepoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNodepoolDeleteNodepoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNodepoolDeleteNodepoolResponse) UnmarshalBinary(b []byte) error {
	var res ClusterNodepoolDeleteNodepoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
