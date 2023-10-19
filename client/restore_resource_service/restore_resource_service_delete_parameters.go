// Code generated by go-swagger; DO NOT EDIT.

package restore_resource_service

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

// NewRestoreResourceServiceDeleteParams creates a new RestoreResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreResourceServiceDeleteParams() *RestoreResourceServiceDeleteParams {
	return &RestoreResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreResourceServiceDeleteParamsWithTimeout creates a new RestoreResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewRestoreResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *RestoreResourceServiceDeleteParams {
	return &RestoreResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewRestoreResourceServiceDeleteParamsWithContext creates a new RestoreResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewRestoreResourceServiceDeleteParamsWithContext(ctx context.Context) *RestoreResourceServiceDeleteParams {
	return &RestoreResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewRestoreResourceServiceDeleteParamsWithHTTPClient creates a new RestoreResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *RestoreResourceServiceDeleteParams {
	return &RestoreResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
RestoreResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the restore resource service delete operation.

	Typically these are written to a http.Request.
*/
type RestoreResourceServiceDeleteParams struct {

	/* FullNameClusterName.

	   Name of Cluster.
	*/
	FullNameClusterName string

	/* FullNameManagementClusterName.

	   Name of management cluster.
	*/
	FullNameManagementClusterName *string

	/* FullNameName.

	   Name of this Restore.
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

// WithDefaults hydrates default values in the restore resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreResourceServiceDeleteParams) WithDefaults() *RestoreResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *RestoreResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithContext(ctx context.Context) *RestoreResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *RestoreResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameClusterName adds the fullNameClusterName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithFullNameClusterName(fullNameClusterName string) *RestoreResourceServiceDeleteParams {
	o.SetFullNameClusterName(fullNameClusterName)
	return o
}

// SetFullNameClusterName adds the fullNameClusterName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetFullNameClusterName(fullNameClusterName string) {
	o.FullNameClusterName = fullNameClusterName
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *RestoreResourceServiceDeleteParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithFullNameName(fullNameName string) *RestoreResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *RestoreResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *RestoreResourceServiceDeleteParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the restore resource service delete params
func (o *RestoreResourceServiceDeleteParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
