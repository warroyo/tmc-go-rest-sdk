// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoAPICoreV1ClientIPConfig ClientIPConfig represents the configurations of Client IP based session affinity.
//
// swagger:model k8s.io.api.core.v1.ClientIPConfig
type K8sIoAPICoreV1ClientIPConfig struct {

	// timeoutSeconds specifies the seconds of ClientIP type session sticky time.
	// The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP".
	// Default value is 10800(for 3 hours).
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
}

// Validate validates this k8s io api core v1 client IP config
func (m *K8sIoAPICoreV1ClientIPConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io api core v1 client IP config based on context it is used
func (m *K8sIoAPICoreV1ClientIPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1ClientIPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1ClientIPConfig) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1ClientIPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}