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

// ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse Streamed response from enumerating TanzuKubernetesClusters.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.EnumerateTanzuKubernetesClustersResponse
type ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse struct {

	// TanzuKubernetesCluster object.
	TanzuKubernetesCluster *ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster `json:"tanzuKubernetesCluster,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster enumerate tanzu kubernetes clusters response
func (m *ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTanzuKubernetesCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse) validateTanzuKubernetesCluster(formats strfmt.Registry) error {
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

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster enumerate tanzu kubernetes clusters response based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTanzuKubernetesCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse) contextValidateTanzuKubernetesCluster(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterEnumerateTanzuKubernetesClustersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
