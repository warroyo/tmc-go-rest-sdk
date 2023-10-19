// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustergroupPolicyDeletePolicyResponse Response from deleting a Policy.
//
// swagger:model clustergroup.policy.DeletePolicyResponse
type ClustergroupPolicyDeletePolicyResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this clustergroup policy delete policy response
func (m *ClustergroupPolicyDeletePolicyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clustergroup policy delete policy response based on context it is used
func (m *ClustergroupPolicyDeletePolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupPolicyDeletePolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupPolicyDeletePolicyResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupPolicyDeletePolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
