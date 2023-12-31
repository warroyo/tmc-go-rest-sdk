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

// AksclusterClusterConfig The config of an AKS cluster.
//
// swagger:model akscluster.ClusterConfig
type AksclusterClusterConfig struct {

	// The access config.
	AccessConfig *AksclusterAccessConfig `json:"accessConfig,omitempty"`

	// The Addons config.
	AddonsConfig *AksclusterAddonsConfig `json:"addonsConfig,omitempty"`

	// The access config for the cluster API server.
	APIServerAccessConfig *AksclusterAPIServerAccessConfig `json:"apiServerAccessConfig,omitempty"`

	// The auto upgrade config.
	AutoUpgradeConfig *AksclusterAutoUpgradeConfig `json:"autoUpgradeConfig,omitempty"`

	// The resource ID of the disk encryption set to use for enabling
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`

	// The linux VMs config.
	LinuxConfig *AksclusterLinuxConfig `json:"linuxConfig,omitempty"`

	// The geo-location where the resource lives for the cluster.
	Location string `json:"location,omitempty"`

	// Kubernetes network config.
	NetworkConfig *AksclusterNetworkConfig `json:"networkConfig,omitempty"`

	// The name of the resource group containing node pool nodes.
	NodeResourceGroupName string `json:"nodeResourceGroupName,omitempty"`

	// The name of the proxy credential to be used for the cluster.
	// If specified, the AKS cluster will be created with the proxy configurations specified in the proxy credential.
	ProxyName string `json:"proxyName,omitempty"`

	// SKU of the cluster.
	Sku *AksclusterClusterSKU `json:"sku,omitempty"`

	// The storage config.
	StorageConfig *AksclusterStorageConfig `json:"storageConfig,omitempty"`

	// The metadata to apply to the cluster to assist with categorization and organization.
	Tags map[string]string `json:"tags,omitempty"`

	// Kubernetes version of the cluster.
	Version string `json:"version,omitempty"`
}

// Validate validates this akscluster cluster config
func (m *AksclusterClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddonsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIServerAccessConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoUpgradeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinuxConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterClusterConfig) validateAccessConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if m.AccessConfig != nil {
		if err := m.AccessConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateAddonsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AddonsConfig) { // not required
		return nil
	}

	if m.AddonsConfig != nil {
		if err := m.AddonsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addonsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addonsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateAPIServerAccessConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.APIServerAccessConfig) { // not required
		return nil
	}

	if m.APIServerAccessConfig != nil {
		if err := m.APIServerAccessConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiServerAccessConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiServerAccessConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateAutoUpgradeConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoUpgradeConfig) { // not required
		return nil
	}

	if m.AutoUpgradeConfig != nil {
		if err := m.AutoUpgradeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpgradeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoUpgradeConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateLinuxConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LinuxConfig) { // not required
		return nil
	}

	if m.LinuxConfig != nil {
		if err := m.LinuxConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linuxConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linuxConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateSku(formats strfmt.Registry) error {
	if swag.IsZero(m.Sku) { // not required
		return nil
	}

	if m.Sku != nil {
		if err := m.Sku.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sku")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sku")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) validateStorageConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageConfig) { // not required
		return nil
	}

	if m.StorageConfig != nil {
		if err := m.StorageConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this akscluster cluster config based on the context it is used
func (m *AksclusterClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddonsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIServerAccessConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoUpgradeConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinuxConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSku(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterClusterConfig) contextValidateAccessConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessConfig != nil {
		if err := m.AccessConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateAddonsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AddonsConfig != nil {
		if err := m.AddonsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addonsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addonsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateAPIServerAccessConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.APIServerAccessConfig != nil {
		if err := m.APIServerAccessConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiServerAccessConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiServerAccessConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateAutoUpgradeConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoUpgradeConfig != nil {
		if err := m.AutoUpgradeConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpgradeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoUpgradeConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateLinuxConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LinuxConfig != nil {
		if err := m.LinuxConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linuxConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linuxConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateSku(ctx context.Context, formats strfmt.Registry) error {

	if m.Sku != nil {
		if err := m.Sku.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sku")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sku")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterConfig) contextValidateStorageConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageConfig != nil {
		if err := m.StorageConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterClusterConfig) UnmarshalBinary(b []byte) error {
	var res AksclusterClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
