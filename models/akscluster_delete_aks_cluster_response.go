// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AksclusterDeleteAksClusterResponse Response from deleting an AksCluster.
//
// swagger:model akscluster.DeleteAksClusterResponse
type AksclusterDeleteAksClusterResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this akscluster delete aks cluster response
func (m *AksclusterDeleteAksClusterResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this akscluster delete aks cluster response based on context it is used
func (m *AksclusterDeleteAksClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterDeleteAksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterDeleteAksClusterResponse) UnmarshalBinary(b []byte) error {
	var res AksclusterDeleteAksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
