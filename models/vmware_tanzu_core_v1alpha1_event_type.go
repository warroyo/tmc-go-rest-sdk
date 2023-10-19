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

// VmwareTanzuCoreV1alpha1EventType Possible types of events.
//
//   - TYPE_UNSPECIFIED: Unspecified event type.
//   - CREATE: Create event.
//   - UPDATE: Update event.
//   - DELETE: Delete event.
//
// swagger:model vmware.tanzu.core.v1alpha1.event.Type
type VmwareTanzuCoreV1alpha1EventType string

func NewVmwareTanzuCoreV1alpha1EventType(value VmwareTanzuCoreV1alpha1EventType) *VmwareTanzuCoreV1alpha1EventType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuCoreV1alpha1EventType.
func (m VmwareTanzuCoreV1alpha1EventType) Pointer() *VmwareTanzuCoreV1alpha1EventType {
	return &m
}

const (

	// VmwareTanzuCoreV1alpha1EventTypeTYPEUNSPECIFIED captures enum value "TYPE_UNSPECIFIED"
	VmwareTanzuCoreV1alpha1EventTypeTYPEUNSPECIFIED VmwareTanzuCoreV1alpha1EventType = "TYPE_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1EventTypeCREATE captures enum value "CREATE"
	VmwareTanzuCoreV1alpha1EventTypeCREATE VmwareTanzuCoreV1alpha1EventType = "CREATE"

	// VmwareTanzuCoreV1alpha1EventTypeUPDATE captures enum value "UPDATE"
	VmwareTanzuCoreV1alpha1EventTypeUPDATE VmwareTanzuCoreV1alpha1EventType = "UPDATE"

	// VmwareTanzuCoreV1alpha1EventTypeDELETE captures enum value "DELETE"
	VmwareTanzuCoreV1alpha1EventTypeDELETE VmwareTanzuCoreV1alpha1EventType = "DELETE"
)

// for schema
var vmwareTanzuCoreV1alpha1EventTypeEnum []interface{}

func init() {
	var res []VmwareTanzuCoreV1alpha1EventType
	if err := json.Unmarshal([]byte(`["TYPE_UNSPECIFIED","CREATE","UPDATE","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareTanzuCoreV1alpha1EventTypeEnum = append(vmwareTanzuCoreV1alpha1EventTypeEnum, v)
	}
}

func (m VmwareTanzuCoreV1alpha1EventType) validateVmwareTanzuCoreV1alpha1EventTypeEnum(path, location string, value VmwareTanzuCoreV1alpha1EventType) error {
	if err := validate.EnumCase(path, location, value, vmwareTanzuCoreV1alpha1EventTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vmware tanzu core v1alpha1 event type
func (m VmwareTanzuCoreV1alpha1EventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVmwareTanzuCoreV1alpha1EventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vmware tanzu core v1alpha1 event type based on context it is used
func (m VmwareTanzuCoreV1alpha1EventType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
