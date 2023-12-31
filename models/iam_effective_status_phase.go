// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IamEffectiveStatusPhase Phase of an iam policy.
//
//   - PHASE_UNSPECIFIED: Phase Unspecified phase is the default state set when an effective iam policy is first created/updated.
//   - APPLYING: Applying state is set when we are working on applying the policy.
//   - APPLIED: Applied state is set when we have successfully applied the policy.
//   - ERROR: Error state is set when there was a failure while applying the policy.
//   - DELAYED: Delayed state is set when we are unable to apply the policy for a long time.
//
// swagger:model iam.effective.Status.Phase
type IamEffectiveStatusPhase string

func NewIamEffectiveStatusPhase(value IamEffectiveStatusPhase) *IamEffectiveStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IamEffectiveStatusPhase.
func (m IamEffectiveStatusPhase) Pointer() *IamEffectiveStatusPhase {
	return &m
}

const (

	// IamEffectiveStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	IamEffectiveStatusPhasePHASEUNSPECIFIED IamEffectiveStatusPhase = "PHASE_UNSPECIFIED"

	// IamEffectiveStatusPhaseAPPLYING captures enum value "APPLYING"
	IamEffectiveStatusPhaseAPPLYING IamEffectiveStatusPhase = "APPLYING"

	// IamEffectiveStatusPhaseAPPLIED captures enum value "APPLIED"
	IamEffectiveStatusPhaseAPPLIED IamEffectiveStatusPhase = "APPLIED"

	// IamEffectiveStatusPhaseERROR captures enum value "ERROR"
	IamEffectiveStatusPhaseERROR IamEffectiveStatusPhase = "ERROR"

	// IamEffectiveStatusPhaseDELAYED captures enum value "DELAYED"
	IamEffectiveStatusPhaseDELAYED IamEffectiveStatusPhase = "DELAYED"
)

// for schema
var iamEffectiveStatusPhaseEnum []interface{}

func init() {
	var res []IamEffectiveStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","APPLYING","APPLIED","ERROR","DELAYED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iamEffectiveStatusPhaseEnum = append(iamEffectiveStatusPhaseEnum, v)
	}
}

func (m IamEffectiveStatusPhase) validateIamEffectiveStatusPhaseEnum(path, location string, value IamEffectiveStatusPhase) error {
	if err := validate.EnumCase(path, location, value, iamEffectiveStatusPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this iam effective status phase
func (m IamEffectiveStatusPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIamEffectiveStatusPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this iam effective status phase based on context it is used
func (m IamEffectiveStatusPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
