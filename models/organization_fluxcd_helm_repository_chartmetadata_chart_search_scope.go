// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model organization.fluxcd.helm.repository.chartmetadata.chart.SearchScope
type OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope struct {

	// Scope search to the specified chart_metadata_name; supports globbing; default (*).
	ChartMetadataName string `json:"chartMetadataName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified repository_name; supports globbing; default (*).
	RepositoryName string `json:"repositoryName,omitempty"`
}

// Validate validates this organization fluxcd helm repository chartmetadata chart search scope
func (m *OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization fluxcd helm repository chartmetadata chart search scope based on context it is used
func (m *OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) UnmarshalBinary(b []byte) error {
	var res OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}