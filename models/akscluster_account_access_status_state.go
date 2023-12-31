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

// AksclusterAccountAccessStatusState The state of our access to the Azure account.
//
//   - STATE_UNSPECIFIED: Unspecified state.
//   - PENDING: Still in the process of accessing the customer's Azure account.
//   - SUCCEEDED: Successfully accessed customer's Azure account.
//   - FAILED: Failed to access customer's Azure account.
//   - INVALID: Invalid access definition for the Azure account.
//
// swagger:model akscluster.AccountAccessStatus.State
type AksclusterAccountAccessStatusState string

func NewAksclusterAccountAccessStatusState(value AksclusterAccountAccessStatusState) *AksclusterAccountAccessStatusState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AksclusterAccountAccessStatusState.
func (m AksclusterAccountAccessStatusState) Pointer() *AksclusterAccountAccessStatusState {
	return &m
}

const (

	// AksclusterAccountAccessStatusStateSTATEUNSPECIFIED captures enum value "STATE_UNSPECIFIED"
	AksclusterAccountAccessStatusStateSTATEUNSPECIFIED AksclusterAccountAccessStatusState = "STATE_UNSPECIFIED"

	// AksclusterAccountAccessStatusStatePENDING captures enum value "PENDING"
	AksclusterAccountAccessStatusStatePENDING AksclusterAccountAccessStatusState = "PENDING"

	// AksclusterAccountAccessStatusStateSUCCEEDED captures enum value "SUCCEEDED"
	AksclusterAccountAccessStatusStateSUCCEEDED AksclusterAccountAccessStatusState = "SUCCEEDED"

	// AksclusterAccountAccessStatusStateFAILED captures enum value "FAILED"
	AksclusterAccountAccessStatusStateFAILED AksclusterAccountAccessStatusState = "FAILED"

	// AksclusterAccountAccessStatusStateINVALID captures enum value "INVALID"
	AksclusterAccountAccessStatusStateINVALID AksclusterAccountAccessStatusState = "INVALID"
)

// for schema
var aksclusterAccountAccessStatusStateEnum []interface{}

func init() {
	var res []AksclusterAccountAccessStatusState
	if err := json.Unmarshal([]byte(`["STATE_UNSPECIFIED","PENDING","SUCCEEDED","FAILED","INVALID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aksclusterAccountAccessStatusStateEnum = append(aksclusterAccountAccessStatusStateEnum, v)
	}
}

func (m AksclusterAccountAccessStatusState) validateAksclusterAccountAccessStatusStateEnum(path, location string, value AksclusterAccountAccessStatusState) error {
	if err := validate.EnumCase(path, location, value, aksclusterAccountAccessStatusStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this akscluster account access status state
func (m AksclusterAccountAccessStatusState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAksclusterAccountAccessStatusStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this akscluster account access status state based on context it is used
func (m AksclusterAccountAccessStatusState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
