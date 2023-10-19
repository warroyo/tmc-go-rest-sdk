// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementclusterProvisionerDeleteProvisionerResponse Response from deleting a Provisioner.
//
// swagger:model managementcluster.provisioner.DeleteProvisionerResponse
type ManagementclusterProvisionerDeleteProvisionerResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// Validate validates this managementcluster provisioner delete provisioner response
func (m *ManagementclusterProvisionerDeleteProvisionerResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this managementcluster provisioner delete provisioner response based on context it is used
func (m *ManagementclusterProvisionerDeleteProvisionerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerDeleteProvisionerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerDeleteProvisionerResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerDeleteProvisionerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
