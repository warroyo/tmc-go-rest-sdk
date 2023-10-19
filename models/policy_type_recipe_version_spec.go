// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTypeRecipeVersionSpec Spec policy recipe version.
//
// swagger:model policy.type.recipe.version.Spec
type PolicyTypeRecipeVersionSpec struct {

	// Deprecated specifies whether this version of the recipe is deprecated.
	// Deprecated recipes will not be assignable to new policy instances nor visible in the UI.
	Deprecated bool `json:"deprecated,omitempty"`

	// InputSchema defines the set of variable inputs needed to create a policy using this recipe, in JsonSchema format.
	InputSchema string `json:"inputSchema,omitempty"`

	// Policy templates are references to kubernetes resources (policy pre-requisites) associated with this recipe.
	// These templates will be applied on clusters where policy instances using this recipe are effective.
	// A recipe can have 0 or more templates associated with it.
	PolicyTemplates []*VmwareTanzuCoreV1alpha1ObjectReference `json:"policyTemplates"`
}

// Validate validates this policy type recipe version spec
func (m *PolicyTypeRecipeVersionSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTypeRecipeVersionSpec) validatePolicyTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyTemplates); i++ {
		if swag.IsZero(m.PolicyTemplates[i]) { // not required
			continue
		}

		if m.PolicyTemplates[i] != nil {
			if err := m.PolicyTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this policy type recipe version spec based on the context it is used
func (m *PolicyTypeRecipeVersionSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTypeRecipeVersionSpec) contextValidatePolicyTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyTemplates); i++ {

		if m.PolicyTemplates[i] != nil {
			if err := m.PolicyTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTypeRecipeVersionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTypeRecipeVersionSpec) UnmarshalBinary(b []byte) error {
	var res PolicyTypeRecipeVersionSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
