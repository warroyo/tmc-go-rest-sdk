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

// ClustergroupNamespaceSecretCreateSecretResponse Response from creating a Secret.
//
// swagger:model clustergroup.namespace.secret.CreateSecretResponse
type ClustergroupNamespaceSecretCreateSecretResponse struct {

	// Secret created.
	Secret *ClustergroupNamespaceSecretSecret `json:"secret,omitempty"`
}

// Validate validates this clustergroup namespace secret create secret response
func (m *ClustergroupNamespaceSecretCreateSecretResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceSecretCreateSecretResponse) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup namespace secret create secret response based on the context it is used
func (m *ClustergroupNamespaceSecretCreateSecretResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceSecretCreateSecretResponse) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {
		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretCreateSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceSecretCreateSecretResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceSecretCreateSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}