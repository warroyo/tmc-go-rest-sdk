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

// AksclusterClusterSKUName Name options of cluster SKU.
//
//   - NAME_UNSPECIFIED: Unspecified name.
//   - BASIC: Basic option for the AKS control plane. Deprecated in favor of BASE to match Azure API changes.
//   - BASE: Base option for the AKS control plane.
//
// swagger:model akscluster.ClusterSKU.Name
type AksclusterClusterSKUName string

func NewAksclusterClusterSKUName(value AksclusterClusterSKUName) *AksclusterClusterSKUName {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AksclusterClusterSKUName.
func (m AksclusterClusterSKUName) Pointer() *AksclusterClusterSKUName {
	return &m
}

const (

	// AksclusterClusterSKUNameNAMEUNSPECIFIED captures enum value "NAME_UNSPECIFIED"
	AksclusterClusterSKUNameNAMEUNSPECIFIED AksclusterClusterSKUName = "NAME_UNSPECIFIED"

	// AksclusterClusterSKUNameBASIC captures enum value "BASIC"
	AksclusterClusterSKUNameBASIC AksclusterClusterSKUName = "BASIC"

	// AksclusterClusterSKUNameBASE captures enum value "BASE"
	AksclusterClusterSKUNameBASE AksclusterClusterSKUName = "BASE"
)

// for schema
var aksclusterClusterSKUNameEnum []interface{}

func init() {
	var res []AksclusterClusterSKUName
	if err := json.Unmarshal([]byte(`["NAME_UNSPECIFIED","BASIC","BASE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aksclusterClusterSKUNameEnum = append(aksclusterClusterSKUNameEnum, v)
	}
}

func (m AksclusterClusterSKUName) validateAksclusterClusterSKUNameEnum(path, location string, value AksclusterClusterSKUName) error {
	if err := validate.EnumCase(path, location, value, aksclusterClusterSKUNameEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this akscluster cluster s k u name
func (m AksclusterClusterSKUName) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAksclusterClusterSKUNameEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this akscluster cluster s k u name based on context it is used
func (m AksclusterClusterSKUName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
