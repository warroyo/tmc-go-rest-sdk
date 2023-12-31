// Code generated by go-swagger; DO NOT EDIT.

package usage_service

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

// NewUsageServiceGetOrgLevelUsageParams creates a new UsageServiceGetOrgLevelUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsageServiceGetOrgLevelUsageParams() *UsageServiceGetOrgLevelUsageParams {
	return &UsageServiceGetOrgLevelUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsageServiceGetOrgLevelUsageParamsWithTimeout creates a new UsageServiceGetOrgLevelUsageParams object
// with the ability to set a timeout on a request.
func NewUsageServiceGetOrgLevelUsageParamsWithTimeout(timeout time.Duration) *UsageServiceGetOrgLevelUsageParams {
	return &UsageServiceGetOrgLevelUsageParams{
		timeout: timeout,
	}
}

// NewUsageServiceGetOrgLevelUsageParamsWithContext creates a new UsageServiceGetOrgLevelUsageParams object
// with the ability to set a context for a request.
func NewUsageServiceGetOrgLevelUsageParamsWithContext(ctx context.Context) *UsageServiceGetOrgLevelUsageParams {
	return &UsageServiceGetOrgLevelUsageParams{
		Context: ctx,
	}
}

// NewUsageServiceGetOrgLevelUsageParamsWithHTTPClient creates a new UsageServiceGetOrgLevelUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsageServiceGetOrgLevelUsageParamsWithHTTPClient(client *http.Client) *UsageServiceGetOrgLevelUsageParams {
	return &UsageServiceGetOrgLevelUsageParams{
		HTTPClient: client,
	}
}

/*
UsageServiceGetOrgLevelUsageParams contains all the parameters to send to the API endpoint

	for the usage service get org level usage operation.

	Typically these are written to a http.Request.
*/
type UsageServiceGetOrgLevelUsageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the usage service get org level usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsageServiceGetOrgLevelUsageParams) WithDefaults() *UsageServiceGetOrgLevelUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the usage service get org level usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsageServiceGetOrgLevelUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the usage service get org level usage params
func (o *UsageServiceGetOrgLevelUsageParams) WithTimeout(timeout time.Duration) *UsageServiceGetOrgLevelUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the usage service get org level usage params
func (o *UsageServiceGetOrgLevelUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the usage service get org level usage params
func (o *UsageServiceGetOrgLevelUsageParams) WithContext(ctx context.Context) *UsageServiceGetOrgLevelUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the usage service get org level usage params
func (o *UsageServiceGetOrgLevelUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the usage service get org level usage params
func (o *UsageServiceGetOrgLevelUsageParams) WithHTTPClient(client *http.Client) *UsageServiceGetOrgLevelUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the usage service get org level usage params
func (o *UsageServiceGetOrgLevelUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UsageServiceGetOrgLevelUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
