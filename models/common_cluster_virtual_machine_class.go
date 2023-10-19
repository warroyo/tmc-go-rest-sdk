// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterVirtualMachineClass Virtual machine class for the TKG Service vSphere provider.
//
// swagger:model common.cluster.VirtualMachineClass
type CommonClusterVirtualMachineClass struct {

	// CPU cores provided by the virtual machine class.
	CPUCores string `json:"cpuCores,omitempty"`

	// Memory provided by the virtual machine class.
	MemoryGb string `json:"memoryGb,omitempty"`

	// Name of the virtual machine class.
	Name string `json:"name,omitempty"`
}

// Validate validates this common cluster virtual machine class
func (m *CommonClusterVirtualMachineClass) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster virtual machine class based on context it is used
func (m *CommonClusterVirtualMachineClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterVirtualMachineClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterVirtualMachineClass) UnmarshalBinary(b []byte) error {
	var res CommonClusterVirtualMachineClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}