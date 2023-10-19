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

// TanzupackageMetadataListMetadatasResponse Response from listing Metadatas.
//
// swagger:model tanzupackage.metadata.ListMetadatasResponse
type TanzupackageMetadataListMetadatasResponse struct {

	// List of metadatas.
	Metadatas []*TanzupackageMetadataMetadata `json:"metadatas"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this tanzupackage metadata list metadatas response
func (m *TanzupackageMetadataListMetadatasResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadatas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TanzupackageMetadataListMetadatasResponse) validateMetadatas(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadatas) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadatas); i++ {
		if swag.IsZero(m.Metadatas[i]) { // not required
			continue
		}

		if m.Metadatas[i] != nil {
			if err := m.Metadatas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadatas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tanzupackage metadata list metadatas response based on the context it is used
func (m *TanzupackageMetadataListMetadatasResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadatas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TanzupackageMetadataListMetadatasResponse) contextValidateMetadatas(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadatas); i++ {

		if m.Metadatas[i] != nil {
			if err := m.Metadatas[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadatas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TanzupackageMetadataListMetadatasResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TanzupackageMetadataListMetadatasResponse) UnmarshalBinary(b []byte) error {
	var res TanzupackageMetadataListMetadatasResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
