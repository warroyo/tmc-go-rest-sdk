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

// K8sIoAPICoreV1ReplicationControllerCondition ReplicationControllerCondition describes the state of a replication controller at a certain point.
//
// swagger:model k8s.io.api.core.v1.ReplicationControllerCondition
type K8sIoAPICoreV1ReplicationControllerCondition struct {

	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *K8sIoApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty"`

	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status,omitempty"`

	// Type of replication controller condition.
	Type string `json:"type,omitempty"`
}

// Validate validates this k8s io api core v1 replication controller condition
func (m *K8sIoAPICoreV1ReplicationControllerCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ReplicationControllerCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if m.LastTransitionTime != nil {
		if err := m.LastTransitionTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTransitionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastTransitionTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 replication controller condition based on the context it is used
func (m *K8sIoAPICoreV1ReplicationControllerCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ReplicationControllerCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastTransitionTime != nil {
		if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTransitionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastTransitionTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ReplicationControllerCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ReplicationControllerCondition) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ReplicationControllerCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
