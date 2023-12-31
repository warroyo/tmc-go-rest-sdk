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

// ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse Response from getting an AntreaAgentInfo.
//
// swagger:model cluster.network.antrea.antreaagentinfo.GetAntreaAgentInfoResponse
type ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse struct {

	// AntreaAgentInfo returned.
	AntreaAgentInfo *ClusterNetworkAntreaAntreaagentinfoAntreaAgentInfo `json:"antreaAgentInfo,omitempty"`
}

// Validate validates this cluster network antrea antreaagentinfo get antrea agent info response
func (m *ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntreaAgentInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse) validateAntreaAgentInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AntreaAgentInfo) { // not required
		return nil
	}

	if m.AntreaAgentInfo != nil {
		if err := m.AntreaAgentInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antreaAgentInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("antreaAgentInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster network antrea antreaagentinfo get antrea agent info response based on the context it is used
func (m *ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntreaAgentInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse) contextValidateAntreaAgentInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AntreaAgentInfo != nil {
		if err := m.AntreaAgentInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antreaAgentInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("antreaAgentInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse) UnmarshalBinary(b []byte) error {
	var res ClusterNetworkAntreaAntreaagentinfoGetAntreaAgentInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
