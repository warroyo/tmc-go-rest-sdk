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

// AccountCredentialStatus Status of the credential.
//
// swagger:model account.credential.Status
type AccountCredentialStatus struct {

	// List of available phases for a credential.
	AvailablePhases []*AccountCredentialStatusPhase `json:"availablePhases"`

	// The conditions describing a credential's status.
	//
	// - "Scheduled"
	//      * True when the credential is being created.
	//      * False if the credential cannot be created.
	// - "Ready"
	//      * True when the credential is created.
	//      * False if the credential creation failed.
	// - "Validating"
	//      * True when and if the intended service starts validating the credential.
	//      * False if the credential validation process failed. This value does not
	//        represent the validity of the credential.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the credential.
	Phase *AccountCredentialStatusPhase `json:"phase,omitempty"`

	// Additional information about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// The list of references. i.e. resources using the credential.
	References []*AccountCredentialReference `json:"references"`
}

// Validate validates this account credential status
func (m *AccountCredentialStatus) Validate(formats strfmt.Registry) error {
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

	if err := m.validateReferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialStatus) validateAvailablePhases(formats strfmt.Registry) error {
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

func (m *AccountCredentialStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *AccountCredentialStatus) validatePhase(formats strfmt.Registry) error {
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

func (m *AccountCredentialStatus) validateReferences(formats strfmt.Registry) error {
	if swag.IsZero(m.References) { // not required
		return nil
	}

	for i := 0; i < len(m.References); i++ {
		if swag.IsZero(m.References[i]) { // not required
			continue
		}

		if m.References[i] != nil {
			if err := m.References[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("references" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account credential status based on the context it is used
func (m *AccountCredentialStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateReferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCredentialStatus) contextValidateAvailablePhases(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AccountCredentialStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		if val, ok := m.Conditions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *AccountCredentialStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AccountCredentialStatus) contextValidateReferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.References); i++ {

		if m.References[i] != nil {
			if err := m.References[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("references" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCredentialStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCredentialStatus) UnmarshalBinary(b []byte) error {
	var res AccountCredentialStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
