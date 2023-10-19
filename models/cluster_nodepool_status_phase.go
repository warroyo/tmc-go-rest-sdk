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

// ClusterNodepoolStatusPhase Phase of the nodepool resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - CREATING: Resource is pending processing.
//   - READY: Resource is in ready state.
//   - ERROR: Error in processing.
//   - DELETING: Resource is being deleted.
//   - RESIZING: Resizing state.
//   - UPGRADING: An upgrade is in progress.
//   - UPGRADE_FAILED: An upgrade has failed.
//   - WAITING: The cluster is not created yet. so wait till then.
//   - UPDATING: A generic phase used for TKGS nodepool as actual individual phase cannot be determined.
//
// swagger:model cluster.nodepool.Status.Phase
type ClusterNodepoolStatusPhase string

func NewClusterNodepoolStatusPhase(value ClusterNodepoolStatusPhase) *ClusterNodepoolStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterNodepoolStatusPhase.
func (m ClusterNodepoolStatusPhase) Pointer() *ClusterNodepoolStatusPhase {
	return &m
}

const (

	// ClusterNodepoolStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	ClusterNodepoolStatusPhasePHASEUNSPECIFIED ClusterNodepoolStatusPhase = "PHASE_UNSPECIFIED"

	// ClusterNodepoolStatusPhaseCREATING captures enum value "CREATING"
	ClusterNodepoolStatusPhaseCREATING ClusterNodepoolStatusPhase = "CREATING"

	// ClusterNodepoolStatusPhaseREADY captures enum value "READY"
	ClusterNodepoolStatusPhaseREADY ClusterNodepoolStatusPhase = "READY"

	// ClusterNodepoolStatusPhaseERROR captures enum value "ERROR"
	ClusterNodepoolStatusPhaseERROR ClusterNodepoolStatusPhase = "ERROR"

	// ClusterNodepoolStatusPhaseDELETING captures enum value "DELETING"
	ClusterNodepoolStatusPhaseDELETING ClusterNodepoolStatusPhase = "DELETING"

	// ClusterNodepoolStatusPhaseRESIZING captures enum value "RESIZING"
	ClusterNodepoolStatusPhaseRESIZING ClusterNodepoolStatusPhase = "RESIZING"

	// ClusterNodepoolStatusPhaseUPGRADING captures enum value "UPGRADING"
	ClusterNodepoolStatusPhaseUPGRADING ClusterNodepoolStatusPhase = "UPGRADING"

	// ClusterNodepoolStatusPhaseUPGRADEFAILED captures enum value "UPGRADE_FAILED"
	ClusterNodepoolStatusPhaseUPGRADEFAILED ClusterNodepoolStatusPhase = "UPGRADE_FAILED"

	// ClusterNodepoolStatusPhaseWAITING captures enum value "WAITING"
	ClusterNodepoolStatusPhaseWAITING ClusterNodepoolStatusPhase = "WAITING"

	// ClusterNodepoolStatusPhaseUPDATING captures enum value "UPDATING"
	ClusterNodepoolStatusPhaseUPDATING ClusterNodepoolStatusPhase = "UPDATING"
)

// for schema
var clusterNodepoolStatusPhaseEnum []interface{}

func init() {
	var res []ClusterNodepoolStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATING","READY","ERROR","DELETING","RESIZING","UPGRADING","UPGRADE_FAILED","WAITING","UPDATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterNodepoolStatusPhaseEnum = append(clusterNodepoolStatusPhaseEnum, v)
	}
}

func (m ClusterNodepoolStatusPhase) validateClusterNodepoolStatusPhaseEnum(path, location string, value ClusterNodepoolStatusPhase) error {
	if err := validate.EnumCase(path, location, value, clusterNodepoolStatusPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster nodepool status phase
func (m ClusterNodepoolStatusPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterNodepoolStatusPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster nodepool status phase based on context it is used
func (m ClusterNodepoolStatusPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}