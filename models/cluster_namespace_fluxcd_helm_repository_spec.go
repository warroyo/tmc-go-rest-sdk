// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceFluxcdHelmRepositorySpec Spec of the helm repository.
//
// swagger:model cluster.namespace.fluxcd.helm.repository.Spec
type ClusterNamespaceFluxcdHelmRepositorySpec struct {

	// URL of helm repository.
	URL string `json:"url,omitempty"`
}

// Validate validates this cluster namespace fluxcd helm repository spec
func (m *ClusterNamespaceFluxcdHelmRepositorySpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster namespace fluxcd helm repository spec based on context it is used
func (m *ClusterNamespaceFluxcdHelmRepositorySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdHelmRepositorySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdHelmRepositorySpec) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceFluxcdHelmRepositorySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}