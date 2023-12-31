// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterFluxcdSourcesecretListSourceSecretsResponse Response from listing SourceSecrets.
//
// swagger:model cluster.fluxcd.sourcesecret.ListSourceSecretsResponse
type ClusterFluxcdSourcesecretListSourceSecretsResponse struct {

	// List of sourcesecrets.
	SourceSecrets []*ClusterFluxcdSourcesecretSourceSecret `json:"sourceSecrets"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this cluster fluxcd sourcesecret list source secrets response
func (m *ClusterFluxcdSourcesecretListSourceSecretsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceSecrets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdSourcesecretListSourceSecretsResponse) validateSourceSecrets(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceSecrets) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceSecrets); i++ {
		if swag.IsZero(m.SourceSecrets[i]) { // not required
			continue
		}

		if m.SourceSecrets[i] != nil {
			if err := m.SourceSecrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceSecrets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceSecrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster fluxcd sourcesecret list source secrets response based on the context it is used
func (m *ClusterFluxcdSourcesecretListSourceSecretsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceSecrets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdSourcesecretListSourceSecretsResponse) contextValidateSourceSecrets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SourceSecrets); i++ {

		if m.SourceSecrets[i] != nil {
			if err := m.SourceSecrets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceSecrets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceSecrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterFluxcdSourcesecretListSourceSecretsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFluxcdSourcesecretListSourceSecretsResponse) UnmarshalBinary(b []byte) error {
	var res ClusterFluxcdSourcesecretListSourceSecretsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
