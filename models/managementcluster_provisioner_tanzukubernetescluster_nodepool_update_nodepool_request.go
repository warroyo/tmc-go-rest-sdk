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

// ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest Request to update (overwrite) a Nodepool.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.nodepool.UpdateNodepoolRequest
type ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest struct {

	// Update Nodepool.
	Nodepool *ManagementclusterProvisionerTanzukubernetesclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster nodepool update nodepool request
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodepool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest) validateNodepool(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodepool) { // not required
		return nil
	}

	if m.Nodepool != nil {
		if err := m.Nodepool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodepool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodepool")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster nodepool update nodepool request based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodepool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest) contextValidateNodepool(ctx context.Context, formats strfmt.Registry) error {

	if m.Nodepool != nil {
		if err := m.Nodepool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodepool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodepool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterNodepoolUpdateNodepoolRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
