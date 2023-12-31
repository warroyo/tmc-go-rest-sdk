// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EksclusterLogging EKS logging configuration.
// Refer https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html for more info
//
// swagger:model ekscluster.Logging
type EksclusterLogging struct {

	// Enable API server logs.
	APIServer bool `json:"apiServer,omitempty"`

	// Enable audit logs.
	Audit bool `json:"audit,omitempty"`

	// Enable authenticator logs.
	Authenticator bool `json:"authenticator,omitempty"`

	// Enable controller manager logs.
	ControllerManager bool `json:"controllerManager,omitempty"`

	// Enable scheduler logs.
	Scheduler bool `json:"scheduler,omitempty"`
}

// Validate validates this ekscluster logging
func (m *EksclusterLogging) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ekscluster logging based on context it is used
func (m *EksclusterLogging) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterLogging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterLogging) UnmarshalBinary(b []byte) error {
	var res EksclusterLogging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
