// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterFluxcdSourcesecretDeleteSourceSecretResponse Response from deleting a SourceSecret.
//
// swagger:model cluster.fluxcd.sourcesecret.DeleteSourceSecretResponse
type ClusterFluxcdSourcesecretDeleteSourceSecretResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this cluster fluxcd sourcesecret delete source secret response
func (m *ClusterFluxcdSourcesecretDeleteSourceSecretResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster fluxcd sourcesecret delete source secret response based on context it is used
func (m *ClusterFluxcdSourcesecretDeleteSourceSecretResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterFluxcdSourcesecretDeleteSourceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFluxcdSourcesecretDeleteSourceSecretResponse) UnmarshalBinary(b []byte) error {
	var res ClusterFluxcdSourcesecretDeleteSourceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}