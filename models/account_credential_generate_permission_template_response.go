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

// AccountCredentialGeneratePermissionTemplateResponse Response containing the generated permission template.
//
// swagger:model account.credential.GeneratePermissionTemplateResponse
type AccountCredentialGeneratePermissionTemplateResponse struct {

	// The Tanzu capability for which the credential shall be used.
	Capability string `json:"capability,omitempty"`

	// The full name of the credential that this template is for.
	FullName *AccountCredentialFullName `json:"fullName,omitempty"`

	// Base64 encoded permission template.
	PermissionTemplate string `json:"permissionTemplate,omitempty"`

	// The infrastructure provider that the permission template is for.
	Provider *AccountCredentialProvider `json:"provider,omitempty"`

	// URL for permission template.
	TemplateURL string `json:"templateUrl,omitempty"`

	// Values which helps in forming the CLI command output and displaying values on UI.
	TemplateValues map[string]string `json:"templateValues,omitempty"`
}

// Validate validates this account credential generate permission template response
func (m *AccountCredentialGeneratePermissionTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialGeneratePermissionTemplateResponse) validateFullName(formats strfmt.Registry) error {
	if swag.IsZero(m.FullName) { // not required
		return nil
	}

	if m.FullName != nil {
		if err := m.FullName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialGeneratePermissionTemplateResponse) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account credential generate permission template response based on the context it is used
func (m *AccountCredentialGeneratePermissionTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialGeneratePermissionTemplateResponse) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

	if m.FullName != nil {
		if err := m.FullName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialGeneratePermissionTemplateResponse) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.Provider != nil {
		if err := m.Provider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialGeneratePermissionTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialGeneratePermissionTemplateResponse) UnmarshalBinary(b []byte) error {
	var res AccountCredentialGeneratePermissionTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}