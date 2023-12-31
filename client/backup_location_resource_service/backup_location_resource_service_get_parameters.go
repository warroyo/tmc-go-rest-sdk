// Code generated by go-swagger; DO NOT EDIT.

package backup_location_resource_service

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

// NewBackupLocationResourceServiceGetParams creates a new BackupLocationResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupLocationResourceServiceGetParams() *BackupLocationResourceServiceGetParams {
	return &BackupLocationResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupLocationResourceServiceGetParamsWithTimeout creates a new BackupLocationResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewBackupLocationResourceServiceGetParamsWithTimeout(timeout time.Duration) *BackupLocationResourceServiceGetParams {
	return &BackupLocationResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewBackupLocationResourceServiceGetParamsWithContext creates a new BackupLocationResourceServiceGetParams object
// with the ability to set a context for a request.
func NewBackupLocationResourceServiceGetParamsWithContext(ctx context.Context) *BackupLocationResourceServiceGetParams {
	return &BackupLocationResourceServiceGetParams{
		Context: ctx,
	}
}

// NewBackupLocationResourceServiceGetParamsWithHTTPClient creates a new BackupLocationResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupLocationResourceServiceGetParamsWithHTTPClient(client *http.Client) *BackupLocationResourceServiceGetParams {
	return &BackupLocationResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
BackupLocationResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the backup location resource service get operation.

	Typically these are written to a http.Request.
*/
type BackupLocationResourceServiceGetParams struct {

	/* FullNameName.

	   Name of the Backup Location.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameProviderName.

	   Name of the data protection provider name.
	*/
	FullNameProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup location resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupLocationResourceServiceGetParams) WithDefaults() *BackupLocationResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup location resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupLocationResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) WithTimeout(timeout time.Duration) *BackupLocationResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) WithContext(ctx context.Context) *BackupLocationResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) WithHTTPClient(client *http.Client) *BackupLocationResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) WithFullNameName(fullNameName string) *BackupLocationResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *BackupLocationResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProviderName adds the fullNameProviderName to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) WithFullNameProviderName(fullNameProviderName string) *BackupLocationResourceServiceGetParams {
	o.SetFullNameProviderName(fullNameProviderName)
	return o
}

// SetFullNameProviderName adds the fullNameProviderName to the backup location resource service get params
func (o *BackupLocationResourceServiceGetParams) SetFullNameProviderName(fullNameProviderName string) {
	o.FullNameProviderName = fullNameProviderName
}

// WriteToRequest writes these params to a swagger request
func (o *BackupLocationResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param fullName.providerName
	if err := r.SetPathParam("fullName.providerName", o.FullNameProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
