// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterAuthTokenGetResponse Response containing an opaque bearer token for in-cluster access.
//
// swagger:model cluster.auth.token.GetResponse
type ClusterAuthTokenGetResponse struct {

	// ClusterUID is the UID for the cluster.
	ClusterUID string `json:"clusterUid,omitempty"`

	// The expiration timestamp for the token.
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// Cluster full_name.
	FullName *ClusterFullName `json:"fullName,omitempty"`

	// Opaque bearer token.
	Token string `json:"token,omitempty"`
}

// Validate validates this cluster auth token get response
func (m *ClusterAuthTokenGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterAuthTokenGetResponse) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterAuthTokenGetResponse) validateFullName(formats strfmt.Registry) error {
	if swag.IsZero(m.FullName) { // not required
		return nil
	}

	if m.FullName != nil {
		if err := m.FullName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster auth token get response based on the context it is used
func (m *ClusterAuthTokenGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterAuthTokenGetResponse) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

	if m.FullName != nil {
		if err := m.FullName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterAuthTokenGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterAuthTokenGetResponse) UnmarshalBinary(b []byte) error {
	var res ClusterAuthTokenGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
