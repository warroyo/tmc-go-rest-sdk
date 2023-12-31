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

// AksclusterClusterSKU The SKU of an AKS cluster
//
// swagger:model akscluster.ClusterSKU
type AksclusterClusterSKU struct {

	// Name of the cluster SKU.
	Name *AksclusterClusterSKUName `json:"name,omitempty"`

	// Tier of the cluster SKU.
	Tier *AksclusterTier `json:"tier,omitempty"`
}

// Validate validates this akscluster cluster s k u
func (m *AksclusterClusterSKU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterClusterSKU) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterSKU) validateTier(formats strfmt.Registry) error {
	if swag.IsZero(m.Tier) { // not required
		return nil
	}

	if m.Tier != nil {
		if err := m.Tier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tier")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this akscluster cluster s k u based on the context it is used
func (m *AksclusterClusterSKU) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterClusterSKU) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {
		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *AksclusterClusterSKU) contextValidateTier(ctx context.Context, formats strfmt.Registry) error {

	if m.Tier != nil {
		if err := m.Tier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AksclusterClusterSKU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterClusterSKU) UnmarshalBinary(b []byte) error {
	var res AksclusterClusterSKU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
