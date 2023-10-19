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

// K8sIoAPIAppsV1StatefulSetUpdateStrategy StatefulSetUpdateStrategy indicates the strategy that the StatefulSet
// controller will use to perform updates. It includes any additional parameters
// necessary to perform the update for the indicated strategy.
//
// swagger:model k8s.io.api.apps.v1.StatefulSetUpdateStrategy
type K8sIoAPIAppsV1StatefulSetUpdateStrategy struct {

	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	// +optional
	RollingUpdate *K8sIoAPIAppsV1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`

	// Type indicates the type of the StatefulSetUpdateStrategy.
	// Default is RollingUpdate.
	// +optional
	Type string `json:"type,omitempty"`
}

// Validate validates this k8s io api apps v1 stateful set update strategy
func (m *K8sIoAPIAppsV1StatefulSetUpdateStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRollingUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIAppsV1StatefulSetUpdateStrategy) validateRollingUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.RollingUpdate) { // not required
		return nil
	}

	if m.RollingUpdate != nil {
		if err := m.RollingUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollingUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollingUpdate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api apps v1 stateful set update strategy based on the context it is used
func (m *K8sIoAPIAppsV1StatefulSetUpdateStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRollingUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIAppsV1StatefulSetUpdateStrategy) contextValidateRollingUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.RollingUpdate != nil {
		if err := m.RollingUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollingUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollingUpdate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPIAppsV1StatefulSetUpdateStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPIAppsV1StatefulSetUpdateStrategy) UnmarshalBinary(b []byte) error {
	var res K8sIoAPIAppsV1StatefulSetUpdateStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
