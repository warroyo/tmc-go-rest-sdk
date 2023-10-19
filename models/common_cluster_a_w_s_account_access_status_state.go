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

// CommonClusterAWSAccountAccessStatusState The state of our access to the AWS account.
//
//   - STATE_UNSPECIFIED: Unspecified state.
//   - PENDING: Still in the process of accessing the customer's AWS account.
//   - SUCCEEDED: Successfully accessed customer's AWS account.
//   - FAILED: Failed to access customer's AWS account.
//
// swagger:model common.cluster.AWSAccountAccessStatus.State
type CommonClusterAWSAccountAccessStatusState string

func NewCommonClusterAWSAccountAccessStatusState(value CommonClusterAWSAccountAccessStatusState) *CommonClusterAWSAccountAccessStatusState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CommonClusterAWSAccountAccessStatusState.
func (m CommonClusterAWSAccountAccessStatusState) Pointer() *CommonClusterAWSAccountAccessStatusState {
	return &m
}

const (

	// CommonClusterAWSAccountAccessStatusStateSTATEUNSPECIFIED captures enum value "STATE_UNSPECIFIED"
	CommonClusterAWSAccountAccessStatusStateSTATEUNSPECIFIED CommonClusterAWSAccountAccessStatusState = "STATE_UNSPECIFIED"

	// CommonClusterAWSAccountAccessStatusStatePENDING captures enum value "PENDING"
	CommonClusterAWSAccountAccessStatusStatePENDING CommonClusterAWSAccountAccessStatusState = "PENDING"

	// CommonClusterAWSAccountAccessStatusStateSUCCEEDED captures enum value "SUCCEEDED"
	CommonClusterAWSAccountAccessStatusStateSUCCEEDED CommonClusterAWSAccountAccessStatusState = "SUCCEEDED"

	// CommonClusterAWSAccountAccessStatusStateFAILED captures enum value "FAILED"
	CommonClusterAWSAccountAccessStatusStateFAILED CommonClusterAWSAccountAccessStatusState = "FAILED"
)

// for schema
var commonClusterAWSAccountAccessStatusStateEnum []interface{}

func init() {
	var res []CommonClusterAWSAccountAccessStatusState
	if err := json.Unmarshal([]byte(`["STATE_UNSPECIFIED","PENDING","SUCCEEDED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonClusterAWSAccountAccessStatusStateEnum = append(commonClusterAWSAccountAccessStatusStateEnum, v)
	}
}

func (m CommonClusterAWSAccountAccessStatusState) validateCommonClusterAWSAccountAccessStatusStateEnum(path, location string, value CommonClusterAWSAccountAccessStatusState) error {
	if err := validate.EnumCase(path, location, value, commonClusterAWSAccountAccessStatusStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this common cluster a w s account access status state
func (m CommonClusterAWSAccountAccessStatusState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCommonClusterAWSAccountAccessStatusStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this common cluster a w s account access status state based on context it is used
func (m CommonClusterAWSAccountAccessStatusState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}