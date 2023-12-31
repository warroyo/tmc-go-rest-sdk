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

// AksclusterNodepoolOsType The operation system type options of the nodepool.
//
//   - OS_TYPE_UNSPECIFIED: Unspecified OS type.
//   - LINUX: The linux operation system.
//
// swagger:model akscluster.nodepool.OsType
type AksclusterNodepoolOsType string

func NewAksclusterNodepoolOsType(value AksclusterNodepoolOsType) *AksclusterNodepoolOsType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AksclusterNodepoolOsType.
func (m AksclusterNodepoolOsType) Pointer() *AksclusterNodepoolOsType {
	return &m
}

const (

	// AksclusterNodepoolOsTypeOSTYPEUNSPECIFIED captures enum value "OS_TYPE_UNSPECIFIED"
	AksclusterNodepoolOsTypeOSTYPEUNSPECIFIED AksclusterNodepoolOsType = "OS_TYPE_UNSPECIFIED"

	// AksclusterNodepoolOsTypeLINUX captures enum value "LINUX"
	AksclusterNodepoolOsTypeLINUX AksclusterNodepoolOsType = "LINUX"
)

// for schema
var aksclusterNodepoolOsTypeEnum []interface{}

func init() {
	var res []AksclusterNodepoolOsType
	if err := json.Unmarshal([]byte(`["OS_TYPE_UNSPECIFIED","LINUX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aksclusterNodepoolOsTypeEnum = append(aksclusterNodepoolOsTypeEnum, v)
	}
}

func (m AksclusterNodepoolOsType) validateAksclusterNodepoolOsTypeEnum(path, location string, value AksclusterNodepoolOsType) error {
	if err := validate.EnumCase(path, location, value, aksclusterNodepoolOsTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this akscluster nodepool os type
func (m AksclusterNodepoolOsType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAksclusterNodepoolOsTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this akscluster nodepool os type based on context it is used
func (m AksclusterNodepoolOsType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
