// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterDataprotectionBackupDeleteBackupResponse Response from deleting a Backup.
//
// swagger:model cluster.dataprotection.backup.DeleteBackupResponse
type ClusterDataprotectionBackupDeleteBackupResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this cluster dataprotection backup delete backup response
func (m *ClusterDataprotectionBackupDeleteBackupResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster dataprotection backup delete backup response based on context it is used
func (m *ClusterDataprotectionBackupDeleteBackupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionBackupDeleteBackupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionBackupDeleteBackupResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionBackupDeleteBackupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}