// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspacePolicyDeletePolicyResponse Response from deleting a Policy.
//
// swagger:model workspace.policy.DeletePolicyResponse
type WorkspacePolicyDeletePolicyResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this workspace policy delete policy response
func (m *WorkspacePolicyDeletePolicyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workspace policy delete policy response based on context it is used
func (m *WorkspacePolicyDeletePolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspacePolicyDeletePolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspacePolicyDeletePolicyResponse) UnmarshalBinary(b []byte) error {
	var res WorkspacePolicyDeletePolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
