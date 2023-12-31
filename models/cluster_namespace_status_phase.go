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

// ClusterNamespaceStatusPhase The overall phase of the namespace.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - CREATING: Creating phase is set when the namespace is being created.
//   - ATTACHING: Attaching phase is set when the namespace is being attached.
//   - UPDATING: Updating phase is set when the namespace is being updated.
//   - READY: Ready phase is set when the namespace is successfully created/attached/updated.
//   - ERROR: Error phase is set when there was a failure while creating/attaching/updating the namespace.
//
// swagger:model cluster.namespace.Status.Phase
type ClusterNamespaceStatusPhase string

func NewClusterNamespaceStatusPhase(value ClusterNamespaceStatusPhase) *ClusterNamespaceStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterNamespaceStatusPhase.
func (m ClusterNamespaceStatusPhase) Pointer() *ClusterNamespaceStatusPhase {
	return &m
}

const (

	// ClusterNamespaceStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	ClusterNamespaceStatusPhasePHASEUNSPECIFIED ClusterNamespaceStatusPhase = "PHASE_UNSPECIFIED"

	// ClusterNamespaceStatusPhaseCREATING captures enum value "CREATING"
	ClusterNamespaceStatusPhaseCREATING ClusterNamespaceStatusPhase = "CREATING"

	// ClusterNamespaceStatusPhaseATTACHING captures enum value "ATTACHING"
	ClusterNamespaceStatusPhaseATTACHING ClusterNamespaceStatusPhase = "ATTACHING"

	// ClusterNamespaceStatusPhaseUPDATING captures enum value "UPDATING"
	ClusterNamespaceStatusPhaseUPDATING ClusterNamespaceStatusPhase = "UPDATING"

	// ClusterNamespaceStatusPhaseREADY captures enum value "READY"
	ClusterNamespaceStatusPhaseREADY ClusterNamespaceStatusPhase = "READY"

	// ClusterNamespaceStatusPhaseERROR captures enum value "ERROR"
	ClusterNamespaceStatusPhaseERROR ClusterNamespaceStatusPhase = "ERROR"
)

// for schema
var clusterNamespaceStatusPhaseEnum []interface{}

func init() {
	var res []ClusterNamespaceStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATING","ATTACHING","UPDATING","READY","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterNamespaceStatusPhaseEnum = append(clusterNamespaceStatusPhaseEnum, v)
	}
}

func (m ClusterNamespaceStatusPhase) validateClusterNamespaceStatusPhaseEnum(path, location string, value ClusterNamespaceStatusPhase) error {
	if err := validate.EnumCase(path, location, value, clusterNamespaceStatusPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster namespace status phase
func (m ClusterNamespaceStatusPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterNamespaceStatusPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster namespace status phase based on context it is used
func (m ClusterNamespaceStatusPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
