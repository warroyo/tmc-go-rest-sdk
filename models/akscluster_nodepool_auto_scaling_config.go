// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AksclusterNodepoolAutoScalingConfig Auto scaling config for the nodepool.
//
// swagger:model akscluster.nodepool.AutoScalingConfig
type AksclusterNodepoolAutoScalingConfig struct {

	// Whether to enable auto-scaler.
	Enabled bool `json:"enabled,omitempty"`

	// The maximum number of nodes for auto-scaling.
	MaxCount int32 `json:"maxCount,omitempty"`

	// The minimum number of nodes for auto-scaling.
	MinCount int32 `json:"minCount,omitempty"`
}

// Validate validates this akscluster nodepool auto scaling config
func (m *AksclusterNodepoolAutoScalingConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this akscluster nodepool auto scaling config based on context it is used
func (m *AksclusterNodepoolAutoScalingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterNodepoolAutoScalingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterNodepoolAutoScalingConfig) UnmarshalBinary(b []byte) error {
	var res AksclusterNodepoolAutoScalingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
