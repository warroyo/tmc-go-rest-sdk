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

// ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse Response from listing ContinuousDeliveries.
//
// swagger:model clustergroup.fluxcd.continuousdelivery.ListContinuousDeliveriesResponse
type ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse struct {

	// List of continuousdeliveries.
	ContinuousDeliveries []*ClustergroupFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDeliveries"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this clustergroup fluxcd continuousdelivery list continuous deliveries response
func (m *ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContinuousDeliveries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) validateContinuousDeliveries(formats strfmt.Registry) error {
	if swag.IsZero(m.ContinuousDeliveries) { // not required
		return nil
	}

	for i := 0; i < len(m.ContinuousDeliveries); i++ {
		if swag.IsZero(m.ContinuousDeliveries[i]) { // not required
			continue
		}

		if m.ContinuousDeliveries[i] != nil {
			if err := m.ContinuousDeliveries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("continuousDeliveries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("continuousDeliveries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this clustergroup fluxcd continuousdelivery list continuous deliveries response based on the context it is used
func (m *ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContinuousDeliveries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) contextValidateContinuousDeliveries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContinuousDeliveries); i++ {

		if m.ContinuousDeliveries[i] != nil {
			if err := m.ContinuousDeliveries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("continuousDeliveries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("continuousDeliveries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
