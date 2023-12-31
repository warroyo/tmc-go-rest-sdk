// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterVirtualMachineImage Virtual machine image for the TKG Service vSphere provider.
//
// swagger:model common.cluster.VirtualMachineImage
type CommonClusterVirtualMachineImage struct {

	// Name of the virtual machine image.
	Name string `json:"name,omitempty"`
}

// Validate validates this common cluster virtual machine image
func (m *CommonClusterVirtualMachineImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster virtual machine image based on context it is used
func (m *CommonClusterVirtualMachineImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterVirtualMachineImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterVirtualMachineImage) UnmarshalBinary(b []byte) error {
	var res CommonClusterVirtualMachineImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
