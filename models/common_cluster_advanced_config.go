// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterAdvancedConfig The advanced configuration for TKGm cluster.
//
// swagger:model common.cluster.AdvancedConfig
type CommonClusterAdvancedConfig struct {

	// The key of the advanced configuration parameters.
	Key string `json:"key,omitempty"`

	// The value of the advanced configuration parameters.
	Value string `json:"value,omitempty"`
}

// Validate validates this common cluster advanced config
func (m *CommonClusterAdvancedConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster advanced config based on context it is used
func (m *CommonClusterAdvancedConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterAdvancedConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterAdvancedConfig) UnmarshalBinary(b []byte) error {
	var res CommonClusterAdvancedConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
