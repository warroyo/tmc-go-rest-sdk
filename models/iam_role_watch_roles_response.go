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

// IamRoleWatchRolesResponse Streamed response from watching Roles.
//
// swagger:model iam.role.WatchRolesResponse
type IamRoleWatchRolesResponse struct {

	// Type of event.
	EventType *VmwareTanzuCoreV1alpha1EventType `json:"eventType,omitempty"`

	// Role event.
	Role *IamRoleRole `json:"role,omitempty"`
}

// Validate validates this iam role watch roles response
func (m *IamRoleWatchRolesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamRoleWatchRolesResponse) validateEventType(formats strfmt.Registry) error {
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

func (m *IamRoleWatchRolesResponse) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iam role watch roles response based on the context it is used
func (m *IamRoleWatchRolesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamRoleWatchRolesResponse) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IamRoleWatchRolesResponse) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamRoleWatchRolesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamRoleWatchRolesResponse) UnmarshalBinary(b []byte) error {
	var res IamRoleWatchRolesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
