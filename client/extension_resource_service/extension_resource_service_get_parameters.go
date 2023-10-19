// Code generated by go-swagger; DO NOT EDIT.

package extension_resource_service

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

// NewExtensionResourceServiceGetParams creates a new ExtensionResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtensionResourceServiceGetParams() *ExtensionResourceServiceGetParams {
	return &ExtensionResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtensionResourceServiceGetParamsWithTimeout creates a new ExtensionResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewExtensionResourceServiceGetParamsWithTimeout(timeout time.Duration) *ExtensionResourceServiceGetParams {
	return &ExtensionResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewExtensionResourceServiceGetParamsWithContext creates a new ExtensionResourceServiceGetParams object
// with the ability to set a context for a request.
func NewExtensionResourceServiceGetParamsWithContext(ctx context.Context) *ExtensionResourceServiceGetParams {
	return &ExtensionResourceServiceGetParams{
		Context: ctx,
	}
}

// NewExtensionResourceServiceGetParamsWithHTTPClient creates a new ExtensionResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtensionResourceServiceGetParamsWithHTTPClient(client *http.Client) *ExtensionResourceServiceGetParams {
	return &ExtensionResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
ExtensionResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the extension resource service get operation.

	Typically these are written to a http.Request.
*/
type ExtensionResourceServiceGetParams struct {

	/* FullNameManagementClusterName.

	   Name of the management cluster.
	*/
	FullNameManagementClusterName string

	/* FullNameName.

	   Unique identifier of this Extension.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extension resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtensionResourceServiceGetParams) WithDefaults() *ExtensionResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extension resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtensionResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) WithTimeout(timeout time.Duration) *ExtensionResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) WithContext(ctx context.Context) *ExtensionResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) WithHTTPClient(client *http.Client) *ExtensionResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *ExtensionResourceServiceGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) WithFullNameName(fullNameName string) *ExtensionResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *ExtensionResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the extension resource service get params
func (o *ExtensionResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ExtensionResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fullName.managementClusterName
	if err := r.SetPathParam("fullName.managementClusterName", o.FullNameManagementClusterName); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}