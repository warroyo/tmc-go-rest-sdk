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

// K8sIoAPICoreV1NodeDaemonEndpoints NodeDaemonEndpoints lists ports opened by daemons running on the Node.
//
// swagger:model k8s.io.api.core.v1.NodeDaemonEndpoints
type K8sIoAPICoreV1NodeDaemonEndpoints struct {

	// Endpoint on which Kubelet is listening.
	// +optional
	KubeletEndpoint *K8sIoAPICoreV1DaemonEndpoint `json:"kubeletEndpoint,omitempty"`
}

// Validate validates this k8s io api core v1 node daemon endpoints
func (m *K8sIoAPICoreV1NodeDaemonEndpoints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKubeletEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1NodeDaemonEndpoints) validateKubeletEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.KubeletEndpoint) { // not required
		return nil
	}

	if m.KubeletEndpoint != nil {
		if err := m.KubeletEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeletEndpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubeletEndpoint")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api core v1 node daemon endpoints based on the context it is used
func (m *K8sIoAPICoreV1NodeDaemonEndpoints) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKubeletEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1NodeDaemonEndpoints) contextValidateKubeletEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.KubeletEndpoint != nil {
		if err := m.KubeletEndpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeletEndpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubeletEndpoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1NodeDaemonEndpoints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1NodeDaemonEndpoints) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1NodeDaemonEndpoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
