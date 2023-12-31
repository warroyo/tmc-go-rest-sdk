// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AksclusterStorageConfig The storage config.
//
// swagger:model akscluster.StorageConfig
type AksclusterStorageConfig struct {

	// Enable the azure disk CSI driver for the storage.
	EnableDiskCsiDriver bool `json:"enableDiskCsiDriver,omitempty"`

	// Enable the azure file CSI driver for the storage.
	EnableFileCsiDriver bool `json:"enableFileCsiDriver,omitempty"`

	// Enable the snapshot controller for the storage.
	EnableSnapshotController bool `json:"enableSnapshotController,omitempty"`
}

// Validate validates this akscluster storage config
func (m *AksclusterStorageConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this akscluster storage config based on context it is used
func (m *AksclusterStorageConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterStorageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterStorageConfig) UnmarshalBinary(b []byte) error {
	var res AksclusterStorageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
