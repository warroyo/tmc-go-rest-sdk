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

// ManagementclusterProvisionerClusterclassVariableSchema VariableSchema defines the schema of a variable.
//
// swagger:model managementcluster.provisioner.clusterclass.VariableSchema
type ManagementclusterProvisionerClusterclassVariableSchema struct {

	// Template values in OpenAPI V3 schema format.
	Template *K8sIoApimachineryPkgRuntimeRawExtension `json:"template,omitempty"`
}

// Validate validates this managementcluster provisioner clusterclass variable schema
func (m *ManagementclusterProvisionerClusterclassVariableSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerClusterclassVariableSchema) validateTemplate(formats strfmt.Registry) error {
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

// ContextValidate validate this managementcluster provisioner clusterclass variable schema based on the context it is used
func (m *ManagementclusterProvisionerClusterclassVariableSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerClusterclassVariableSchema) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ManagementclusterProvisionerClusterclassVariableSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerClusterclassVariableSchema) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerClusterclassVariableSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}