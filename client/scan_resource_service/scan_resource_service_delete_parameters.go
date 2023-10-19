// Code generated by go-swagger; DO NOT EDIT.

package scan_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewScanResourceServiceDeleteParams creates a new ScanResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScanResourceServiceDeleteParams() *ScanResourceServiceDeleteParams {
	return &ScanResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScanResourceServiceDeleteParamsWithTimeout creates a new ScanResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewScanResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *ScanResourceServiceDeleteParams {
	return &ScanResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewScanResourceServiceDeleteParamsWithContext creates a new ScanResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewScanResourceServiceDeleteParamsWithContext(ctx context.Context) *ScanResourceServiceDeleteParams {
	return &ScanResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewScanResourceServiceDeleteParamsWithHTTPClient creates a new ScanResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewScanResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *ScanResourceServiceDeleteParams {
	return &ScanResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
ScanResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the scan resource service delete operation.

	Typically these are written to a http.Request.
*/
type ScanResourceServiceDeleteParams struct {

	/* FullNameClusterName.

	   Name of Cluster.
	*/
	FullNameClusterName string

	/* FullNameManagementClusterName.

	   Name of management cluster.
	*/
	FullNameManagementClusterName *string

	/* FullNameName.

	   Name of the inspection scan.
	*/
	FullNameName string

	/* FullNameOrgID.

	   Org ID of the inspection scan.
	*/
	FullNameOrgID *string

	/* FullNameProvisionerName.

	   Name of Provisioner.
	*/
	FullNameProvisionerName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scan resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScanResourceServiceDeleteParams) WithDefaults() *ScanResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scan resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScanResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *ScanResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithContext(ctx context.Context) *ScanResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *ScanResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameClusterName adds the fullNameClusterName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithFullNameClusterName(fullNameClusterName string) *ScanResourceServiceDeleteParams {
	o.SetFullNameClusterName(fullNameClusterName)
	return o
}

// SetFullNameClusterName adds the fullNameClusterName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetFullNameClusterName(fullNameClusterName string) {
	o.FullNameClusterName = fullNameClusterName
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *ScanResourceServiceDeleteParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithFullNameName(fullNameName string) *ScanResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *ScanResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *ScanResourceServiceDeleteParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the scan resource service delete params
func (o *ScanResourceServiceDeleteParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *ScanResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fullName.clusterName
	if err := r.SetPathParam("fullName.clusterName", o.FullNameClusterName); err != nil {
		return err
	}

	if o.FullNameManagementClusterName != nil {

		// query param fullName.managementClusterName
		var qrFullNameManagementClusterName string

		if o.FullNameManagementClusterName != nil {
			qrFullNameManagementClusterName = *o.FullNameManagementClusterName
		}
		qFullNameManagementClusterName := qrFullNameManagementClusterName
		if qFullNameManagementClusterName != "" {

			if err := r.SetQueryParam("fullName.managementClusterName", qFullNameManagementClusterName); err != nil {
				return err
			}
		}
	}

	// path param fullName.name
	if err := r.SetPathParam("fullName.name", o.FullNameName); err != nil {
		return err
	}

	if o.FullNameOrgID != nil {

		// query param fullName.orgId
		var qrFullNameOrgID string

		if o.FullNameOrgID != nil {
			qrFullNameOrgID = *o.FullNameOrgID
		}
		qFullNameOrgID := qrFullNameOrgID
		if qFullNameOrgID != "" {

			if err := r.SetQueryParam("fullName.orgId", qFullNameOrgID); err != nil {
				return err
			}
		}
	}

	if o.FullNameProvisionerName != nil {

		// query param fullName.provisionerName
		var qrFullNameProvisionerName string

		if o.FullNameProvisionerName != nil {
			qrFullNameProvisionerName = *o.FullNameProvisionerName
		}
		qFullNameProvisionerName := qrFullNameProvisionerName
		if qFullNameProvisionerName != "" {

			if err := r.SetQueryParam("fullName.provisionerName", qFullNameProvisionerName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}