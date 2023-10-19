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

// ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse Streamed response from watching Nodepools.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.nodepool.WatchNodepoolsResponse
type ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse struct {

	// Type of event.
	EventType *VmwareTanzuCoreV1alpha1EventType `json:"eventType,omitempty"`

	// Nodepool event.
	Nodepool *ManagementclusterProvisionerTanzukubernetesclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster nodepool watch nodepools response
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodepool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) validateEventType(formats strfmt.Registry) error {
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

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) validateNodepool(formats strfmt.Registry) error {
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

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster nodepool watch nodepools response based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodepool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) contextValidateNodepool(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterNodepoolWatchNodepoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
