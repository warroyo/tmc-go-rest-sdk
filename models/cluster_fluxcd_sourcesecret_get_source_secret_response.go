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

// ClusterFluxcdSourcesecretGetSourceSecretResponse Response from getting a SourceSecret.
//
// swagger:model cluster.fluxcd.sourcesecret.GetSourceSecretResponse
type ClusterFluxcdSourcesecretGetSourceSecretResponse struct {

	// SourceSecret returned.
	SourceSecret *ClusterFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// Validate validates this cluster fluxcd sourcesecret get source secret response
func (m *ClusterFluxcdSourcesecretGetSourceSecretResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdSourcesecretGetSourceSecretResponse) validateSourceSecret(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster fluxcd sourcesecret get source secret response based on the context it is used
func (m *ClusterFluxcdSourcesecretGetSourceSecretResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdSourcesecretGetSourceSecretResponse) contextValidateSourceSecret(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterFluxcdSourcesecretGetSourceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFluxcdSourcesecretGetSourceSecretResponse) UnmarshalBinary(b []byte) error {
	var res ClusterFluxcdSourcesecretGetSourceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
