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

// PolicyTemplateCreateTemplateResponse Response from creating a Template.
//
// swagger:model policy.template.CreateTemplateResponse
type PolicyTemplateCreateTemplateResponse struct {

	// Template created.
	Template *PolicyTemplateTemplate `json:"template,omitempty"`
}

// Validate validates this policy template create template response
func (m *PolicyTemplateCreateTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTemplateCreateTemplateResponse) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy template create template response based on the context it is used
func (m *PolicyTemplateCreateTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTemplateCreateTemplateResponse) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTemplateCreateTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplateCreateTemplateResponse) UnmarshalBinary(b []byte) error {
	var res PolicyTemplateCreateTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
