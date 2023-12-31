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

// ClusterNamespaceFluxcdHelmReleaseSpec Spec of the Helm Release.
//
// swagger:model cluster.namespace.fluxcd.helm.release.Spec
type ClusterNamespaceFluxcdHelmReleaseSpec struct {

	// Reference to the chart which will be installed.
	ChartRef *ClusterNamespaceFluxcdHelmReleaseChartRef `json:"chartRef,omitempty"`

	// Inline values in yaml format.
	InlineConfiguration string `json:"inlineConfiguration,omitempty"`

	// Interval at which to reconcile the Helm release.
	Interval string `json:"interval,omitempty"`

	// Name of target namespace.
	TargetNamespace string `json:"targetNamespace,omitempty"`
}

// Validate validates this cluster namespace fluxcd helm release spec
func (m *ClusterNamespaceFluxcdHelmReleaseSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceFluxcdHelmReleaseSpec) validateChartRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartRef) { // not required
		return nil
	}

	if m.ChartRef != nil {
		if err := m.ChartRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chartRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chartRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster namespace fluxcd helm release spec based on the context it is used
func (m *ClusterNamespaceFluxcdHelmReleaseSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChartRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNamespaceFluxcdHelmReleaseSpec) contextValidateChartRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ChartRef != nil {
		if err := m.ChartRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chartRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chartRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdHelmReleaseSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceFluxcdHelmReleaseSpec) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceFluxcdHelmReleaseSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
