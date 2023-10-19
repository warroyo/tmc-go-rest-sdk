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

// ClusterExtensionListExtensionsResponse Response from listing Extensions.
//
// swagger:model cluster.extension.ListExtensionsResponse
type ClusterExtensionListExtensionsResponse struct {

	// List of extensions.
	Extensions []*ClusterExtensionExtension `json:"extensions"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this cluster extension list extensions response
func (m *ClusterExtensionListExtensionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterExtensionListExtensionsResponse) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster extension list extensions response based on the context it is used
func (m *ClusterExtensionListExtensionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterExtensionListExtensionsResponse) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Extensions); i++ {

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterExtensionListExtensionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterExtensionListExtensionsResponse) UnmarshalBinary(b []byte) error {
	var res ClusterExtensionListExtensionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}