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

// K8sIoAPICoreV1PodCondition PodCondition contains details for the current condition of this pod.
//
// swagger:model k8s.io.api.core.v1.PodCondition
type K8sIoAPICoreV1PodCondition struct {

	// Last time we probed the condition.
	// +optional
	LastProbeTime *K8sIoApimachineryPkgApisMetaV1Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *K8sIoApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition.
	// Can be True, False, Unknown.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Status string `json:"status,omitempty"`

	// Type is the type of the condition.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Type string `json:"type,omitempty"`
}

// Validate validates this k8s io api core v1 pod condition
func (m *K8sIoAPICoreV1PodCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastProbeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1PodCondition) validateLastProbeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastProbeTime) { // not required
		return nil
	}

	if m.LastProbeTime != nil {
		if err := m.LastProbeTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastProbeTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastProbeTime")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1PodCondition) validateLastTransitionTime(formats strfmt.Registry) error {
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

// ContextValidate validate this k8s io api core v1 pod condition based on the context it is used
func (m *K8sIoAPICoreV1PodCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastProbeTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1PodCondition) contextValidateLastProbeTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastProbeTime != nil {
		if err := m.LastProbeTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastProbeTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastProbeTime")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1PodCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

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
func (m *K8sIoAPICoreV1PodCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1PodCondition) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1PodCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
