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

// ClustergroupIntegrationListIntegrationsResponse Response from listing Integrations.
//
// swagger:model clustergroup.integration.ListIntegrationsResponse
type ClustergroupIntegrationListIntegrationsResponse struct {

	// List of integrations.
	Integrations []*ClustergroupIntegrationIntegration `json:"integrations"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this clustergroup integration list integrations response
func (m *ClustergroupIntegrationListIntegrationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupIntegrationListIntegrationsResponse) validateIntegrations(formats strfmt.Registry) error {
	if swag.IsZero(m.Integrations) { // not required
		return nil
	}

	for i := 0; i < len(m.Integrations); i++ {
		if swag.IsZero(m.Integrations[i]) { // not required
			continue
		}

		if m.Integrations[i] != nil {
			if err := m.Integrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this clustergroup integration list integrations response based on the context it is used
func (m *ClustergroupIntegrationListIntegrationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupIntegrationListIntegrationsResponse) contextValidateIntegrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Integrations); i++ {

		if m.Integrations[i] != nil {
			if err := m.Integrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupIntegrationListIntegrationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupIntegrationListIntegrationsResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupIntegrationListIntegrationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
