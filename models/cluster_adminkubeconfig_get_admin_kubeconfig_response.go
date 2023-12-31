// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterAdminkubeconfigGetAdminKubeconfigResponse Response with cluster admin kubeconfig.
//
// swagger:model cluster.adminkubeconfig.GetAdminKubeconfigResponse
type ClusterAdminkubeconfigGetAdminKubeconfigResponse struct {

	// Cluster Admin Kubeconfig.
	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// Validate validates this cluster adminkubeconfig get admin kubeconfig response
func (m *ClusterAdminkubeconfigGetAdminKubeconfigResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster adminkubeconfig get admin kubeconfig response based on context it is used
func (m *ClusterAdminkubeconfigGetAdminKubeconfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterAdminkubeconfigGetAdminKubeconfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterAdminkubeconfigGetAdminKubeconfigResponse) UnmarshalBinary(b []byte) error {
	var res ClusterAdminkubeconfigGetAdminKubeconfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
