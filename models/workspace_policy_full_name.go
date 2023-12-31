// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspacePolicyFullName Full name of the workspace policy. This includes the object
// name along with any parents or further identifiers.
//
// swagger:model workspace.policy.FullName
type WorkspacePolicyFullName struct {

	// Name of the policy.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the workspace.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this workspace policy full name
func (m *WorkspacePolicyFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workspace policy full name based on context it is used
func (m *WorkspacePolicyFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspacePolicyFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspacePolicyFullName) UnmarshalBinary(b []byte) error {
	var res WorkspacePolicyFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
