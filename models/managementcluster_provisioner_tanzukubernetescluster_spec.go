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

// ManagementclusterProvisionerTanzukubernetesclusterSpec Spec of the cluster.
//
// swagger:model managementcluster.provisioner.tanzukubernetescluster.Spec
type ManagementclusterProvisionerTanzukubernetesclusterSpec struct {

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Name of the image registry configuration to use.
	ImageRegistry string `json:"imageRegistry,omitempty"`

	// Name of the proxy configuration to use.
	ProxyName string `json:"proxyName,omitempty"`

	// TKG AWS cluster spec.
	TkgAws *ClusterInfrastructureTkgawsSpec `json:"tkgAws,omitempty"`

	// TKG Azure cluster spec.
	TkgAzure *ClusterInfrastructureTkgazureSpec `json:"tkgAzure,omitempty"`

	// TKG Service vSphere cluster spec.
	TkgServiceVsphere *ClusterInfrastructureTkgservicevsphereSpec `json:"tkgServiceVsphere,omitempty"`

	// TKG vSphere cluster spec.
	TkgVsphere *ClusterInfrastructureTkgvsphereSpec `json:"tkgVsphere,omitempty"`

	// TMC-managed flag indicates if the cluster is managed by tmc.
	TmcManaged bool `json:"tmcManaged,omitempty"`

	// The cluster topology.
	Topology *ManagementclusterProvisionerTanzukubernetesclusterTopology `json:"topology,omitempty"`
}

// Validate validates this managementcluster provisioner tanzukubernetescluster spec
func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTkgAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTkgAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTkgServiceVsphere(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTkgVsphere(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) validateTkgAws(formats strfmt.Registry) error {
	if swag.IsZero(m.TkgAws) { // not required
		return nil
	}

	if m.TkgAws != nil {
		if err := m.TkgAws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgAws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgAws")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) validateTkgAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.TkgAzure) { // not required
		return nil
	}

	if m.TkgAzure != nil {
		if err := m.TkgAzure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgAzure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgAzure")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) validateTkgServiceVsphere(formats strfmt.Registry) error {
	if swag.IsZero(m.TkgServiceVsphere) { // not required
		return nil
	}

	if m.TkgServiceVsphere != nil {
		if err := m.TkgServiceVsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgServiceVsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgServiceVsphere")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) validateTkgVsphere(formats strfmt.Registry) error {
	if swag.IsZero(m.TkgVsphere) { // not required
		return nil
	}

	if m.TkgVsphere != nil {
		if err := m.TkgVsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgVsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgVsphere")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) validateTopology(formats strfmt.Registry) error {
	if swag.IsZero(m.Topology) { // not required
		return nil
	}

	if m.Topology != nil {
		if err := m.Topology.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this managementcluster provisioner tanzukubernetescluster spec based on the context it is used
func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTkgAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTkgAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTkgServiceVsphere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTkgVsphere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopology(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) contextValidateTkgAws(ctx context.Context, formats strfmt.Registry) error {

	if m.TkgAws != nil {
		if err := m.TkgAws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgAws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgAws")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) contextValidateTkgAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.TkgAzure != nil {
		if err := m.TkgAzure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgAzure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgAzure")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) contextValidateTkgServiceVsphere(ctx context.Context, formats strfmt.Registry) error {

	if m.TkgServiceVsphere != nil {
		if err := m.TkgServiceVsphere.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgServiceVsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgServiceVsphere")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) contextValidateTkgVsphere(ctx context.Context, formats strfmt.Registry) error {

	if m.TkgVsphere != nil {
		if err := m.TkgVsphere.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tkgVsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tkgVsphere")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) contextValidateTopology(ctx context.Context, formats strfmt.Registry) error {

	if m.Topology != nil {
		if err := m.Topology.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterProvisionerTanzukubernetesclusterSpec) UnmarshalBinary(b []byte) error {
	var res ManagementclusterProvisionerTanzukubernetesclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
