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

// ClustergroupNamespaceFluxcdHelmReleaseStatus Status of the Release.
//
// swagger:model clustergroup.namespace.fluxcd.helm.release.Status
type ClustergroupNamespaceFluxcdHelmReleaseStatus struct {

	// Details contains information about the Cluster Group helm release being applied on member Clusters.
	Details *CommonBatchDetails `json:"details,omitempty"`

	// Generation value at the time this status was updated.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// Phase of the Cluster Group helm release application on member Clusters.
	Phase *CommonBatchPhase `json:"phase,omitempty"`
}

// Validate validates this clustergroup namespace fluxcd helm release status
func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clustergroup namespace fluxcd helm release status based on the context it is used
func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.Details != nil {
		if err := m.Details.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustergroupNamespaceFluxcdHelmReleaseStatus) UnmarshalBinary(b []byte) error {
	var res ClustergroupNamespaceFluxcdHelmReleaseStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
