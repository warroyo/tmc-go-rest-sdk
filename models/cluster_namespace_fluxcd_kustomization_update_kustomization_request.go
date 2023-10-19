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

// ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest Request to update (overwrite) a Kustomization.
//
// swagger:model cluster.namespace.fluxcd.kustomization.UpdateKustomizationRequest
type ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest struct {

	// Update Kustomization.
	Kustomization *ClusterNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// Validate validates this cluster namespace fluxcd kustomization update kustomization request
func (m *ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKustomization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) validateKustomization(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster namespace fluxcd kustomization update kustomization request based on the context it is used
func (m *ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKustomization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) contextValidateKustomization(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}