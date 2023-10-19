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

// PolicyInsightCategory Defines category for the policy insight.
//
//   - CATEGORY_UNSPECIFIED: Unknown category for policy insight.
//   - ACCESS: Category for policy insight which is based on the access policy.
//   - CUSTOM: Category for policy insight which is based on the custom policy.
//   - IMAGE_REGISTRY: Category for policy insight which is based on the image policy.
//   - NETWORK: Category for policy insight which is based on the network policy.
//   - SECURITY: Category for policy insight which is based on the security policy.
//   - QUOTA: Category for policy insight which is based on the quota policy.
//   - MUTATION: Category for policy insight which is based on the quota policy.
//
// swagger:model policy.insight.Category
type PolicyInsightCategory string

func NewPolicyInsightCategory(value PolicyInsightCategory) *PolicyInsightCategory {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PolicyInsightCategory.
func (m PolicyInsightCategory) Pointer() *PolicyInsightCategory {
	return &m
}

const (

	// PolicyInsightCategoryCATEGORYUNSPECIFIED captures enum value "CATEGORY_UNSPECIFIED"
	PolicyInsightCategoryCATEGORYUNSPECIFIED PolicyInsightCategory = "CATEGORY_UNSPECIFIED"

	// PolicyInsightCategoryACCESS captures enum value "ACCESS"
	PolicyInsightCategoryACCESS PolicyInsightCategory = "ACCESS"

	// PolicyInsightCategoryCUSTOM captures enum value "CUSTOM"
	PolicyInsightCategoryCUSTOM PolicyInsightCategory = "CUSTOM"

	// PolicyInsightCategoryIMAGEREGISTRY captures enum value "IMAGE_REGISTRY"
	PolicyInsightCategoryIMAGEREGISTRY PolicyInsightCategory = "IMAGE_REGISTRY"

	// PolicyInsightCategoryNETWORK captures enum value "NETWORK"
	PolicyInsightCategoryNETWORK PolicyInsightCategory = "NETWORK"

	// PolicyInsightCategorySECURITY captures enum value "SECURITY"
	PolicyInsightCategorySECURITY PolicyInsightCategory = "SECURITY"

	// PolicyInsightCategoryQUOTA captures enum value "QUOTA"
	PolicyInsightCategoryQUOTA PolicyInsightCategory = "QUOTA"

	// PolicyInsightCategoryMUTATION captures enum value "MUTATION"
	PolicyInsightCategoryMUTATION PolicyInsightCategory = "MUTATION"
)

// for schema
var policyInsightCategoryEnum []interface{}

func init() {
	var res []PolicyInsightCategory
	if err := json.Unmarshal([]byte(`["CATEGORY_UNSPECIFIED","ACCESS","CUSTOM","IMAGE_REGISTRY","NETWORK","SECURITY","QUOTA","MUTATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyInsightCategoryEnum = append(policyInsightCategoryEnum, v)
	}
}

func (m PolicyInsightCategory) validatePolicyInsightCategoryEnum(path, location string, value PolicyInsightCategory) error {
	if err := validate.EnumCase(path, location, value, policyInsightCategoryEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy insight category
func (m PolicyInsightCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicyInsightCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this policy insight category based on context it is used
func (m PolicyInsightCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
