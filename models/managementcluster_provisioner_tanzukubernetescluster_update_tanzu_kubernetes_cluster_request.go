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

// ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest Request to update (overwrite) a TanzuKubernetesCluster.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.UpdateTanzuKubernetesClusterRequest
type ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest struct {

	// Update TanzuKubernetesCluster.
	TanzuKubernetesCluster *ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster `json:"tanzuKubernetesCluster,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster update tanzu kubernetes cluster request
func (m *ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTanzuKubernetesCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest) validateTanzuKubernetesCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.TanzuKubernetesCluster) { // not required
		return nil
	}

	if m.TanzuKubernetesCluster != nil {
		if err := m.TanzuKubernetesCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tanzuKubernetesCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tanzuKubernetesCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster update tanzu kubernetes cluster request based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTanzuKubernetesCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest) contextValidateTanzuKubernetesCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.TanzuKubernetesCluster != nil {
		if err := m.TanzuKubernetesCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tanzuKubernetesCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tanzuKubernetesCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterUpdateTanzuKubernetesClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
