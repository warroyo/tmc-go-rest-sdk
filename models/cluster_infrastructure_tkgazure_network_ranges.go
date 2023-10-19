// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInfrastructureTkgazureNetworkRanges Network range for the workload cluster.
//
// swagger:model cluster.infrastructure.tkgazure.NetworkRanges
type ClusterInfrastructureTkgazureNetworkRanges struct {

	// CIDRBlocks specifies one or more of IP address ranges.
	CidrBlocks []string `json:"cidrBlocks"`
}

// Validate validates this cluster infrastructure tkgazure network ranges
func (m *ClusterInfrastructureTkgazureNetworkRanges) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster infrastructure tkgazure network ranges based on context it is used
func (m *ClusterInfrastructureTkgazureNetworkRanges) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureNetworkRanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInfrastructureTkgazureNetworkRanges) UnmarshalBinary(b []byte) error {
	var res ClusterInfrastructureTkgazureNetworkRanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
