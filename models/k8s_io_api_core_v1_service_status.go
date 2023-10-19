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

// K8sIoAPICoreV1ServiceStatus ServiceStatus represents the current status of a service.
//
// swagger:model k8s.io.api.core.v1.ServiceStatus
type K8sIoAPICoreV1ServiceStatus struct {

	// LoadBalancer contains the current status of the load-balancer,
	// if one is present.
	// +optional
	LoadBalancer *K8sIoAPICoreV1LoadBalancerStatus `json:"loadBalancer,omitempty"`
}

// Validate validates this k8s io api core v1 service status
func (m *K8sIoAPICoreV1ServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadBalancer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ServiceStatus) validateLoadBalancer(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadBalancer) { // not required
		return nil
	}

	if m.LoadBalancer != nil {
		if err := m.LoadBalancer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadBalancer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loadBalancer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 service status based on the context it is used
func (m *K8sIoAPICoreV1ServiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoadBalancer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1ServiceStatus) contextValidateLoadBalancer(ctx context.Context, formats strfmt.Registry) error {

	if m.LoadBalancer != nil {
		if err := m.LoadBalancer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadBalancer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loadBalancer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ServiceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ServiceStatus) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ServiceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}