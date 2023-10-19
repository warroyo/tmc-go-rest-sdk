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

// ClusterNamespaceTanzupackageInstallSpec Spec of Package Install.
//
// swagger:model cluster.namespace.tanzupackage.install.Spec
type ClusterNamespaceTanzupackageInstallSpec struct {

	// Inline values to configure the Package Install.
	InlineValues interface{} `json:"inlineValues,omitempty"`

	// Reference to the Package which will be installed.
	PackageRef *ClusterNamespaceTanzupackageMetadataPackagePackageRef `json:"packageRef,omitempty"`

	// Role binding scope for service account which will be used by Package Install.
	RoleBindingScope *ClusterNamespaceTanzupackageInstallRoleBindingScope `json:"roleBindingScope,omitempty"`
}

// Validate validates this cluster namespace tanzupackage install spec
func (m *ClusterNamespaceTanzupackageInstallSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleBindingScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageInstallSpec) validatePackageRef(formats strfmt.Registry) error {
	if swag.IsZero(m.PackageRef) { // not required
		return nil
	}

	if m.PackageRef != nil {
		if err := m.PackageRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageRef")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNamespaceTanzupackageInstallSpec) validateRoleBindingScope(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleBindingScope) { // not required
		return nil
	}

	if m.RoleBindingScope != nil {
		if err := m.RoleBindingScope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleBindingScope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleBindingScope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster namespace tanzupackage install spec based on the context it is used
func (m *ClusterNamespaceTanzupackageInstallSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackageRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleBindingScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageInstallSpec) contextValidatePackageRef(ctx context.Context, formats strfmt.Registry) error {

	if m.PackageRef != nil {
		if err := m.PackageRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageRef")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNamespaceTanzupackageInstallSpec) contextValidateRoleBindingScope(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleBindingScope != nil {
		if err := m.RoleBindingScope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleBindingScope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleBindingScope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageInstallSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageInstallSpec) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageInstallSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}