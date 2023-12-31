// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTypeRecipeGetRecipeResponse Response from getting a Recipe.
//
// swagger:model policy.type.recipe.GetRecipeResponse
type PolicyTypeRecipeGetRecipeResponse struct {

	// Recipe returned.
	Recipe *PolicyTypeRecipeRecipe `json:"recipe,omitempty"`
}

// Validate validates this policy type recipe get recipe response
func (m *PolicyTypeRecipeGetRecipeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipe(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTypeRecipeGetRecipeResponse) validateRecipe(formats strfmt.Registry) error {
	if swag.IsZero(m.Recipe) { // not required
		return nil
	}

	if m.Recipe != nil {
		if err := m.Recipe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recipe")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy type recipe get recipe response based on the context it is used
func (m *PolicyTypeRecipeGetRecipeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecipe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTypeRecipeGetRecipeResponse) contextValidateRecipe(ctx context.Context, formats strfmt.Registry) error {

	if m.Recipe != nil {
		if err := m.Recipe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recipe")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTypeRecipeGetRecipeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTypeRecipeGetRecipeResponse) UnmarshalBinary(b []byte) error {
	var res PolicyTypeRecipeGetRecipeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
