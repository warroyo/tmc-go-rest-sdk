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

// ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse Response from updating a GitRepository.
//
// swagger:model clustergroup.namespace.fluxcd.gitrepository.UpdateGitRepositoryResponse
type ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse struct {

	// GitRepository updated.
	GitRepository *ClustergroupNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// Validate validates this clustergroup namespace fluxcd gitrepository update git repository response
func (m *ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGitRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse) validateGitRepository(formats strfmt.Registry) error {
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

// ContextValidate validate this clustergroup namespace fluxcd gitrepository update git repository response based on the context it is used
func (m *ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGitRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse) contextValidateGitRepository(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
