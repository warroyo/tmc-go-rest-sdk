// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationFluxcdHelmRepositorySearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model organization.fluxcd.helm.repository.SearchScope
type OrganizationFluxcdHelmRepositorySearchScope struct {

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`
}

// Validate validates this organization fluxcd helm repository search scope
func (m *OrganizationFluxcdHelmRepositorySearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization fluxcd helm repository search scope based on context it is used
func (m *OrganizationFluxcdHelmRepositorySearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationFluxcdHelmRepositorySearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationFluxcdHelmRepositorySearchScope) UnmarshalBinary(b []byte) error {
	var res OrganizationFluxcdHelmRepositorySearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
