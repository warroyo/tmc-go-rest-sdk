// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterInspectionScanProgressInfo Progress information about the inspection scan.
//
// swagger:model cluster.inspection.scan.ProgressInfo
type ClusterInspectionScanProgressInfo struct {

	// Number of tests completed.
	NumTestsCompleted string `json:"numTestsCompleted,omitempty"`

	// Number of tests run as part of the inspection scan.
	TotalNumTests string `json:"totalNumTests,omitempty"`
}

// Validate validates this cluster inspection scan progress info
func (m *ClusterInspectionScanProgressInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster inspection scan progress info based on context it is used
func (m *ClusterInspectionScanProgressInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInspectionScanProgressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInspectionScanProgressInfo) UnmarshalBinary(b []byte) error {
	var res ClusterInspectionScanProgressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
