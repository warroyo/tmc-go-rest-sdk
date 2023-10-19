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

// ClusterInspectionScanGetScanResponse Response from getting a Scan.
//
// swagger:model cluster.inspection.scan.GetScanResponse
type ClusterInspectionScanGetScanResponse struct {

	// Scan returned.
	Scan *ClusterInspectionScanScan `json:"scan,omitempty"`
}

// Validate validates this cluster inspection scan get scan response
func (m *ClusterInspectionScanGetScanResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInspectionScanGetScanResponse) validateScan(formats strfmt.Registry) error {
	if swag.IsZero(m.Scan) { // not required
		return nil
	}

	if m.Scan != nil {
		if err := m.Scan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster inspection scan get scan response based on the context it is used
func (m *ClusterInspectionScanGetScanResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInspectionScanGetScanResponse) contextValidateScan(ctx context.Context, formats strfmt.Registry) error {

	if m.Scan != nil {
		if err := m.Scan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInspectionScanGetScanResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInspectionScanGetScanResponse) UnmarshalBinary(b []byte) error {
	var res ClusterInspectionScanGetScanResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
