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

// ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec Spec for the cluster nodepool.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.nodepool.Spec
type ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec struct {

	// The name of the machine deployment class used to create the nodepool.
	Class string `json:"class,omitempty"`

	// The failure domain the machines will be created in.
	FailureDomain string `json:"failureDomain,omitempty"`

	// The metadata of the nodepool.
	Metadata *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata `json:"metadata,omitempty"`

	// The OS image of the nodepool.
	OsImage *ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage `json:"osImage,omitempty"`

	// Overrides can be used to override cluster level variables.
	Overrides []*ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable `json:"overrides"`

	// The replicas of the nodepool.
	Replicas int32 `json:"replicas,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster nodepool spec
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) validateMetadata(formats strfmt.Registry) error {
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

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) validateOsImage(formats strfmt.Registry) error {
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

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) validateOverrides(formats strfmt.Registry) error {
	if swag.IsZero(m.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(m.Overrides); i++ {
		if swag.IsZero(m.Overrides[i]) { // not required
			continue
		}

		if m.Overrides[i] != nil {
			if err := m.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster nodepool spec based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOsImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) contextValidateOsImage(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) contextValidateOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Overrides); i++ {

		if m.Overrides[i] != nil {
			if err := m.Overrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
