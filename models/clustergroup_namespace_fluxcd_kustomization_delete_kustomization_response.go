// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse Response from deleting a Kustomization.
//
// swagger:model clustergroup.namespace.fluxcd.kustomization.DeleteKustomizationResponse
type ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this clustergroup namespace fluxcd kustomization delete kustomization response
func (m *ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup namespace fluxcd kustomization delete kustomization response based on context it is used
func (m *ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceFluxcdKustomizationDeleteKustomizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
