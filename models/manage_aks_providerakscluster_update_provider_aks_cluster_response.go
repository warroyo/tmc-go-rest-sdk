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

// ManageAksProvideraksclusterUpdateProviderAksClusterResponse Response from updating a ProviderAksCluster.
//
// swagger:model manage.aks.providerakscluster.UpdateProviderAksClusterResponse
type ManageAksProvideraksclusterUpdateProviderAksClusterResponse struct {

	// ProviderAksCluster updated.
	ProviderAksCluster *ManageAksProvideraksclusterProviderAksCluster `json:"providerAksCluster,omitempty"`
}

// Validate validates this manage aks providerakscluster update provider aks cluster response
func (m *ManageAksProvideraksclusterUpdateProviderAksClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviderAksCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManageAksProvideraksclusterUpdateProviderAksClusterResponse) validateProviderAksCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderAksCluster) { // not required
		return nil
	}

	if m.ProviderAksCluster != nil {
		if err := m.ProviderAksCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providerAksCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("providerAksCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this manage aks providerakscluster update provider aks cluster response based on the context it is used
func (m *ManageAksProvideraksclusterUpdateProviderAksClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProviderAksCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManageAksProvideraksclusterUpdateProviderAksClusterResponse) contextValidateProviderAksCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.ProviderAksCluster != nil {
		if err := m.ProviderAksCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providerAksCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("providerAksCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManageAksProvideraksclusterUpdateProviderAksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManageAksProvideraksclusterUpdateProviderAksClusterResponse) UnmarshalBinary(b []byte) error {
	var res ManageAksProvideraksclusterUpdateProviderAksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
