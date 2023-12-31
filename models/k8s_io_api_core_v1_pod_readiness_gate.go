// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1PodReadinessGate PodReadinessGate contains the reference to a pod condition
//
// swagger:model k8s.io.api.core.v1.PodReadinessGate
type K8sIoAPICoreV1PodReadinessGate struct {

	// ConditionType refers to a condition in the pod's condition list with matching type.
	ConditionType string `json:"conditionType,omitempty"`
}

// Validate validates this k8s io api core v1 pod readiness gate
func (m *K8sIoAPICoreV1PodReadinessGate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 pod readiness gate based on context it is used
func (m *K8sIoAPICoreV1PodReadinessGate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1PodReadinessGate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1PodReadinessGate) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1PodReadinessGate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
