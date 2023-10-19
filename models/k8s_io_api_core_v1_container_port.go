// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1ContainerPort ContainerPort represents a network port in a single container.
//
// swagger:model k8s.io.api.core.v1.ContainerPort
type K8sIoAPICoreV1ContainerPort struct {

	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort,omitempty"`

	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"hostIP,omitempty"`

	// Number of port to expose on the host.
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"hostPort,omitempty"`

	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
	// named port in a pod must have a unique name. Name for the port that can be
	// referred to by services.
	// +optional
	Name string `json:"name,omitempty"`

	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this k8s io api core v1 container port
func (m *K8sIoAPICoreV1ContainerPort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 container port based on context it is used
func (m *K8sIoAPICoreV1ContainerPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ContainerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ContainerPort) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ContainerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
