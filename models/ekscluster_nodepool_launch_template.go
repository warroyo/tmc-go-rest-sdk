// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EksclusterNodepoolLaunchTemplate Launch template for the nodepool.
//
// swagger:model ekscluster.nodepool.LaunchTemplate
type EksclusterNodepoolLaunchTemplate struct {

	// The ID of the launch template.
	ID string `json:"id,omitempty"`

	// The name of the launch template.
	Name string `json:"name,omitempty"`

	// The version of the launch template to use. If no version is specified, then the template's default version is used.
	Version string `json:"version,omitempty"`
}

// Validate validates this ekscluster nodepool launch template
func (m *EksclusterNodepoolLaunchTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ekscluster nodepool launch template based on context it is used
func (m *EksclusterNodepoolLaunchTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterNodepoolLaunchTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterNodepoolLaunchTemplate) UnmarshalBinary(b []byte) error {
	var res EksclusterNodepoolLaunchTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
