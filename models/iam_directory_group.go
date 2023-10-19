// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IamDirectoryGroup Group represents a named group from the upstream identity provider.
//
// swagger:model iam.directory.Group
type IamDirectoryGroup struct {

	// Unique group ID (e.g., "be9a51ac-e154-4a45-b0b8-edd80a658c0b")
	ID string `json:"id,omitempty"`

	// Human readable display name (e.g., "my group")
	Name string `json:"name,omitempty"`

	// Number of users who belong to this group.
	UserCount int64 `json:"userCount,omitempty"`
}

// Validate validates this iam directory group
func (m *IamDirectoryGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this iam directory group based on context it is used
func (m *IamDirectoryGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamDirectoryGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamDirectoryGroup) UnmarshalBinary(b []byte) error {
	var res IamDirectoryGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}