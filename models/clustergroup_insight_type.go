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

// ClustergroupInsightType Type specifies the different types of insights.
//
//   - TYPE_UNSPECIFIED: Unknown type.
//   - ERROR: Insight about an error scenario.
//   - OVERRIDE: Insight about override.
//
// For example: source resource on cluster-group was overridden on member cluster by another resource.
//
// swagger:model clustergroup.insight.Type
type ClustergroupInsightType string

func NewClustergroupInsightType(value ClustergroupInsightType) *ClustergroupInsightType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClustergroupInsightType.
func (m ClustergroupInsightType) Pointer() *ClustergroupInsightType {
	return &m
}

const (

	// ClustergroupInsightTypeTYPEUNSPECIFIED captures enum value "TYPE_UNSPECIFIED"
	ClustergroupInsightTypeTYPEUNSPECIFIED ClustergroupInsightType = "TYPE_UNSPECIFIED"

	// ClustergroupInsightTypeERROR captures enum value "ERROR"
	ClustergroupInsightTypeERROR ClustergroupInsightType = "ERROR"

	// ClustergroupInsightTypeOVERRIDE captures enum value "OVERRIDE"
	ClustergroupInsightTypeOVERRIDE ClustergroupInsightType = "OVERRIDE"
)

// for schema
var clustergroupInsightTypeEnum []interface{}

func init() {
	var res []ClustergroupInsightType
	if err := json.Unmarshal([]byte(`["TYPE_UNSPECIFIED","ERROR","OVERRIDE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clustergroupInsightTypeEnum = append(clustergroupInsightTypeEnum, v)
	}
}

func (m ClustergroupInsightType) validateClustergroupInsightTypeEnum(path, location string, value ClustergroupInsightType) error {
	if err := validate.EnumCase(path, location, value, clustergroupInsightTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this clustergroup insight type
func (m ClustergroupInsightType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClustergroupInsightTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this clustergroup insight type based on context it is used
func (m ClustergroupInsightType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
