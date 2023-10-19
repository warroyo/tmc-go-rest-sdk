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

// ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse Response from listing ContinuousDeliveries.
//
// swagger:model cluster.fluxcd.continuousdelivery.ListContinuousDeliveriesResponse
type ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse struct {

	// List of continuousdeliveries.
	ContinuousDeliveries []*ClusterFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDeliveries"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this cluster fluxcd continuousdelivery list continuous deliveries response
func (m *ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContinuousDeliveries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) validateContinuousDeliveries(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster fluxcd continuousdelivery list continuous deliveries response based on the context it is used
func (m *ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContinuousDeliveries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) contextValidateContinuousDeliveries(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) UnmarshalBinary(b []byte) error {
	var res ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
