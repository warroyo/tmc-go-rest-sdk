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

// AuditCreateAuditRequest Request to create an Audit.
//
// swagger:model audit.CreateAuditRequest
type AuditCreateAuditRequest struct {

	// Audit to create.
	Audit *AuditAudit `json:"audit,omitempty"`
}

// Validate validates this audit create audit request
func (m *AuditCreateAuditRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditCreateAuditRequest) validateAudit(formats strfmt.Registry) error {
	if swag.IsZero(m.Audit) { // not required
		return nil
	}

	if m.Audit != nil {
		if err := m.Audit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this audit create audit request based on the context it is used
func (m *AuditCreateAuditRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditCreateAuditRequest) contextValidateAudit(ctx context.Context, formats strfmt.Registry) error {

	if m.Audit != nil {
		if err := m.Audit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditCreateAuditRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditCreateAuditRequest) UnmarshalBinary(b []byte) error {
	var res AuditCreateAuditRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}