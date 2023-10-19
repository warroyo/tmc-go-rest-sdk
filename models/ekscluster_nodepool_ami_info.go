// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EksclusterNodepoolAmiInfo AMI info for the nodepool.
//
// swagger:model ekscluster.nodepool.AmiInfo
type EksclusterNodepoolAmiInfo struct {

	// AMI id if custom AMI is specified. It cannot be used if launch template id is specified.
	AmiID string `json:"amiId,omitempty"`

	// Override bootstrap command for custom AMI.
	OverrideBootstrapCmd string `json:"overrideBootstrapCmd,omitempty"`
}

// Validate validates this ekscluster nodepool ami info
func (m *EksclusterNodepoolAmiInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ekscluster nodepool ami info based on context it is used
func (m *EksclusterNodepoolAmiInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EksclusterNodepoolAmiInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EksclusterNodepoolAmiInfo) UnmarshalBinary(b []byte) error {
	var res EksclusterNodepoolAmiInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
