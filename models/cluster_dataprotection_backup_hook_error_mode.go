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

// ClusterDataprotectionBackupHookErrorMode Specifies how Velero should behave if it encounters an error executing this hook.
//
//   - MODE_UNSPECIFIED: The default mode.
//   - CONTINUE: Means that an error from a hook is acceptable, and the operation can proceed.
//   - FAIL: Means that an error from a hook is problematic, and the operation should be in error.
//
// swagger:model cluster.dataprotection.backup.HookErrorMode
type ClusterDataprotectionBackupHookErrorMode string

func NewClusterDataprotectionBackupHookErrorMode(value ClusterDataprotectionBackupHookErrorMode) *ClusterDataprotectionBackupHookErrorMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterDataprotectionBackupHookErrorMode.
func (m ClusterDataprotectionBackupHookErrorMode) Pointer() *ClusterDataprotectionBackupHookErrorMode {
	return &m
}

const (

	// ClusterDataprotectionBackupHookErrorModeMODEUNSPECIFIED captures enum value "MODE_UNSPECIFIED"
	ClusterDataprotectionBackupHookErrorModeMODEUNSPECIFIED ClusterDataprotectionBackupHookErrorMode = "MODE_UNSPECIFIED"

	// ClusterDataprotectionBackupHookErrorModeCONTINUE captures enum value "CONTINUE"
	ClusterDataprotectionBackupHookErrorModeCONTINUE ClusterDataprotectionBackupHookErrorMode = "CONTINUE"

	// ClusterDataprotectionBackupHookErrorModeFAIL captures enum value "FAIL"
	ClusterDataprotectionBackupHookErrorModeFAIL ClusterDataprotectionBackupHookErrorMode = "FAIL"
)

// for schema
var clusterDataprotectionBackupHookErrorModeEnum []interface{}

func init() {
	var res []ClusterDataprotectionBackupHookErrorMode
	if err := json.Unmarshal([]byte(`["MODE_UNSPECIFIED","CONTINUE","FAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterDataprotectionBackupHookErrorModeEnum = append(clusterDataprotectionBackupHookErrorModeEnum, v)
	}
}

func (m ClusterDataprotectionBackupHookErrorMode) validateClusterDataprotectionBackupHookErrorModeEnum(path, location string, value ClusterDataprotectionBackupHookErrorMode) error {
	if err := validate.EnumCase(path, location, value, clusterDataprotectionBackupHookErrorModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster dataprotection backup hook error mode
func (m ClusterDataprotectionBackupHookErrorMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterDataprotectionBackupHookErrorModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster dataprotection backup hook error mode based on context it is used
func (m ClusterDataprotectionBackupHookErrorMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
