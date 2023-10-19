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

// ManagementclusterProvisionerWatchprovisionersResponse Streamed response from watching provisioners.
//
// swagger:model managementcluster.provisioner.WatchprovisionersResponse
type ManagementclusterProvisionerWatchprovisionersResponse struct {

	// Type of event.
	EventType *VmwareTanzuCoreV1alpha1EventType `json:"eventType,omitempty"`

	// Provisioner event.
	Provisioner *ManagementclusterProvisionerProvisioner `json:"provisioner,omitempty"`
}

// Validate validates this managementcluster provisioner watchprovisioners response
func (m *ManagementclusterProvisionerWatchprovisionersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisioner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerWatchprovisionersResponse) validateEventType(formats strfmt.Registry) error {
	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	if m.EventType != nil {
		if err := m.EventType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventType")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerWatchprovisionersResponse) validateProvisioner(formats strfmt.Registry) error {
	if swag.IsZero(m.Provisioner) { // not required
		return nil
	}

	if m.Provisioner != nil {
		if err := m.Provisioner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisioner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provisioner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster provisioner watchprovisioners response based on the context it is used
func (m *ManagementclusterProvisionerWatchprovisionersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvisioner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerWatchprovisionersResponse) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

	if m.EventType != nil {
		if err := m.EventType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventType")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerWatchprovisionersResponse) contextValidateProvisioner(ctx context.Context, formats strfmt.Registry) error {

	if m.Provisioner != nil {
		if err := m.Provisioner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisioner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provisioner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerWatchprovisionersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerWatchprovisionersResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerWatchprovisionersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
