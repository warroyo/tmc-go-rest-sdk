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

// ClusterDataprotectionBackuplocationStatus Status of the backup location resource.
//
// swagger:model cluster.dataprotection.backuplocation.Status
type ClusterDataprotectionBackuplocationStatus struct {

	// Permission for the backup location.
	AccessMode *ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode `json:"accessMode,omitempty"`

	// A list of available phases for backup location object.
	AvailablePhases []*ClusterDataprotectionBackuplocationStatusPhase `json:"availablePhases"`

	// The conditions attached to this backup location object.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// The SHA value inside the file `<bucket>/metadata/revision`. Updating the SHA in that file will force
	// a sync to happen sooner.
	LastSyncRevision string `json:"lastSyncRevision,omitempty"`

	// The timestamp when last sync happened.
	// Format: date-time
	LastSyncTimestamp strfmt.DateTime `json:"lastSyncTimestamp,omitempty"`

	// A message about the backup storage location's status.
	Message string `json:"message,omitempty"`

	// The resource generation the current status applies to.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// The current state of the backup location.
	Phase *ClusterDataprotectionBackuplocationStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`
}

// Validate validates this cluster dataprotection backuplocation status
func (m *ClusterDataprotectionBackuplocationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailablePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSyncTimestamp(formats); err != nil {
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

func (m *ClusterDataprotectionBackuplocationStatus) validateAccessMode(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessMode) { // not required
		return nil
	}

	if m.AccessMode != nil {
		if err := m.AccessMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessMode")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackuplocationStatus) validateAvailablePhases(formats strfmt.Registry) error {
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

func (m *ClusterDataprotectionBackuplocationStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *ClusterDataprotectionBackuplocationStatus) validateLastSyncTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSyncTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSyncTimestamp", "body", "date-time", m.LastSyncTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDataprotectionBackuplocationStatus) validatePhase(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster dataprotection backuplocation status based on the context it is used
func (m *ClusterDataprotectionBackuplocationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailablePhases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
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

func (m *ClusterDataprotectionBackuplocationStatus) contextValidateAccessMode(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessMode != nil {
		if err := m.AccessMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessMode")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackuplocationStatus) contextValidateAvailablePhases(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterDataprotectionBackuplocationStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		if val, ok := m.Conditions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackuplocationStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterDataprotectionBackuplocationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionBackuplocationStatus) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionBackuplocationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
