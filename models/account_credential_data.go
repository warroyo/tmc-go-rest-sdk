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

// AccountCredentialData Credential data.
//
// swagger:model account.credential.Data
type AccountCredentialData struct {

	// AWS credential.
	AwsCredential *AccountCredentialTypeAwsSpec `json:"awsCredential,omitempty"`

	// Azure credential.
	AzureCredential *AccountCredentialTypeAzureSpec `json:"azureCredential,omitempty"`

	// Generic credential.
	GenericCredential string `json:"genericCredential,omitempty"`

	// Key Value credential.
	KeyValue *AccountCredentialTypeKeyvalueSpec `json:"keyValue,omitempty"`
}

// Validate validates this account credential data
func (m *AccountCredentialData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialData) validateAwsCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsCredential) { // not required
		return nil
	}

	if m.AwsCredential != nil {
		if err := m.AwsCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsCredential")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialData) validateAzureCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureCredential) { // not required
		return nil
	}

	if m.AzureCredential != nil {
		if err := m.AzureCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureCredential")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialData) validateKeyValue(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyValue) { // not required
		return nil
	}

	if m.KeyValue != nil {
		if err := m.KeyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyValue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account credential data based on the context it is used
func (m *AccountCredentialData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialData) contextValidateAwsCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsCredential != nil {
		if err := m.AwsCredential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsCredential")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialData) contextValidateAzureCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureCredential != nil {
		if err := m.AzureCredential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureCredential")
			}
			return err
		}
	}

	return nil
}

func (m *AccountCredentialData) contextValidateKeyValue(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyValue != nil {
		if err := m.KeyValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialData) UnmarshalBinary(b []byte) error {
	var res AccountCredentialData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
