// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse Response from deleting a ContinuousDelivery.
//
// swagger:model clustergroup.fluxcd.continuousdelivery.DeleteContinuousDeliveryResponse
type ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this clustergroup fluxcd continuousdelivery delete continuous delivery response
func (m *ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup fluxcd continuousdelivery delete continuous delivery response based on context it is used
func (m *ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupFluxcdContinuousdeliveryDeleteContinuousDeliveryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
