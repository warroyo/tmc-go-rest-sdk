// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementclusterProvisionerTanzukubernetesclusterTopology The cluster topology.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.Topology
type ManagementclusterProvisionerTanzukubernetesclusterTopology struct {

	// The name of the cluster class for the cluster.
	ClusterClass string `json:"clusterClass,omitempty"`

	// Control plane specific configuration.
	ControlPlane *ManagementclusterProvisionerTanzukubernetesclusterControlPlane `json:"controlPlane,omitempty"`

	// The core addons.
	CoreAddons []*ManagementclusterProvisionerTanzukubernetesclusterCoreAddon `json:"coreAddons"`

	// Network specific configuration.
	Network *ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings `json:"network,omitempty"`

	// Nodepool definition for the cluster.
	NodePools []*ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition `json:"nodePools"`

	// Variables configuration for the cluster.
	Variables []*ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable `json:"variables"`

	// Kubernetes version of the cluster.
	Version string `json:"version,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster topology
func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlPlane(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoreAddons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodePools(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) validateControlPlane(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlPlane) { // not required
		return nil
	}

	if m.ControlPlane != nil {
		if err := m.ControlPlane.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlPlane")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlPlane")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) validateCoreAddons(formats strfmt.Registry) error {
	if swag.IsZero(m.CoreAddons) { // not required
		return nil
	}

	for i := 0; i < len(m.CoreAddons); i++ {
		if swag.IsZero(m.CoreAddons[i]) { // not required
			continue
		}

		if m.CoreAddons[i] != nil {
			if err := m.CoreAddons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coreAddons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("coreAddons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) validateNodePools(formats strfmt.Registry) error {
	if swag.IsZero(m.NodePools) { // not required
		return nil
	}

	for i := 0; i < len(m.NodePools); i++ {
		if swag.IsZero(m.NodePools[i]) { // not required
			continue
		}

		if m.NodePools[i] != nil {
			if err := m.NodePools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodePools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodePools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) validateVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster topology based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlPlane(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoreAddons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodePools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) contextValidateControlPlane(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlPlane != nil {
		if err := m.ControlPlane.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlPlane")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlPlane")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) contextValidateCoreAddons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CoreAddons); i++ {

		if m.CoreAddons[i] != nil {
			if err := m.CoreAddons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coreAddons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("coreAddons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) contextValidateNodePools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodePools); i++ {

		if m.NodePools[i] != nil {
			if err := m.NodePools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodePools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodePools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variables); i++ {

		if m.Variables[i] != nil {
			if err := m.Variables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterTopology) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
