// Code generated by go-swagger; DO NOT EDIT.

package schedule_resource_service

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

// NewScheduleResourceServiceDeleteParams creates a new ScheduleResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduleResourceServiceDeleteParams() *ScheduleResourceServiceDeleteParams {
	return &ScheduleResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleResourceServiceDeleteParamsWithTimeout creates a new ScheduleResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewScheduleResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *ScheduleResourceServiceDeleteParams {
	return &ScheduleResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewScheduleResourceServiceDeleteParamsWithContext creates a new ScheduleResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewScheduleResourceServiceDeleteParamsWithContext(ctx context.Context) *ScheduleResourceServiceDeleteParams {
	return &ScheduleResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewScheduleResourceServiceDeleteParamsWithHTTPClient creates a new ScheduleResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduleResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *ScheduleResourceServiceDeleteParams {
	return &ScheduleResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
ScheduleResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the schedule resource service delete operation.

	Typically these are written to a http.Request.
*/
type ScheduleResourceServiceDeleteParams struct {

	/* FullNameClusterName.

	   Name of Cluster.
	*/
	FullNameClusterName string

	/* FullNameManagementClusterName.

	   Name of ManagementCluster.
	*/
	FullNameManagementClusterName *string

	/* FullNameName.

	   Name of Inspection Schedule.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
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

// WithDefaults hydrates default values in the schedule resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleResourceServiceDeleteParams) WithDefaults() *ScheduleResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *ScheduleResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithContext(ctx context.Context) *ScheduleResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *ScheduleResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameClusterName adds the fullNameClusterName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithFullNameClusterName(fullNameClusterName string) *ScheduleResourceServiceDeleteParams {
	o.SetFullNameClusterName(fullNameClusterName)
	return o
}

// SetFullNameClusterName adds the fullNameClusterName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetFullNameClusterName(fullNameClusterName string) {
	o.FullNameClusterName = fullNameClusterName
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *ScheduleResourceServiceDeleteParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithFullNameName(fullNameName string) *ScheduleResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *ScheduleResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *ScheduleResourceServiceDeleteParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the schedule resource service delete params
func (o *ScheduleResourceServiceDeleteParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
