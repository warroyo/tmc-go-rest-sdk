// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EksclusterNodepoolDeleteNodepoolResponse Response from deleting a Nodepool.
//
// swagger:model ekscluster.nodepool.DeleteNodepoolResponse
type EksclusterNodepoolDeleteNodepoolResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this ekscluster nodepool delete nodepool response
func (m *EksclusterNodepoolDeleteNodepoolResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ekscluster nodepool delete nodepool response based on context it is used
func (m *EksclusterNodepoolDeleteNodepoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterNodepoolDeleteNodepoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterNodepoolDeleteNodepoolResponse) UnmarshalBinary(b []byte) error {
	var res EksclusterNodepoolDeleteNodepoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}