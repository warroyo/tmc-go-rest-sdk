// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1ConfigMapEnvSource ConfigMapEnvSource selects a ConfigMap to populate the environment
// variables with.
//
// The contents of the target ConfigMap's Data field will represent the
// key-value pairs as environment variables.
//
// swagger:model k8s.io.api.core.v1.ConfigMapEnvSource
type K8sIoAPICoreV1ConfigMapEnvSource struct {

	// The ConfigMap to select from.
	LocalObjectReference *K8sIoAPICoreV1LocalObjectReference `json:"localObjectReference,omitempty"`

	// Specify whether the ConfigMap must be defined
	// +optional
	Optional bool `json:"optional,omitempty"`
}

// Validate validates this k8s io api core v1 config map env source
func (m *K8sIoAPICoreV1ConfigMapEnvSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalObjectReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ConfigMapEnvSource) validateLocalObjectReference(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalObjectReference) { // not required
		return nil
	}

	if m.LocalObjectReference != nil {
		if err := m.LocalObjectReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localObjectReference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localObjectReference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 config map env source based on the context it is used
func (m *K8sIoAPICoreV1ConfigMapEnvSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocalObjectReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ConfigMapEnvSource) contextValidateLocalObjectReference(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalObjectReference != nil {
		if err := m.LocalObjectReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localObjectReference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localObjectReference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ConfigMapEnvSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ConfigMapEnvSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ConfigMapEnvSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
