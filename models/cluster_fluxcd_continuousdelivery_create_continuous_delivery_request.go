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

// ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest Request to create a ContinuousDelivery.
//
// swagger:model cluster.fluxcd.continuousdelivery.CreateContinuousDeliveryRequest
type ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest struct {

	// ContinuousDelivery to create.
	ContinuousDelivery *ClusterFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDelivery,omitempty"`
}

// Validate validates this cluster fluxcd continuousdelivery create continuous delivery request
func (m *ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContinuousDelivery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest) validateContinuousDelivery(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster fluxcd continuousdelivery create continuous delivery request based on the context it is used
func (m *ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContinuousDelivery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest) contextValidateContinuousDelivery(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest) UnmarshalBinary(b []byte) error {
	var res ClusterFluxcdContinuousdeliveryCreateContinuousDeliveryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}