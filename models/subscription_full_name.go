// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubscriptionFullName Full name of the subscription. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model subscription.FullName
type SubscriptionFullName struct {

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this subscription full name
func (m *SubscriptionFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscription full name based on context it is used
func (m *SubscriptionFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionFullName) UnmarshalBinary(b []byte) error {
	var res SubscriptionFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}