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

// ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest Request to update (overwrite) a Repository.
//
// swagger:model cluster.namespace.tanzupackage.repository.UpdateRepositoryRequest
type ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest struct {

	// Update Repository.
	Repository *ClusterNamespaceTanzupackageRepositoryRepository `json:"repository,omitempty"`
}

// Validate validates this cluster namespace tanzupackage repository update repository request
func (m *ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest) validateRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster namespace tanzupackage repository update repository request based on the context it is used
func (m *ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {
		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageRepositoryUpdateRepositoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
