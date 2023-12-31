// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterAdvancedConfigOption The options of TKGm advanced configuration.
//
// swagger:model common.cluster.AdvancedConfigOption
type CommonClusterAdvancedConfigOption struct {

	// The default value of the advanced configuration parameters.
	Default string `json:"default,omitempty"`

	// The description of the advanced configuration parameters.
	Description string `json:"description,omitempty"`

	// The key of the advanced configuration parameters.
	Key string `json:"key,omitempty"`
}

// Validate validates this common cluster advanced config option
func (m *CommonClusterAdvancedConfigOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster advanced config option based on context it is used
func (m *CommonClusterAdvancedConfigOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterAdvancedConfigOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterAdvancedConfigOption) UnmarshalBinary(b []byte) error {
	var res CommonClusterAdvancedConfigOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
