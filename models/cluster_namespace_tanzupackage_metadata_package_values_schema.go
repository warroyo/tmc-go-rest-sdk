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

// ClusterNamespaceTanzupackageMetadataPackageValuesSchema ValuesSchema is used to show template values that can be configured by users while installing Package.
//
// swagger:model cluster.namespace.tanzupackage.metadata.package.ValuesSchema
type ClusterNamespaceTanzupackageMetadataPackageValuesSchema struct {

	// Template values in OpenAPI V3 schema format.
	Template *K8sIoApimachineryPkgRuntimeRawExtension `json:"template,omitempty"`
}

// Validate validates this cluster namespace tanzupackage metadata package values schema
func (m *ClusterNamespaceTanzupackageMetadataPackageValuesSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageMetadataPackageValuesSchema) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster namespace tanzupackage metadata package values schema based on the context it is used
func (m *ClusterNamespaceTanzupackageMetadataPackageValuesSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageMetadataPackageValuesSchema) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageMetadataPackageValuesSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageMetadataPackageValuesSchema) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageMetadataPackageValuesSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}