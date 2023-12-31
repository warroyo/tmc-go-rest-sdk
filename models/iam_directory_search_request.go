// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IamDirectorySearchRequest Request to list users and groups.
//
// swagger:model iam.directory.SearchRequest
type IamDirectorySearchRequest struct {

	// Search query (for some implementation-specific algorithm). For example,
	// a query of "use" might match the user "user1@example.com" and the group
	// "Warehouse" which both contain it as a substring.
	//
	// This query would typically come from a user entering text into an
	// autocompleted UI element.
	Query string `json:"query,omitempty"`
}

// Validate validates this iam directory search request
func (m *IamDirectorySearchRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this iam directory search request based on context it is used
func (m *IamDirectorySearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamDirectorySearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamDirectorySearchRequest) UnmarshalBinary(b []byte) error {
	var res IamDirectorySearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
