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

// K8sIoAPIAppsV1DeploymentCondition DeploymentCondition describes the state of a deployment at a certain point.
//
// swagger:model k8s.io.api.apps.v1.DeploymentCondition
type K8sIoAPIAppsV1DeploymentCondition struct {

	// Last time the condition transitioned from one status to another.
	LastTransitionTime *K8sIoApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// The last time this condition was updated.
	LastUpdateTime *K8sIoApimachineryPkgApisMetaV1Time `json:"lastUpdateTime,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`

	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status,omitempty"`

	// Type of deployment condition.
	Type string `json:"type,omitempty"`
}

// Validate validates this k8s io api apps v1 deployment condition
func (m *K8sIoAPIAppsV1DeploymentCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIAppsV1DeploymentCondition) validateLastTransitionTime(formats strfmt.Registry) error {
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

func (m *K8sIoAPIAppsV1DeploymentCondition) validateLastUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdateTime) { // not required
		return nil
	}

	if m.LastUpdateTime != nil {
		if err := m.LastUpdateTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdateTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastUpdateTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api apps v1 deployment condition based on the context it is used
func (m *K8sIoAPIAppsV1DeploymentCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIAppsV1DeploymentCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

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

func (m *K8sIoAPIAppsV1DeploymentCondition) contextValidateLastUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastUpdateTime != nil {
		if err := m.LastUpdateTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdateTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastUpdateTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPIAppsV1DeploymentCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPIAppsV1DeploymentCondition) UnmarshalBinary(b []byte) error {
	var res K8sIoAPIAppsV1DeploymentCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
