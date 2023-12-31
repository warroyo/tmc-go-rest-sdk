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

// K8sIoAPICoreV1NodeConfigSource NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil.
//
// swagger:model k8s.io.api.core.v1.NodeConfigSource
type K8sIoAPICoreV1NodeConfigSource struct {

	// ConfigMap is a reference to a Node's ConfigMap
	ConfigMap *K8sIoAPICoreV1ConfigMapNodeConfigSource `json:"configMap,omitempty"`
}

// Validate validates this k8s io api core v1 node config source
func (m *K8sIoAPICoreV1NodeConfigSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1NodeConfigSource) validateConfigMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigMap) { // not required
		return nil
	}

	if m.ConfigMap != nil {
		if err := m.ConfigMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configMap")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 node config source based on the context it is used
func (m *K8sIoAPICoreV1NodeConfigSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1NodeConfigSource) contextValidateConfigMap(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigMap != nil {
		if err := m.ConfigMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configMap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1NodeConfigSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1NodeConfigSource) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1NodeConfigSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
