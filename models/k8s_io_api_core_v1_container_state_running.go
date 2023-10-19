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

// K8sIoAPICoreV1ContainerStateRunning ContainerStateRunning is a running state of a container.
//
// swagger:model k8s.io.api.core.v1.ContainerStateRunning
type K8sIoAPICoreV1ContainerStateRunning struct {

	// Time at which the container was last (re-)started
	// +optional
	StartedAt *K8sIoApimachineryPkgApisMetaV1Time `json:"startedAt,omitempty"`
}

// Validate validates this k8s io api core v1 container state running
func (m *K8sIoAPICoreV1ContainerStateRunning) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ContainerStateRunning) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if m.StartedAt != nil {
		if err := m.StartedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startedAt")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 container state running based on the context it is used
func (m *K8sIoAPICoreV1ContainerStateRunning) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStartedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ContainerStateRunning) contextValidateStartedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.StartedAt != nil {
		if err := m.StartedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startedAt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ContainerStateRunning) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ContainerStateRunning) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ContainerStateRunning
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}