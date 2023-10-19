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

// ClusterDataprotectionScheduleStatusPhase The lifecycle phase of a schedule backup.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the schedule object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when schedule is being created on the cluster.
//   - NEW: The schedule has been created but not yet processed by velero.
//   - ENABLED: The schedule has been validated and will now be triggering backups according to the schedule spec.
//   - FAILEDVALIDATION: The schedule has failed the velero controller's validations and therefore will not trigger backups.
//   - PENDING_DELETE: Pending delete is set when the object deletion is being processed by the service.
//   - DELETING: The phase when schedule is being deleted.
//   - UPDATING: The phase when schedule is being updated.
//   - PAUSED: The phase when schedule is being paused.
//
// swagger:model cluster.dataprotection.schedule.Status.Phase
type ClusterDataprotectionScheduleStatusPhase string

func NewClusterDataprotectionScheduleStatusPhase(value ClusterDataprotectionScheduleStatusPhase) *ClusterDataprotectionScheduleStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterDataprotectionScheduleStatusPhase.
func (m ClusterDataprotectionScheduleStatusPhase) Pointer() *ClusterDataprotectionScheduleStatusPhase {
	return &m
}

const (

	// ClusterDataprotectionScheduleStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	ClusterDataprotectionScheduleStatusPhasePHASEUNSPECIFIED ClusterDataprotectionScheduleStatusPhase = "PHASE_UNSPECIFIED"

	// ClusterDataprotectionScheduleStatusPhasePENDING captures enum value "PENDING"
	ClusterDataprotectionScheduleStatusPhasePENDING ClusterDataprotectionScheduleStatusPhase = "PENDING"

	// ClusterDataprotectionScheduleStatusPhaseCREATING captures enum value "CREATING"
	ClusterDataprotectionScheduleStatusPhaseCREATING ClusterDataprotectionScheduleStatusPhase = "CREATING"

	// ClusterDataprotectionScheduleStatusPhaseNEW captures enum value "NEW"
	ClusterDataprotectionScheduleStatusPhaseNEW ClusterDataprotectionScheduleStatusPhase = "NEW"

	// ClusterDataprotectionScheduleStatusPhaseENABLED captures enum value "ENABLED"
	ClusterDataprotectionScheduleStatusPhaseENABLED ClusterDataprotectionScheduleStatusPhase = "ENABLED"

	// ClusterDataprotectionScheduleStatusPhaseFAILEDVALIDATION captures enum value "FAILEDVALIDATION"
	ClusterDataprotectionScheduleStatusPhaseFAILEDVALIDATION ClusterDataprotectionScheduleStatusPhase = "FAILEDVALIDATION"

	// ClusterDataprotectionScheduleStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE"
	ClusterDataprotectionScheduleStatusPhasePENDINGDELETE ClusterDataprotectionScheduleStatusPhase = "PENDING_DELETE"

	// ClusterDataprotectionScheduleStatusPhaseDELETING captures enum value "DELETING"
	ClusterDataprotectionScheduleStatusPhaseDELETING ClusterDataprotectionScheduleStatusPhase = "DELETING"

	// ClusterDataprotectionScheduleStatusPhaseUPDATING captures enum value "UPDATING"
	ClusterDataprotectionScheduleStatusPhaseUPDATING ClusterDataprotectionScheduleStatusPhase = "UPDATING"

	// ClusterDataprotectionScheduleStatusPhasePAUSED captures enum value "PAUSED"
	ClusterDataprotectionScheduleStatusPhasePAUSED ClusterDataprotectionScheduleStatusPhase = "PAUSED"
)

// for schema
var clusterDataprotectionScheduleStatusPhaseEnum []interface{}

func init() {
	var res []ClusterDataprotectionScheduleStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","NEW","ENABLED","FAILEDVALIDATION","PENDING_DELETE","DELETING","UPDATING","PAUSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterDataprotectionScheduleStatusPhaseEnum = append(clusterDataprotectionScheduleStatusPhaseEnum, v)
	}
}

func (m ClusterDataprotectionScheduleStatusPhase) validateClusterDataprotectionScheduleStatusPhaseEnum(path, location string, value ClusterDataprotectionScheduleStatusPhase) error {
	if err := validate.EnumCase(path, location, value, clusterDataprotectionScheduleStatusPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster dataprotection schedule status phase
func (m ClusterDataprotectionScheduleStatusPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterDataprotectionScheduleStatusPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster dataprotection schedule status phase based on context it is used
func (m ClusterDataprotectionScheduleStatusPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}