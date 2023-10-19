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

// AksclusterNodepoolTaintEffect The effect of the taint on pods
// that do not tolerate the taint.
// Valid effects are NoSchedule, NoExecute, PreferNoSchedule and EffectUnspecified.
//
//   - EFFECT_UNSPECIFIED: Unspecified effect.
//   - NO_SCHEDULE: Pods that do not tolerate this taint are not scheduled on the node.
//   - NO_EXECUTE: Pods are evicted from the node if are already running on the node.
//   - PREFER_NO_SCHEDULE: Avoids scheduling Pods that do not tolerate this taint onto the node.
//
// swagger:model akscluster.nodepool.Taint.Effect
type AksclusterNodepoolTaintEffect string

func NewAksclusterNodepoolTaintEffect(value AksclusterNodepoolTaintEffect) *AksclusterNodepoolTaintEffect {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AksclusterNodepoolTaintEffect.
func (m AksclusterNodepoolTaintEffect) Pointer() *AksclusterNodepoolTaintEffect {
	return &m
}

const (

	// AksclusterNodepoolTaintEffectEFFECTUNSPECIFIED captures enum value "EFFECT_UNSPECIFIED"
	AksclusterNodepoolTaintEffectEFFECTUNSPECIFIED AksclusterNodepoolTaintEffect = "EFFECT_UNSPECIFIED"

	// AksclusterNodepoolTaintEffectNOSCHEDULE captures enum value "NO_SCHEDULE"
	AksclusterNodepoolTaintEffectNOSCHEDULE AksclusterNodepoolTaintEffect = "NO_SCHEDULE"

	// AksclusterNodepoolTaintEffectNOEXECUTE captures enum value "NO_EXECUTE"
	AksclusterNodepoolTaintEffectNOEXECUTE AksclusterNodepoolTaintEffect = "NO_EXECUTE"

	// AksclusterNodepoolTaintEffectPREFERNOSCHEDULE captures enum value "PREFER_NO_SCHEDULE"
	AksclusterNodepoolTaintEffectPREFERNOSCHEDULE AksclusterNodepoolTaintEffect = "PREFER_NO_SCHEDULE"
)

// for schema
var aksclusterNodepoolTaintEffectEnum []interface{}

func init() {
	var res []AksclusterNodepoolTaintEffect
	if err := json.Unmarshal([]byte(`["EFFECT_UNSPECIFIED","NO_SCHEDULE","NO_EXECUTE","PREFER_NO_SCHEDULE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aksclusterNodepoolTaintEffectEnum = append(aksclusterNodepoolTaintEffectEnum, v)
	}
}

func (m AksclusterNodepoolTaintEffect) validateAksclusterNodepoolTaintEffectEnum(path, location string, value AksclusterNodepoolTaintEffect) error {
	if err := validate.EnumCase(path, location, value, aksclusterNodepoolTaintEffectEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this akscluster nodepool taint effect
func (m AksclusterNodepoolTaintEffect) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAksclusterNodepoolTaintEffectEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this akscluster nodepool taint effect based on context it is used
func (m AksclusterNodepoolTaintEffect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}