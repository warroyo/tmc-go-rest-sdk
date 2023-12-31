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

// ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse Response from applying a TanzuKubernetesCluster.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.ApplyTanzuKubernetesClusterResponse
type ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse struct {

	// TanzuKubernetesCluster applied.
	TanzuKubernetesCluster *ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster `json:"tanzuKubernetesCluster,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster apply tanzu kubernetes cluster response
func (m *ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTanzuKubernetesCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse) validateTanzuKubernetesCluster(formats strfmt.Registry) error {
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

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster apply tanzu kubernetes cluster response based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTanzuKubernetesCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse) contextValidateTanzuKubernetesCluster(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterApplyTanzuKubernetesClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
