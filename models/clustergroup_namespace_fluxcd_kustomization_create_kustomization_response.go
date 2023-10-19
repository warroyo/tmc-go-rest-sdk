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

// ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse Response from creating a Kustomization.
//
// swagger:model clustergroup.namespace.fluxcd.kustomization.CreateKustomizationResponse
type ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse struct {

	// Kustomization created.
	Kustomization *ClustergroupNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// Validate validates this clustergroup namespace fluxcd kustomization create kustomization response
func (m *ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKustomization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse) validateKustomization(formats strfmt.Registry) error {
	if swag.IsZero(m.Kustomization) { // not required
		return nil
	}

	if m.Kustomization != nil {
		if err := m.Kustomization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kustomization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kustomization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup namespace fluxcd kustomization create kustomization response based on the context it is used
func (m *ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKustomization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse) contextValidateKustomization(ctx context.Context, formats strfmt.Registry) error {

	if m.Kustomization != nil {
		if err := m.Kustomization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kustomization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kustomization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceFluxcdKustomizationCreateKustomizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
