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

// K8sIoAPICoreV1ContainerState ContainerState holds a possible state of container.
// Only one of its members may be specified.
// If none of them is specified, the default one is ContainerStateWaiting.
//
// swagger:model k8s.io.api.core.v1.ContainerState
type K8sIoAPICoreV1ContainerState struct {

	// Details about a running container
	// +optional
	Running *K8sIoAPICoreV1ContainerStateRunning `json:"running,omitempty"`

	// Details about a terminated container
	// +optional
	Terminated *K8sIoAPICoreV1ContainerStateTerminated `json:"terminated,omitempty"`

	// Details about a waiting container
	// +optional
	Waiting *K8sIoAPICoreV1ContainerStateWaiting `json:"waiting,omitempty"`
}

// Validate validates this k8s io api core v1 container state
func (m *K8sIoAPICoreV1ContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWaiting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ContainerState) validateRunning(formats strfmt.Registry) error {
	if swag.IsZero(m.Running) { // not required
		return nil
	}

	if m.Running != nil {
		if err := m.Running.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1ContainerState) validateTerminated(formats strfmt.Registry) error {
	if swag.IsZero(m.Terminated) { // not required
		return nil
	}

	if m.Terminated != nil {
		if err := m.Terminated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terminated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terminated")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1ContainerState) validateWaiting(formats strfmt.Registry) error {
	if swag.IsZero(m.Waiting) { // not required
		return nil
	}

	if m.Waiting != nil {
		if err := m.Waiting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waiting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("waiting")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 container state based on the context it is used
func (m *K8sIoAPICoreV1ContainerState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRunning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerminated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWaiting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ContainerState) contextValidateRunning(ctx context.Context, formats strfmt.Registry) error {

	if m.Running != nil {
		if err := m.Running.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1ContainerState) contextValidateTerminated(ctx context.Context, formats strfmt.Registry) error {

	if m.Terminated != nil {
		if err := m.Terminated.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terminated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terminated")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1ContainerState) contextValidateWaiting(ctx context.Context, formats strfmt.Registry) error {

	if m.Waiting != nil {
		if err := m.Waiting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waiting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("waiting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ContainerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ContainerState) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ContainerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
