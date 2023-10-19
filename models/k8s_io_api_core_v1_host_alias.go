// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1HostAlias HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the
// pod's hosts file.
//
// swagger:model k8s.io.api.core.v1.HostAlias
type K8sIoAPICoreV1HostAlias struct {

	// Hostnames for the above IP address.
	Hostnames []string `json:"hostnames"`

	// IP address of the host file entry.
	IP string `json:"ip,omitempty"`
}

// Validate validates this k8s io api core v1 host alias
func (m *K8sIoAPICoreV1HostAlias) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 host alias based on context it is used
func (m *K8sIoAPICoreV1HostAlias) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1HostAlias) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1HostAlias) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1HostAlias
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
