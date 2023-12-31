// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AksclusterNetworkConfig The network configuration for the AKS cluster.
//
// swagger:model akscluster.NetworkConfig
type AksclusterNetworkConfig struct {

	// DNS prefix of the cluster.
	DNSPrefix string `json:"dnsPrefix,omitempty"`

	// An IP address assigned to the Kubernetes DNS service.
	// It must be within the Kubernetes service address range specified in serviceCidr.
	DNSServiceIP string `json:"dnsServiceIp,omitempty"`

	// A CIDR notation IP range assigned to the Docker bridge network.
	// It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
	DockerBridgeCidr string `json:"dockerBridgeCidr,omitempty"`

	// The load balancer SKU for the cluster. The valid value is basic and standard.
	LoadBalancerSku string `json:"loadBalancerSku,omitempty"`

	// Network plugin of the cluster. The valid value is azure, kubenet and none.
	NetworkPlugin string `json:"networkPlugin,omitempty"`

	// Network policy of the cluster. The valid value is azure and calico.
	NetworkPolicy string `json:"networkPolicy,omitempty"`

	// The CIDR notation IP ranges from which to assign pod IPs.
	PodCidrs []string `json:"podCidrs"`

	// The CIDR notation IP ranges from which to assign service cluster IPs.
	ServiceCidrs []string `json:"serviceCidrs"`
}

// Validate validates this akscluster network config
func (m *AksclusterNetworkConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this akscluster network config based on context it is used
func (m *AksclusterNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterNetworkConfig) UnmarshalBinary(b []byte) error {
	var res AksclusterNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
