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

// ClusterObjectStatus The Object status.
//
// swagger:model cluster.object.Status
type ClusterObjectStatus struct {

	// Health of the object.
	Health *ClusterObjectResourceHealth `json:"health,omitempty"`

	// Phase of the resource.
	Phase *ClusterObjectStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`
}

// Validate validates this cluster object status
func (m *ClusterObjectStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectStatus) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterObjectStatus) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster object status based on the context it is used
func (m *ClusterObjectStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectStatus) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {
		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterObjectStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterObjectStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterObjectStatus) UnmarshalBinary(b []byte) error {
	var res ClusterObjectStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}