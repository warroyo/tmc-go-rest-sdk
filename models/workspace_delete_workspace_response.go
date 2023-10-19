// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspaceDeleteWorkspaceResponse Response from deleting a Workspace.
//
// swagger:model workspace.DeleteWorkspaceResponse
type WorkspaceDeleteWorkspaceResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this workspace delete workspace response
func (m *WorkspaceDeleteWorkspaceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workspace delete workspace response based on context it is used
func (m *WorkspaceDeleteWorkspaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceDeleteWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceDeleteWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res WorkspaceDeleteWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}