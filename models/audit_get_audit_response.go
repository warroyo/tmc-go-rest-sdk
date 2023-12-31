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

// AuditGetAuditResponse Response from getting an Audit.
//
// swagger:model audit.GetAuditResponse
type AuditGetAuditResponse struct {

	// Audit returned.
	Audit *AuditAudit `json:"audit,omitempty"`
}

// Validate validates this audit get audit response
func (m *AuditGetAuditResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditGetAuditResponse) validateAudit(formats strfmt.Registry) error {
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

// ContextValidate validate this audit get audit response based on the context it is used
func (m *AuditGetAuditResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditGetAuditResponse) contextValidateAudit(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AuditGetAuditResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditGetAuditResponse) UnmarshalBinary(b []byte) error {
	var res AuditGetAuditResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
