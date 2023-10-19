// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode The permissions for a BackupStorageLocation.
//
//   - READONLY: The read only access.
//   - READWRITE: Read and write access.
//
// swagger:model cluster.dataprotection.backuplocation.Status.BackupStorageLocationAccessMode
type ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode string

func NewClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode(value ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) *ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode.
func (m ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) Pointer() *ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode {
	return &m
}

const (

	// ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADONLY captures enum value "READONLY"
	ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADONLY ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode = "READONLY"

	// ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADWRITE captures enum value "READWRITE"
	ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADWRITE ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode = "READWRITE"
)

// for schema
var clusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum []interface{}

func init() {
	var res []ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode
	if err := json.Unmarshal([]byte(`["READONLY","READWRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum = append(clusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum, v)
	}
}

func (m ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) validateClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum(path, location string, value ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) error {
	if err := validate.EnumCase(path, location, value, clusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster dataprotection backuplocation status backup storage location access mode
func (m ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster dataprotection backuplocation status backup storage location access mode based on context it is used
func (m ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}