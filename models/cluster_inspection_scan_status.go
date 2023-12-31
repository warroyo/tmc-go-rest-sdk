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
	"github.com/go-openapi/validate"
)

// ClusterInspectionScanStatus Status of the scan inspection.
//
// swagger:model cluster.inspection.scan.Status
type ClusterInspectionScanStatus struct {

	// Available phases of the inspection scan resource.
	AvailablePhases []*ClusterInspectionScanStatusPhase `json:"availablePhases"`

	// Condition 'Scheduled' with Status 'Unknown' indicates that the inspection is pending
	// Condition 'Scheduled' with 'True' and Condition 'Ready' with 'Unknown' indicates that the inspection is running
	// Condition 'Ready' with 'True' indicates that the inspection is complete
	// Condition 'Ready' with 'False' indicates that the inspection is in error state.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the inspection scan based on conditions. If state is 'PHASE_UNSPECIFIED', use conditions to
	// interpret the state of the inspection.
	Phase *ClusterInspectionScanStatusPhase `json:"phase,omitempty"`

	// Additional information e.g., reason for ERROR state.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// Report details.
	Report *ClusterInspectionScanReport `json:"report,omitempty"`
}

// Validate validates this cluster inspection scan status
func (m *ClusterInspectionScanStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailablePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInspectionScanStatus) validateAvailablePhases(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailablePhases) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailablePhases); i++ {
		if swag.IsZero(m.AvailablePhases[i]) { // not required
			continue
		}

		if m.AvailablePhases[i] != nil {
			if err := m.AvailablePhases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterInspectionScanStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for k := range m.Conditions {

		if err := validate.Required("conditions"+"."+k, "body", m.Conditions[k]); err != nil {
			return err
		}
		if val, ok := m.Conditions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterInspectionScanStatus) validatePhase(formats strfmt.Registry) error {
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

func (m *ClusterInspectionScanStatus) validateReport(formats strfmt.Registry) error {
	if swag.IsZero(m.Report) { // not required
		return nil
	}

	if m.Report != nil {
		if err := m.Report.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster inspection scan status based on the context it is used
func (m *ClusterInspectionScanStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailablePhases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInspectionScanStatus) contextValidateAvailablePhases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailablePhases); i++ {

		if m.AvailablePhases[i] != nil {
			if err := m.AvailablePhases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterInspectionScanStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		if val, ok := m.Conditions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterInspectionScanStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterInspectionScanStatus) contextValidateReport(ctx context.Context, formats strfmt.Registry) error {

	if m.Report != nil {
		if err := m.Report.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("report")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInspectionScanStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInspectionScanStatus) UnmarshalBinary(b []byte) error {
	var res ClusterInspectionScanStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
