// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse Response from deleting a GitRepository.
//
// swagger:model clustergroup.namespace.fluxcd.gitrepository.DeleteGitRepositoryResponse
type ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this clustergroup namespace fluxcd gitrepository delete git repository response
func (m *ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup namespace fluxcd gitrepository delete git repository response based on context it is used
func (m *ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
