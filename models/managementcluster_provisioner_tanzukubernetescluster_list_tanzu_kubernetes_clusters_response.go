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

// ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse Response from listing TanzuKubernetesClusters.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.ListTanzuKubernetesClustersResponse
type ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse struct {

	// List of tanzukubernetesclusters.
	TanzuKubernetesClusters []*ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster `json:"tanzuKubernetesClusters"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster list tanzu kubernetes clusters response
func (m *ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTanzuKubernetesClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse) validateTanzuKubernetesClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.TanzuKubernetesClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.TanzuKubernetesClusters); i++ {
		if swag.IsZero(m.TanzuKubernetesClusters[i]) { // not required
			continue
		}

		if m.TanzuKubernetesClusters[i] != nil {
			if err := m.TanzuKubernetesClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tanzuKubernetesClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tanzuKubernetesClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster list tanzu kubernetes clusters response based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTanzuKubernetesClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse) contextValidateTanzuKubernetesClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TanzuKubernetesClusters); i++ {

		if m.TanzuKubernetesClusters[i] != nil {
			if err := m.TanzuKubernetesClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tanzuKubernetesClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tanzuKubernetesClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
