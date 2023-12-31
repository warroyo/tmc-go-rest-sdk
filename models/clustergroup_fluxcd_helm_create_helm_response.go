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

// ClustergroupFluxcdHelmCreateHelmResponse Response from creating a Helm.
//
// swagger:model clustergroup.fluxcd.helm.CreateHelmResponse
type ClustergroupFluxcdHelmCreateHelmResponse struct {

	// Helm created.
	Helm *ClustergroupFluxcdHelmHelm `json:"helm,omitempty"`
}

// Validate validates this clustergroup fluxcd helm create helm response
func (m *ClustergroupFluxcdHelmCreateHelmResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHelm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdHelmCreateHelmResponse) validateHelm(formats strfmt.Registry) error {
	if swag.IsZero(m.Helm) { // not required
		return nil
	}

	if m.Helm != nil {
		if err := m.Helm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup fluxcd helm create helm response based on the context it is used
func (m *ClustergroupFluxcdHelmCreateHelmResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHelm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdHelmCreateHelmResponse) contextValidateHelm(ctx context.Context, formats strfmt.Registry) error {

	if m.Helm != nil {
		if err := m.Helm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupFluxcdHelmCreateHelmResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupFluxcdHelmCreateHelmResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupFluxcdHelmCreateHelmResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
