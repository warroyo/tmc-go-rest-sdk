// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonBatchDetails Details contains information about a source resource being applied on its atomic targets.
//
// swagger:model common.batch.Details
type CommonBatchDetails struct {

	// Number of atomic targets on which this source resource is successfully applied.
	Applied int32 `json:"applied,omitempty"`

	// Total number of targets available for this source resource.
	AvailableTargets int32 `json:"availableTargets,omitempty"`

	// Number of atomic targets on which this source resource is currently being deleted (only applicable on some source resource types).
	Deleting int32 `json:"deleting,omitempty"`

	// Number of atomic targets on which this source resource failed to apply due to some error.
	Error int32 `json:"error,omitempty"`

	// Number of atomic targets on which this source resource is overridden by another resource.
	Overridden int32 `json:"overridden,omitempty"`

	// Number of atomic targets on which this source resource is still being applied.
	Pending int32 `json:"pending,omitempty"`
}

// Validate validates this common batch details
func (m *CommonBatchDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common batch details based on context it is used
func (m *CommonBatchDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonBatchDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonBatchDetails) UnmarshalBinary(b []byte) error {
	var res CommonBatchDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
