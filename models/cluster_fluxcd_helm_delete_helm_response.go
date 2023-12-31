// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterFluxcdHelmDeleteHelmResponse Response from deleting a Helm.
//
// swagger:model cluster.fluxcd.helm.DeleteHelmResponse
type ClusterFluxcdHelmDeleteHelmResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this cluster fluxcd helm delete helm response
func (m *ClusterFluxcdHelmDeleteHelmResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster fluxcd helm delete helm response based on context it is used
func (m *ClusterFluxcdHelmDeleteHelmResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterFluxcdHelmDeleteHelmResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFluxcdHelmDeleteHelmResponse) UnmarshalBinary(b []byte) error {
	var res ClusterFluxcdHelmDeleteHelmResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
