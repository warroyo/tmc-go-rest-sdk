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

// ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse Response from creating a ContinuousDelivery.
//
// swagger:model clustergroup.fluxcd.continuousdelivery.CreateContinuousDeliveryResponse
type ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse struct {

	// ContinuousDelivery created.
	ContinuousDelivery *ClustergroupFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDelivery,omitempty"`
}

// Validate validates this clustergroup fluxcd continuousdelivery create continuous delivery response
func (m *ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContinuousDelivery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse) validateContinuousDelivery(formats strfmt.Registry) error {
	if swag.IsZero(m.ContinuousDelivery) { // not required
		return nil
	}

	if m.ContinuousDelivery != nil {
		if err := m.ContinuousDelivery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("continuousDelivery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("continuousDelivery")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup fluxcd continuousdelivery create continuous delivery response based on the context it is used
func (m *ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContinuousDelivery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse) contextValidateContinuousDelivery(ctx context.Context, formats strfmt.Registry) error {

	if m.ContinuousDelivery != nil {
		if err := m.ContinuousDelivery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("continuousDelivery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("continuousDelivery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupFluxcdContinuousdeliveryCreateContinuousDeliveryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
