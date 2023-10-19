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

// ManagementclusterGetManagementClusterIAMPolicyResponse GetManagementClusterIAMPolicy response message.
//
// swagger:model managementcluster.GetManagementClusterIAMPolicyResponse
type ManagementclusterGetManagementClusterIAMPolicyResponse struct {

	// ManagementCluster policy.
	PolicyList []*VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policyList"`
}

// Validate validates this managementcluster get management cluster i a m policy response
func (m *ManagementclusterGetManagementClusterIAMPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterGetManagementClusterIAMPolicyResponse) validatePolicyList(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyList) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyList); i++ {
		if swag.IsZero(m.PolicyList[i]) { // not required
			continue
		}

		if m.PolicyList[i] != nil {
			if err := m.PolicyList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this managementcluster get management cluster i a m policy response based on the context it is used
func (m *ManagementclusterGetManagementClusterIAMPolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementclusterGetManagementClusterIAMPolicyResponse) contextValidatePolicyList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyList); i++ {

		if m.PolicyList[i] != nil {
			if err := m.PolicyList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementclusterGetManagementClusterIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementclusterGetManagementClusterIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res ManagementclusterGetManagementClusterIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}