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

// PolicyTypeRecipeVersionEnumerateVersionsResponse Streamed response from enumerating Versions.
//
// swagger:model policy.type.recipe.version.EnumerateVersionsResponse
type PolicyTypeRecipeVersionEnumerateVersionsResponse struct {

	// Version object.
	Version *PolicyTypeRecipeVersionVersion `json:"version,omitempty"`
}

// Validate validates this policy type recipe version enumerate versions response
func (m *PolicyTypeRecipeVersionEnumerateVersionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTypeRecipeVersionEnumerateVersionsResponse) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy type recipe version enumerate versions response based on the context it is used
func (m *PolicyTypeRecipeVersionEnumerateVersionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTypeRecipeVersionEnumerateVersionsResponse) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTypeRecipeVersionEnumerateVersionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTypeRecipeVersionEnumerateVersionsResponse) UnmarshalBinary(b []byte) error {
	var res PolicyTypeRecipeVersionEnumerateVersionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}