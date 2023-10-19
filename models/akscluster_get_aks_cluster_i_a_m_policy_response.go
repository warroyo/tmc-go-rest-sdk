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

// AksclusterGetAksClusterIAMPolicyResponse GetAksClusterIAMPolicy response message.
//
// swagger:model akscluster.GetAksClusterIAMPolicyResponse
type AksclusterGetAksClusterIAMPolicyResponse struct {

	// AksCluster policy.
	PolicyList []*VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policyList"`
}

// Validate validates this akscluster get aks cluster i a m policy response
func (m *AksclusterGetAksClusterIAMPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterGetAksClusterIAMPolicyResponse) validatePolicyList(formats strfmt.Registry) error {
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

// ContextValidate validate this akscluster get aks cluster i a m policy response based on the context it is used
func (m *AksclusterGetAksClusterIAMPolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AksclusterGetAksClusterIAMPolicyResponse) contextValidatePolicyList(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AksclusterGetAksClusterIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AksclusterGetAksClusterIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res AksclusterGetAksClusterIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}