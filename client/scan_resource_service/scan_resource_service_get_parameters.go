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

// NewScanResourceServiceGetParams creates a new ScanResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScanResourceServiceGetParams() *ScanResourceServiceGetParams {
	return &ScanResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScanResourceServiceGetParamsWithTimeout creates a new ScanResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewScanResourceServiceGetParamsWithTimeout(timeout time.Duration) *ScanResourceServiceGetParams {
	return &ScanResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewScanResourceServiceGetParamsWithContext creates a new ScanResourceServiceGetParams object
// with the ability to set a context for a request.
func NewScanResourceServiceGetParamsWithContext(ctx context.Context) *ScanResourceServiceGetParams {
	return &ScanResourceServiceGetParams{
		Context: ctx,
	}
}

// NewScanResourceServiceGetParamsWithHTTPClient creates a new ScanResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewScanResourceServiceGetParamsWithHTTPClient(client *http.Client) *ScanResourceServiceGetParams {
	return &ScanResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
ScanResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the scan resource service get operation.

	Typically these are written to a http.Request.
*/
type ScanResourceServiceGetParams struct {

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

// WithDefaults hydrates default values in the scan resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScanResourceServiceGetParams) WithDefaults() *ScanResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scan resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScanResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithTimeout(timeout time.Duration) *ScanResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithContext(ctx context.Context) *ScanResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithHTTPClient(client *http.Client) *ScanResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameClusterName adds the fullNameClusterName to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithFullNameClusterName(fullNameClusterName string) *ScanResourceServiceGetParams {
	o.SetFullNameClusterName(fullNameClusterName)
	return o
}

// SetFullNameClusterName adds the fullNameClusterName to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetFullNameClusterName(fullNameClusterName string) {
	o.FullNameClusterName = fullNameClusterName
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *ScanResourceServiceGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithFullNameName(fullNameName string) *ScanResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *ScanResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the scan resource service get params
func (o *ScanResourceServiceGetParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *ScanResourceServiceGetParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the scan resource service get params
func (o *ScanResourceServiceGetParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *ScanResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
