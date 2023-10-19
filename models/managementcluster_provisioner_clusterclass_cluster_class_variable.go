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

// ManagementclusterProvisionerClusterclassClusterClassVariable ClusterClassVariable defines a variable which can be configured in the Cluster topology.
//
// swagger:model managementcluster.provisioner.clusterclass.ClusterClassVariable
type ManagementclusterProvisionerClusterclassClusterClassVariable struct {

	// Name of the cluster class variable.
	Name string `json:"name,omitempty"`

	// Required specifies if the variable is required.
	Required bool `json:"required,omitempty"`

	// Schema defines the schema of the variable.
	Schema *ManagementclusterProvisionerClusterclassVariableSchema `json:"schema,omitempty"`
}

// Validate validates this managementcluster provisioner clusterclass cluster class variable
func (m *ManagementclusterProvisionerClusterclassClusterClassVariable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerClusterclassClusterClassVariable) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster provisioner clusterclass cluster class variable based on the context it is used
func (m *ManagementclusterProvisionerClusterclassClusterClassVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerClusterclassClusterClassVariable) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerClusterclassClusterClassVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerClusterclassClusterClassVariable) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerClusterclassClusterClassVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
