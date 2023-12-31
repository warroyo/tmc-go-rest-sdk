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

// SubscriptionTier The tier of TMC that the Org has signed up for.
//
//   - TIER_UNSPECIFIED: UNSPECIFIED Tier.
//   - TMC_STANDARD: Standard Tier.
//   - TMC_ADVANCED: Advanced Tier.
//   - TMC_ENTERPRISE: Enterprise Tier.
//   - TMC_ESSENTIALS: Essentials Tier.
//   - TMC_STARTER: STARTER Tier.
//
// swagger:model subscription.Tier
type SubscriptionTier string

func NewSubscriptionTier(value SubscriptionTier) *SubscriptionTier {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SubscriptionTier.
func (m SubscriptionTier) Pointer() *SubscriptionTier {
	return &m
}

const (

	// SubscriptionTierTIERUNSPECIFIED captures enum value "TIER_UNSPECIFIED"
	SubscriptionTierTIERUNSPECIFIED SubscriptionTier = "TIER_UNSPECIFIED"

	// SubscriptionTierTMCSTANDARD captures enum value "TMC_STANDARD"
	SubscriptionTierTMCSTANDARD SubscriptionTier = "TMC_STANDARD"

	// SubscriptionTierTMCADVANCED captures enum value "TMC_ADVANCED"
	SubscriptionTierTMCADVANCED SubscriptionTier = "TMC_ADVANCED"

	// SubscriptionTierTMCENTERPRISE captures enum value "TMC_ENTERPRISE"
	SubscriptionTierTMCENTERPRISE SubscriptionTier = "TMC_ENTERPRISE"

	// SubscriptionTierTMCESSENTIALS captures enum value "TMC_ESSENTIALS"
	SubscriptionTierTMCESSENTIALS SubscriptionTier = "TMC_ESSENTIALS"

	// SubscriptionTierTMCSTARTER captures enum value "TMC_STARTER"
	SubscriptionTierTMCSTARTER SubscriptionTier = "TMC_STARTER"
)

// for schema
var subscriptionTierEnum []interface{}

func init() {
	var res []SubscriptionTier
	if err := json.Unmarshal([]byte(`["TIER_UNSPECIFIED","TMC_STANDARD","TMC_ADVANCED","TMC_ENTERPRISE","TMC_ESSENTIALS","TMC_STARTER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTierEnum = append(subscriptionTierEnum, v)
	}
}

func (m SubscriptionTier) validateSubscriptionTierEnum(path, location string, value SubscriptionTier) error {
	if err := validate.EnumCase(path, location, value, subscriptionTierEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this subscription tier
func (m SubscriptionTier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSubscriptionTierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this subscription tier based on context it is used
func (m SubscriptionTier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
