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

// AuditListAuditsResponse Response from listing Audits.
//
// swagger:model audit.ListAuditsResponse
type AuditListAuditsResponse struct {

	// List of audits.
	Audits []*AuditAudit `json:"audits"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this audit list audits response
func (m *AuditListAuditsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditListAuditsResponse) validateAudits(formats strfmt.Registry) error {
	if swag.IsZero(m.Audits) { // not required
		return nil
	}

	for i := 0; i < len(m.Audits); i++ {
		if swag.IsZero(m.Audits[i]) { // not required
			continue
		}

		if m.Audits[i] != nil {
			if err := m.Audits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("audits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("audits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this audit list audits response based on the context it is used
func (m *AuditListAuditsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditListAuditsResponse) contextValidateAudits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Audits); i++ {

		if m.Audits[i] != nil {
			if err := m.Audits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("audits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("audits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditListAuditsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditListAuditsResponse) UnmarshalBinary(b []byte) error {
	var res AuditListAuditsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}