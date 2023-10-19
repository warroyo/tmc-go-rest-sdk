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

// ClustergroupFluxcdSourcesecretGetSourceSecretResponse Response from getting a SourceSecret.
//
// swagger:model clustergroup.fluxcd.sourcesecret.GetSourceSecretResponse
type ClustergroupFluxcdSourcesecretGetSourceSecretResponse struct {

	// SourceSecret returned.
	SourceSecret *ClustergroupFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// Validate validates this clustergroup fluxcd sourcesecret get source secret response
func (m *ClustergroupFluxcdSourcesecretGetSourceSecretResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdSourcesecretGetSourceSecretResponse) validateSourceSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceSecret) { // not required
		return nil
	}

	if m.SourceSecret != nil {
		if err := m.SourceSecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceSecret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceSecret")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup fluxcd sourcesecret get source secret response based on the context it is used
func (m *ClustergroupFluxcdSourcesecretGetSourceSecretResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdSourcesecretGetSourceSecretResponse) contextValidateSourceSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceSecret != nil {
		if err := m.SourceSecret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceSecret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceSecret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupFluxcdSourcesecretGetSourceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupFluxcdSourcesecretGetSourceSecretResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupFluxcdSourcesecretGetSourceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
