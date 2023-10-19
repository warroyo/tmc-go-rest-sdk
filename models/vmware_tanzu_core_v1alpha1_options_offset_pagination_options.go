// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions Options to paginate a response using offsets.
//
// swagger:model vmware.tanzu.core.v1alpha1.options.OffsetPaginationOptions
type VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions struct {

	// Offset at which to start returning records.
	Offset string `json:"offset,omitempty"`

	// Number of records to return.
	Size string `json:"size,omitempty"`
}

// Validate validates this vmware tanzu core v1alpha1 options offset pagination options
func (m *VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vmware tanzu core v1alpha1 options offset pagination options based on context it is used
func (m *VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}