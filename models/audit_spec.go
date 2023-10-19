// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditSpec The audit spec.
//
// swagger:model audit.Spec
type AuditSpec struct {

	// The end timestamp of the audit period for which to collect audit logs.
	// Defaults to today.
	// Format: date-time
	EndDatetime strfmt.DateTime `json:"endDatetime,omitempty"`

	// The start timestamp of the audit period for which to collect audit logs.
	// Defaults to 60 days before today.
	// Format: date-time
	StartDatetime strfmt.DateTime `json:"startDatetime,omitempty"`
}

// Validate validates this audit spec
func (m *AuditSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSpec) validateEndDatetime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("endDatetime", "body", "date-time", m.EndDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditSpec) validateStartDatetime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("startDatetime", "body", "date-time", m.StartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this audit spec based on context it is used
func (m *AuditSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditSpec) UnmarshalBinary(b []byte) error {
	var res AuditSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
