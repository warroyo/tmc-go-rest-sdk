// Code generated by go-swagger; DO NOT EDIT.

package audit_resource_service

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

// NewAuditResourceServiceDeleteParams creates a new AuditResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditResourceServiceDeleteParams() *AuditResourceServiceDeleteParams {
	return &AuditResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditResourceServiceDeleteParamsWithTimeout creates a new AuditResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewAuditResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *AuditResourceServiceDeleteParams {
	return &AuditResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewAuditResourceServiceDeleteParamsWithContext creates a new AuditResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewAuditResourceServiceDeleteParamsWithContext(ctx context.Context) *AuditResourceServiceDeleteParams {
	return &AuditResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewAuditResourceServiceDeleteParamsWithHTTPClient creates a new AuditResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *AuditResourceServiceDeleteParams {
	return &AuditResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
AuditResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the audit resource service delete operation.

	Typically these are written to a http.Request.
*/
type AuditResourceServiceDeleteParams struct {

	/* FullNameName.

	   Name of the audit run (id).
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of the organization.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditResourceServiceDeleteParams) WithDefaults() *AuditResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *AuditResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) WithContext(ctx context.Context) *AuditResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *AuditResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) WithFullNameName(fullNameName string) *AuditResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *AuditResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the audit resource service delete params
func (o *AuditResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *AuditResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
