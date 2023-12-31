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

// CommonClusterTKGVsphereReleases Kubernetes releases for TKG vSphere.
//
// swagger:model common.cluster.TKGVsphereReleases
type CommonClusterTKGVsphereReleases struct {

	// Version of Kubernetes used in this release.
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`

	// The vSphere virtual machine templates associated with this TKG release.
	Templates []*CommonClusterTKGVsphereTemplates `json:"templates"`

	// Version of this TKG release.
	Version string `json:"version,omitempty"`
}

// Validate validates this common cluster t k g vsphere releases
func (m *CommonClusterTKGVsphereReleases) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterTKGVsphereReleases) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	for i := 0; i < len(m.Templates); i++ {
		if swag.IsZero(m.Templates[i]) { // not required
			continue
		}

		if m.Templates[i] != nil {
			if err := m.Templates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common cluster t k g vsphere releases based on the context it is used
func (m *CommonClusterTKGVsphereReleases) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterTKGVsphereReleases) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Templates); i++ {

		if m.Templates[i] != nil {
			if err := m.Templates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGVsphereReleases) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGVsphereReleases) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGVsphereReleases
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
