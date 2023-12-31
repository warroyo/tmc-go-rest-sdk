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

// EksclusterNodepoolWatchNodepoolsResponse Streamed response from watching Nodepools.
//
// swagger:model ekscluster.nodepool.WatchNodepoolsResponse
type EksclusterNodepoolWatchNodepoolsResponse struct {

	// Type of event.
	EventType *VmwareTanzuCoreV1alpha1EventType `json:"eventType,omitempty"`

	// Nodepool event.
	Nodepool *EksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// Validate validates this ekscluster nodepool watch nodepools response
func (m *EksclusterNodepoolWatchNodepoolsResponse) Validate(formats strfmt.Registry) error {
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

func (m *EksclusterNodepoolWatchNodepoolsResponse) validateEventType(formats strfmt.Registry) error {
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

func (m *EksclusterNodepoolWatchNodepoolsResponse) validateNodepool(formats strfmt.Registry) error {
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

// ContextValidate validate this ekscluster nodepool watch nodepools response based on the context it is used
func (m *EksclusterNodepoolWatchNodepoolsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *EksclusterNodepoolWatchNodepoolsResponse) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EksclusterNodepoolWatchNodepoolsResponse) contextValidateNodepool(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EksclusterNodepoolWatchNodepoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterNodepoolWatchNodepoolsResponse) UnmarshalBinary(b []byte) error {
	var res EksclusterNodepoolWatchNodepoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
