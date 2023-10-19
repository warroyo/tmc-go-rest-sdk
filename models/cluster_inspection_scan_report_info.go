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

// ClusterInspectionScanReportInfo Contains the metadata for a single report
// (e.g. report id, etc).
//
// swagger:model cluster.inspection.scan.ReportInfo
type ClusterInspectionScanReportInfo struct {

	// Kubernetes Server Version.
	KubeServerVersion string `json:"kubeServerVersion,omitempty"`

	// Number of inspections failed.
	NumFailed string `json:"numFailed,omitempty"`

	// Total number of inspections as part of the scan.
	NumInspections string `json:"numInspections,omitempty"`

	// Number of inspection tests in warning state.
	NumWarning string `json:"numWarning,omitempty"`

	// Progress information about the inspection scan.
	ProgressInfo *ClusterInspectionScanProgressInfo `json:"progressInfo,omitempty"`

	// Internal ID of the run.
	ReportID string `json:"reportId,omitempty"`

	// Result is a success / failure condition based on the result of the scan.
	Result *ClusterInspectionScanReportInfoResult `json:"result,omitempty"`

	// Date and time of the run.
	// Format: date-time
	RunDatetime strfmt.DateTime `json:"runDatetime,omitempty"`

	// The scan type.
	ScanType string `json:"scanType,omitempty"`
}

// Validate validates this cluster inspection scan report info
func (m *ClusterInspectionScanReportInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProgressInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInspectionScanReportInfo) validateProgressInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ProgressInfo) { // not required
		return nil
	}

	if m.ProgressInfo != nil {
		if err := m.ProgressInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progressInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progressInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInspectionScanReportInfo) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInspectionScanReportInfo) validateRunDatetime(formats strfmt.Registry) error {
	if swag.IsZero(m.RunDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("runDatetime", "body", "date-time", m.RunDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster inspection scan report info based on the context it is used
func (m *ClusterInspectionScanReportInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProgressInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInspectionScanReportInfo) contextValidateProgressInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ProgressInfo != nil {
		if err := m.ProgressInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progressInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progressInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInspectionScanReportInfo) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {
		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInspectionScanReportInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInspectionScanReportInfo) UnmarshalBinary(b []byte) error {
	var res ClusterInspectionScanReportInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}