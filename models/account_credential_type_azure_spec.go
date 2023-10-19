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

// AccountCredentialTypeAzureSpec Azure credential spec.
//
// swagger:model account.credential.type.azure.Spec
type AccountCredentialTypeAzureSpec struct {

	// Service Principal.
	ServicePrincipal *AccountCredentialTypeAzureServicePrincipal `json:"servicePrincipal,omitempty"`

	// ServicePrincipalWithCertificate.
	ServicePrincipalWithCertificate *AccountCredentialTypeAzureServicePrincipalWithCertificate `json:"servicePrincipalWithCertificate,omitempty"`
}

// Validate validates this account credential type azure spec
func (m *AccountCredentialTypeAzureSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServicePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePrincipalWithCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialTypeAzureSpec) validateServicePrincipal(formats strfmt.Registry) error {
	if swag.IsZero(m.ServicePrincipal) { // not required
		return nil
	}

	if m.ServicePrincipal != nil {
		if err := m.ServicePrincipal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicePrincipal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("servicePrincipal")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialTypeAzureSpec) validateServicePrincipalWithCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.ServicePrincipalWithCertificate) { // not required
		return nil
	}

	if m.ServicePrincipalWithCertificate != nil {
		if err := m.ServicePrincipalWithCertificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicePrincipalWithCertificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("servicePrincipalWithCertificate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account credential type azure spec based on the context it is used
func (m *AccountCredentialTypeAzureSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServicePrincipal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServicePrincipalWithCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialTypeAzureSpec) contextValidateServicePrincipal(ctx context.Context, formats strfmt.Registry) error {

	if m.ServicePrincipal != nil {
		if err := m.ServicePrincipal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicePrincipal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("servicePrincipal")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialTypeAzureSpec) contextValidateServicePrincipalWithCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.ServicePrincipalWithCertificate != nil {
		if err := m.ServicePrincipalWithCertificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicePrincipalWithCertificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("servicePrincipalWithCertificate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialTypeAzureSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialTypeAzureSpec) UnmarshalBinary(b []byte) error {
	var res AccountCredentialTypeAzureSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
