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

// ClustergroupWatchClusterGroupsResponse Streamed response from watching ClusterGroups.
//
// swagger:model clustergroup.WatchClusterGroupsResponse
type ClustergroupWatchClusterGroupsResponse struct {

	// ClusterGroup event.
	ClusterGroup *ClustergroupClusterGroup `json:"clusterGroup,omitempty"`

	// Type of event.
	EventType *VmwareTanzuCoreV1alpha1EventType `json:"eventType,omitempty"`
}

// Validate validates this clustergroup watch cluster groups response
func (m *ClustergroupWatchClusterGroupsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupWatchClusterGroupsResponse) validateClusterGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterGroup) { // not required
		return nil
	}

	if m.ClusterGroup != nil {
		if err := m.ClusterGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ClustergroupWatchClusterGroupsResponse) validateEventType(formats strfmt.Registry) error {
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

// ContextValidate validate this clustergroup watch cluster groups response based on the context it is used
func (m *ClustergroupWatchClusterGroupsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupWatchClusterGroupsResponse) contextValidateClusterGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterGroup != nil {
		if err := m.ClusterGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ClustergroupWatchClusterGroupsResponse) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ClustergroupWatchClusterGroupsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupWatchClusterGroupsResponse) UnmarshalBinary(b []byte) error {
	var res ClustergroupWatchClusterGroupsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
