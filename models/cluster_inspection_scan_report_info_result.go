// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ClusterInspectionScanReportInfoResult Result describes the status of the inspection.
//
//   - RESULT_UNSPECIFIED: Unspecified result indicates the scan result is unknown.
//   - SUCCESS: Success - to be used if all the tests are part of the scan are successful.
//   - FAILURE: Failure - to indicate that one or more test as part of the scan failed.
//   - WARNING: Warning - to indicate that one or more test as part of the scan has a warning error.
//
// swagger:model cluster.inspection.scan.ReportInfo.Result
type ClusterInspectionScanReportInfoResult string

func NewClusterInspectionScanReportInfoResult(value ClusterInspectionScanReportInfoResult) *ClusterInspectionScanReportInfoResult {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterInspectionScanReportInfoResult.
func (m ClusterInspectionScanReportInfoResult) Pointer() *ClusterInspectionScanReportInfoResult {
	return &m
}

const (

	// ClusterInspectionScanReportInfoResultRESULTUNSPECIFIED captures enum value "RESULT_UNSPECIFIED"
	ClusterInspectionScanReportInfoResultRESULTUNSPECIFIED ClusterInspectionScanReportInfoResult = "RESULT_UNSPECIFIED"

	// ClusterInspectionScanReportInfoResultSUCCESS captures enum value "SUCCESS"
	ClusterInspectionScanReportInfoResultSUCCESS ClusterInspectionScanReportInfoResult = "SUCCESS"

	// ClusterInspectionScanReportInfoResultFAILURE captures enum value "FAILURE"
	ClusterInspectionScanReportInfoResultFAILURE ClusterInspectionScanReportInfoResult = "FAILURE"

	// ClusterInspectionScanReportInfoResultWARNING captures enum value "WARNING"
	ClusterInspectionScanReportInfoResultWARNING ClusterInspectionScanReportInfoResult = "WARNING"
)

// for schema
var clusterInspectionScanReportInfoResultEnum []interface{}

func init() {
	var res []ClusterInspectionScanReportInfoResult
	if err := json.Unmarshal([]byte(`["RESULT_UNSPECIFIED","SUCCESS","FAILURE","WARNING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterInspectionScanReportInfoResultEnum = append(clusterInspectionScanReportInfoResultEnum, v)
	}
}

func (m ClusterInspectionScanReportInfoResult) validateClusterInspectionScanReportInfoResultEnum(path, location string, value ClusterInspectionScanReportInfoResult) error {
	if err := validate.EnumCase(path, location, value, clusterInspectionScanReportInfoResultEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster inspection scan report info result
func (m ClusterInspectionScanReportInfoResult) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterInspectionScanReportInfoResultEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster inspection scan report info result based on context it is used
func (m ClusterInspectionScanReportInfoResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
