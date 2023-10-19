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

// AuditStatusPhase Phase describes the current state of the audit report generation.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING: Pending - to indicate that the audit is waiting to be started.
//   - RUNNING: Running - to indicate the audit is currently running
//
// i.e., the system is collecting the logs and generating the report.
//   - UPLOADING: Uploading - to indicate that the audit results are being uploaded to S3.
//   - COMPLETE: Complete - to indicate that the audit has completed and uploaded results to S3 successfully.
//   - PARTIALLY_COMPLETE: Partially_Complete - to indicate that the audit has completed with partial success.
//
// Only a subset of the requested date range is uploaded because of file size limits.
//   - ERROR: Error - to indicate that an error had occurred during the report generation.
//   - CANCEL: CANCEL - to indicate that the report generation is canceled.
//
// swagger:model audit.Status.Phase
type AuditStatusPhase string

func NewAuditStatusPhase(value AuditStatusPhase) *AuditStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AuditStatusPhase.
func (m AuditStatusPhase) Pointer() *AuditStatusPhase {
	return &m
}

const (

	// AuditStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	AuditStatusPhasePHASEUNSPECIFIED AuditStatusPhase = "PHASE_UNSPECIFIED"

	// AuditStatusPhasePENDING captures enum value "PENDING"
	AuditStatusPhasePENDING AuditStatusPhase = "PENDING"

	// AuditStatusPhaseRUNNING captures enum value "RUNNING"
	AuditStatusPhaseRUNNING AuditStatusPhase = "RUNNING"

	// AuditStatusPhaseUPLOADING captures enum value "UPLOADING"
	AuditStatusPhaseUPLOADING AuditStatusPhase = "UPLOADING"

	// AuditStatusPhaseCOMPLETE captures enum value "COMPLETE"
	AuditStatusPhaseCOMPLETE AuditStatusPhase = "COMPLETE"

	// AuditStatusPhasePARTIALLYCOMPLETE captures enum value "PARTIALLY_COMPLETE"
	AuditStatusPhasePARTIALLYCOMPLETE AuditStatusPhase = "PARTIALLY_COMPLETE"

	// AuditStatusPhaseERROR captures enum value "ERROR"
	AuditStatusPhaseERROR AuditStatusPhase = "ERROR"

	// AuditStatusPhaseCANCEL captures enum value "CANCEL"
	AuditStatusPhaseCANCEL AuditStatusPhase = "CANCEL"
)

// for schema
var auditStatusPhaseEnum []interface{}

func init() {
	var res []AuditStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","RUNNING","UPLOADING","COMPLETE","PARTIALLY_COMPLETE","ERROR","CANCEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditStatusPhaseEnum = append(auditStatusPhaseEnum, v)
	}
}

func (m AuditStatusPhase) validateAuditStatusPhaseEnum(path, location string, value AuditStatusPhase) error {
	if err := validate.EnumCase(path, location, value, auditStatusPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this audit status phase
func (m AuditStatusPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuditStatusPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this audit status phase based on context it is used
func (m AuditStatusPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}