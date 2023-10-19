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

// ClusterNamespaceTanzupackageInstallUpdateInstallResponse Response from updating an Install.
//
// swagger:model cluster.namespace.tanzupackage.install.UpdateInstallResponse
type ClusterNamespaceTanzupackageInstallUpdateInstallResponse struct {

	// Install updated.
	Install *ClusterNamespaceTanzupackageInstallInstall `json:"install,omitempty"`
}

// Validate validates this cluster namespace tanzupackage install update install response
func (m *ClusterNamespaceTanzupackageInstallUpdateInstallResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageInstallUpdateInstallResponse) validateInstall(formats strfmt.Registry) error {
	if swag.IsZero(m.Install) { // not required
		return nil
	}

	if m.Install != nil {
		if err := m.Install.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("install")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("install")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster namespace tanzupackage install update install response based on the context it is used
func (m *ClusterNamespaceTanzupackageInstallUpdateInstallResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceTanzupackageInstallUpdateInstallResponse) contextValidateInstall(ctx context.Context, formats strfmt.Registry) error {

	if m.Install != nil {
		if err := m.Install.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("install")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("install")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageInstallUpdateInstallResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageInstallUpdateInstallResponse) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageInstallUpdateInstallResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}