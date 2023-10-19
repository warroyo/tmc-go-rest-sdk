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

// ClustergroupNamespaceSecretexportCreateSecretExportResponse Response from creating a SecretExport.
//
// swagger:model clustergroup.namespace.secretexport.CreateSecretExportResponse
type ClustergroupNamespaceSecretexportCreateSecretExportResponse struct {

	// SecretExport created.
	SecretExport *ClustergroupNamespaceSecretexportSecretExport `json:"secretExport,omitempty"`
}

// Validate validates this clustergroup namespace secretexport create secret export response
func (m *ClustergroupNamespaceSecretexportCreateSecretExportResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretExport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceSecretexportCreateSecretExportResponse) validateSecretExport(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretExport) { // not required
		return nil
	}

	if m.SecretExport != nil {
		if err := m.SecretExport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretExport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secretExport")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup namespace secretexport create secret export response based on the context it is used
func (m *ClustergroupNamespaceSecretexportCreateSecretExportResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretExport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceSecretexportCreateSecretExportResponse) contextValidateSecretExport(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretExport != nil {
		if err := m.SecretExport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretExport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secretExport")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretexportCreateSecretExportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretexportCreateSecretExportResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceSecretexportCreateSecretExportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
