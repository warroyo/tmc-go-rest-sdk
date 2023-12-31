// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest Request to update (overwrite) a GitRepository.
//
// swagger:model cluster.namespace.fluxcd.gitrepository.UpdateGitRepositoryRequest
type ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest struct {

	// Update GitRepository.
	GitRepository *ClusterNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// Validate validates this cluster namespace fluxcd gitrepository update git repository request
func (m *ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGitRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) validateGitRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.GitRepository) { // not required
		return nil
	}

	if m.GitRepository != nil {
		if err := m.GitRepository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitRepository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitRepository")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster namespace fluxcd gitrepository update git repository request based on the context it is used
func (m *ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGitRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) contextValidateGitRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.GitRepository != nil {
		if err := m.GitRepository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitRepository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitRepository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
