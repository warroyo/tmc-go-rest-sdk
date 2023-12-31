// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterTKGVsphereFolders TKG vSphere virtual machine folders.
//
// swagger:model common.cluster.TKGVsphereFolders
type CommonClusterTKGVsphereFolders struct {

	// Name of the vSphere object.
	Name string `json:"name,omitempty"`

	// Path to the object in the vSphere API.
	Path string `json:"path,omitempty"`
}

// Validate validates this common cluster t k g vsphere folders
func (m *CommonClusterTKGVsphereFolders) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster t k g vsphere folders based on context it is used
func (m *CommonClusterTKGVsphereFolders) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGVsphereFolders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGVsphereFolders) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGVsphereFolders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
