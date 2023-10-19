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

// ClustergroupFluxcdSourcesecretSpec Spec for the Source Secret.
//
// swagger:model clustergroup.fluxcd.sourcesecret.Spec
type ClustergroupFluxcdSourcesecretSpec struct {

	// Spec of the source secret defined at atomic level.
	AtomicSpec *ClusterFluxcdSourcesecretSpec `json:"atomicSpec,omitempty"`
}

// Validate validates this clustergroup fluxcd sourcesecret spec
func (m *ClustergroupFluxcdSourcesecretSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtomicSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdSourcesecretSpec) validateAtomicSpec(formats strfmt.Registry) error {
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

// ContextValidate validate this clustergroup fluxcd sourcesecret spec based on the context it is used
func (m *ClustergroupFluxcdSourcesecretSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAtomicSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdSourcesecretSpec) contextValidateAtomicSpec(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClustergroupFluxcdSourcesecretSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupFluxcdSourcesecretSpec) UnmarshalBinary(b []byte) error {
	var res ClustergroupFluxcdSourcesecretSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
