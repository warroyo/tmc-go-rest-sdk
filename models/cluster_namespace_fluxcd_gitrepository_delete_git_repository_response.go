// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse Response from deleting a GitRepository.
//
// swagger:model cluster.namespace.fluxcd.gitrepository.DeleteGitRepositoryResponse
type ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this cluster namespace fluxcd gitrepository delete git repository response
func (m *ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster namespace fluxcd gitrepository delete git repository response based on context it is used
func (m *ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
