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

// ManagementclusterProvisionerTanzukubernetesclusterControlPlane The cluster specific control plane configuration.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.ControlPlane
type ManagementclusterProvisionerTanzukubernetesclusterControlPlane struct {

	// The metadata of the control plane.
	Metadata *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata `json:"metadata,omitempty"`

	// The OS image of the control plane.
	OsImage *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage `json:"osImage,omitempty"`

	// The replicas of the control plane.
	Replicas int32 `json:"replicas,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster control plane
func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) validateOsImage(formats strfmt.Registry) error {
	if swag.IsZero(m.OsImage) { // not required
		return nil
	}

	if m.OsImage != nil {
		if err := m.OsImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osImage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster control plane based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOsImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) contextValidateOsImage(ctx context.Context, formats strfmt.Registry) error {

	if m.OsImage != nil {
		if err := m.OsImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osImage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterControlPlane) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
