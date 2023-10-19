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

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// NewAuditResourceServiceCreateParams creates a new AuditResourceServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditResourceServiceCreateParams() *AuditResourceServiceCreateParams {
	return &AuditResourceServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditResourceServiceCreateParamsWithTimeout creates a new AuditResourceServiceCreateParams object
// with the ability to set a timeout on a request.
func NewAuditResourceServiceCreateParamsWithTimeout(timeout time.Duration) *AuditResourceServiceCreateParams {
	return &AuditResourceServiceCreateParams{
		timeout: timeout,
	}
}

// NewAuditResourceServiceCreateParamsWithContext creates a new AuditResourceServiceCreateParams object
// with the ability to set a context for a request.
func NewAuditResourceServiceCreateParamsWithContext(ctx context.Context) *AuditResourceServiceCreateParams {
	return &AuditResourceServiceCreateParams{
		Context: ctx,
	}
}

// NewAuditResourceServiceCreateParamsWithHTTPClient creates a new AuditResourceServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditResourceServiceCreateParamsWithHTTPClient(client *http.Client) *AuditResourceServiceCreateParams {
	return &AuditResourceServiceCreateParams{
		HTTPClient: client,
	}
}

/*
AuditResourceServiceCreateParams contains all the parameters to send to the API endpoint

	for the audit resource service create operation.

	Typically these are written to a http.Request.
*/
type AuditResourceServiceCreateParams struct {

	// Body.
	Body *models.AuditCreateAuditRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditResourceServiceCreateParams) WithDefaults() *AuditResourceServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditResourceServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the audit resource service create params
func (o *AuditResourceServiceCreateParams) WithTimeout(timeout time.Duration) *AuditResourceServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit resource service create params
func (o *AuditResourceServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit resource service create params
func (o *AuditResourceServiceCreateParams) WithContext(ctx context.Context) *AuditResourceServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit resource service create params
func (o *AuditResourceServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit resource service create params
func (o *AuditResourceServiceCreateParams) WithHTTPClient(client *http.Client) *AuditResourceServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit resource service create params
func (o *AuditResourceServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the audit resource service create params
func (o *AuditResourceServiceCreateParams) WithBody(body *models.AuditCreateAuditRequest) *AuditResourceServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the audit resource service create params
func (o *AuditResourceServiceCreateParams) SetBody(body *models.AuditCreateAuditRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuditResourceServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}