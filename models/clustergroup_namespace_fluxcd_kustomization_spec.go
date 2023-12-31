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

// ClustergroupNamespaceFluxcdKustomizationSpec Spec for the Kustomization.
//
// swagger:model clustergroup.namespace.fluxcd.kustomization.Spec
type ClustergroupNamespaceFluxcdKustomizationSpec struct {

	// Spec for the Kustomization as defined at atomic level.
	AtomicSpec *ClusterNamespaceFluxcdKustomizationSpec `json:"atomicSpec,omitempty"`
}

// Validate validates this clustergroup namespace fluxcd kustomization spec
func (m *ClustergroupNamespaceFluxcdKustomizationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtomicSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdKustomizationSpec) validateAtomicSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.AtomicSpec) { // not required
		return nil
	}

	if m.AtomicSpec != nil {
		if err := m.AtomicSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atomicSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("atomicSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup namespace fluxcd kustomization spec based on the context it is used
func (m *ClustergroupNamespaceFluxcdKustomizationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAtomicSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdKustomizationSpec) contextValidateAtomicSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.AtomicSpec != nil {
		if err := m.AtomicSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atomicSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("atomicSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdKustomizationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdKustomizationSpec) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceFluxcdKustomizationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
