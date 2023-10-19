// Code generated by go-swagger; DO NOT EDIT.

package core_cli_service

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

// NewCoreCliServiceGetParams creates a new CoreCliServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCoreCliServiceGetParams() *CoreCliServiceGetParams {
	return &CoreCliServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCoreCliServiceGetParamsWithTimeout creates a new CoreCliServiceGetParams object
// with the ability to set a timeout on a request.
func NewCoreCliServiceGetParamsWithTimeout(timeout time.Duration) *CoreCliServiceGetParams {
	return &CoreCliServiceGetParams{
		timeout: timeout,
	}
}

// NewCoreCliServiceGetParamsWithContext creates a new CoreCliServiceGetParams object
// with the ability to set a context for a request.
func NewCoreCliServiceGetParamsWithContext(ctx context.Context) *CoreCliServiceGetParams {
	return &CoreCliServiceGetParams{
		Context: ctx,
	}
}

// NewCoreCliServiceGetParamsWithHTTPClient creates a new CoreCliServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCoreCliServiceGetParamsWithHTTPClient(client *http.Client) *CoreCliServiceGetParams {
	return &CoreCliServiceGetParams{
		HTTPClient: client,
	}
}

/*
CoreCliServiceGetParams contains all the parameters to send to the API endpoint

	for the core cli service get operation.

	Typically these are written to a http.Request.
*/
type CoreCliServiceGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the core cli service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CoreCliServiceGetParams) WithDefaults() *CoreCliServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the core cli service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CoreCliServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the core cli service get params
func (o *CoreCliServiceGetParams) WithTimeout(timeout time.Duration) *CoreCliServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the core cli service get params
func (o *CoreCliServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the core cli service get params
func (o *CoreCliServiceGetParams) WithContext(ctx context.Context) *CoreCliServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the core cli service get params
func (o *CoreCliServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the core cli service get params
func (o *CoreCliServiceGetParams) WithHTTPClient(client *http.Client) *CoreCliServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the core cli service get params
func (o *CoreCliServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CoreCliServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
