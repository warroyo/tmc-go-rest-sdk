// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1ConfigMapProjection Adapts a ConfigMap into a projected volume.
//
// The contents of the target ConfigMap's Data field will be presented in a
// projected volume as files using the keys in the Data field as the file names,
// unless the items element is populated with specific mappings of keys to paths.
// Note that this is identical to a configmap volume source without the default
// mode.
//
// swagger:model k8s.io.api.core.v1.ConfigMapProjection
type K8sIoAPICoreV1ConfigMapProjection struct {

	// If unspecified, each key-value pair in the Data field of the referenced
	// ConfigMap will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the ConfigMap,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	Items []*K8sIoAPICoreV1KeyToPath `json:"items"`

	// local object reference
	LocalObjectReference *K8sIoAPICoreV1LocalObjectReference `json:"localObjectReference,omitempty"`

	// Specify whether the ConfigMap or it's keys must be defined
	// +optional
	Optional bool `json:"optional,omitempty"`
}

// Validate validates this k8s io api core v1 config map projection
func (m *K8sIoAPICoreV1ConfigMapProjection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalObjectReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ConfigMapProjection) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1ConfigMapProjection) validateLocalObjectReference(formats strfmt.Registry) error {
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

// ContextValidate validate this k8s io api core v1 config map projection based on the context it is used
func (m *K8sIoAPICoreV1ConfigMapProjection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalObjectReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ConfigMapProjection) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1ConfigMapProjection) contextValidateLocalObjectReference(ctx context.Context, formats strfmt.Registry) error {

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
func (m *K8sIoAPICoreV1ConfigMapProjection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ConfigMapProjection) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ConfigMapProjection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
