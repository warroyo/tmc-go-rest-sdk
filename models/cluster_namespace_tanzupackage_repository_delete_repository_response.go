// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse Response from deleting a Repository.
//
// swagger:model cluster.namespace.tanzupackage.repository.DeleteRepositoryResponse
type ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this cluster namespace tanzupackage repository delete repository response
func (m *ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster namespace tanzupackage repository delete repository response based on context it is used
func (m *ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageRepositoryDeleteRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}