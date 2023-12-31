// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MeteringUsageUsage Usage and limit of resource.
//
// swagger:model metering.usage.Usage
type MeteringUsageUsage struct {

	// Max limit on the resource usage.
	Limit float32 `json:"limit,omitempty"`

	// Current usage of the resource.
	Usage float32 `json:"usage,omitempty"`
}

// Validate validates this metering usage usage
func (m *MeteringUsageUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metering usage usage based on context it is used
func (m *MeteringUsageUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeteringUsageUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteringUsageUsage) UnmarshalBinary(b []byte) error {
	var res MeteringUsageUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
