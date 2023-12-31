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

// AksclusterPhase Phase of the cluster resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING: Resource is pending processing.
//   - CREATING: Resource is being created.
//   - READY: Resource is ready state.
//   - DELETING: Resource is being deleted.
//   - ERROR: Error in processing.
//   - UPDATING: This phase is used to reflect the UPDATING state of AKS cluster.
//   - OVER_LIMIT: This phase indicates cluster has crossed resource limits set for the organization.
//
// For such cluster we no longer sync data back to TMC.
//   - UPGRADING: This phase indicates kubernetes version is being upgraded for the cluster.
//   - STARTING: The AKS cluster is being started.
//   - STOPPING: The AKS cluster is being stopped.
//   - STOPPED: The AKS cluster is stopped.
//   - PENDING_MANAGE: This phase indicates the cluster is in the process of being managed by TMC.
//   - PENDING_UNMANAGE: This phase indicates the cluster is in the process of being unmanaged by TMC.
//
// swagger:model akscluster.Phase
type AksclusterPhase string

func NewAksclusterPhase(value AksclusterPhase) *AksclusterPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AksclusterPhase.
func (m AksclusterPhase) Pointer() *AksclusterPhase {
	return &m
}

const (

	// AksclusterPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	AksclusterPhasePHASEUNSPECIFIED AksclusterPhase = "PHASE_UNSPECIFIED"

	// AksclusterPhasePENDING captures enum value "PENDING"
	AksclusterPhasePENDING AksclusterPhase = "PENDING"

	// AksclusterPhaseCREATING captures enum value "CREATING"
	AksclusterPhaseCREATING AksclusterPhase = "CREATING"

	// AksclusterPhaseREADY captures enum value "READY"
	AksclusterPhaseREADY AksclusterPhase = "READY"

	// AksclusterPhaseDELETING captures enum value "DELETING"
	AksclusterPhaseDELETING AksclusterPhase = "DELETING"

	// AksclusterPhaseERROR captures enum value "ERROR"
	AksclusterPhaseERROR AksclusterPhase = "ERROR"

	// AksclusterPhaseUPDATING captures enum value "UPDATING"
	AksclusterPhaseUPDATING AksclusterPhase = "UPDATING"

	// AksclusterPhaseOVERLIMIT captures enum value "OVER_LIMIT"
	AksclusterPhaseOVERLIMIT AksclusterPhase = "OVER_LIMIT"

	// AksclusterPhaseUPGRADING captures enum value "UPGRADING"
	AksclusterPhaseUPGRADING AksclusterPhase = "UPGRADING"

	// AksclusterPhaseSTARTING captures enum value "STARTING"
	AksclusterPhaseSTARTING AksclusterPhase = "STARTING"

	// AksclusterPhaseSTOPPING captures enum value "STOPPING"
	AksclusterPhaseSTOPPING AksclusterPhase = "STOPPING"

	// AksclusterPhaseSTOPPED captures enum value "STOPPED"
	AksclusterPhaseSTOPPED AksclusterPhase = "STOPPED"

	// AksclusterPhasePENDINGMANAGE captures enum value "PENDING_MANAGE"
	AksclusterPhasePENDINGMANAGE AksclusterPhase = "PENDING_MANAGE"

	// AksclusterPhasePENDINGUNMANAGE captures enum value "PENDING_UNMANAGE"
	AksclusterPhasePENDINGUNMANAGE AksclusterPhase = "PENDING_UNMANAGE"
)

// for schema
var aksclusterPhaseEnum []interface{}

func init() {
	var res []AksclusterPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","READY","DELETING","ERROR","UPDATING","OVER_LIMIT","UPGRADING","STARTING","STOPPING","STOPPED","PENDING_MANAGE","PENDING_UNMANAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aksclusterPhaseEnum = append(aksclusterPhaseEnum, v)
	}
}

func (m AksclusterPhase) validateAksclusterPhaseEnum(path, location string, value AksclusterPhase) error {
	if err := validate.EnumCase(path, location, value, aksclusterPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this akscluster phase
func (m AksclusterPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAksclusterPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this akscluster phase based on context it is used
func (m AksclusterPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
