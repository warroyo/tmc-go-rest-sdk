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

// ClusterObjectObjectNamespace Namespace Object containing aggregated resources.
//
// swagger:model cluster.object.ObjectNamespace
type ClusterObjectObjectNamespace struct {

	// Namespace object.
	ResourceNamespace *K8sIoAPICoreV1Namespace `json:"resourceNamespace,omitempty"`

	// Type of the namespace.
	Type *ClusterObjectObjectNamespaceType `json:"type,omitempty"`
}

// Validate validates this cluster object object namespace
func (m *ClusterObjectObjectNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectObjectNamespace) validateResourceNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceNamespace) { // not required
		return nil
	}

	if m.ResourceNamespace != nil {
		if err := m.ResourceNamespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceNamespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceNamespace")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterObjectObjectNamespace) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster object object namespace based on the context it is used
func (m *ClusterObjectObjectNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectObjectNamespace) contextValidateResourceNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceNamespace != nil {
		if err := m.ResourceNamespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceNamespace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceNamespace")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterObjectObjectNamespace) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterObjectObjectNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterObjectObjectNamespace) UnmarshalBinary(b []byte) error {
	var res ClusterObjectObjectNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
