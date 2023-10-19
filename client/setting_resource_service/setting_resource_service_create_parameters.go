// Code generated by go-swagger; DO NOT EDIT.

package setting_resource_service

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

// NewSettingResourceServiceCreateParams creates a new SettingResourceServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingResourceServiceCreateParams() *SettingResourceServiceCreateParams {
	return &SettingResourceServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingResourceServiceCreateParamsWithTimeout creates a new SettingResourceServiceCreateParams object
// with the ability to set a timeout on a request.
func NewSettingResourceServiceCreateParamsWithTimeout(timeout time.Duration) *SettingResourceServiceCreateParams {
	return &SettingResourceServiceCreateParams{
		timeout: timeout,
	}
}

// NewSettingResourceServiceCreateParamsWithContext creates a new SettingResourceServiceCreateParams object
// with the ability to set a context for a request.
func NewSettingResourceServiceCreateParamsWithContext(ctx context.Context) *SettingResourceServiceCreateParams {
	return &SettingResourceServiceCreateParams{
		Context: ctx,
	}
}

// NewSettingResourceServiceCreateParamsWithHTTPClient creates a new SettingResourceServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingResourceServiceCreateParamsWithHTTPClient(client *http.Client) *SettingResourceServiceCreateParams {
	return &SettingResourceServiceCreateParams{
		HTTPClient: client,
	}
}

/*
SettingResourceServiceCreateParams contains all the parameters to send to the API endpoint

	for the setting resource service create operation.

	Typically these are written to a http.Request.
*/
type SettingResourceServiceCreateParams struct {

	// Body.
	Body *models.OrganizationSettingCreateSettingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the setting resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingResourceServiceCreateParams) WithDefaults() *SettingResourceServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the setting resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingResourceServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the setting resource service create params
func (o *SettingResourceServiceCreateParams) WithTimeout(timeout time.Duration) *SettingResourceServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the setting resource service create params
func (o *SettingResourceServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the setting resource service create params
func (o *SettingResourceServiceCreateParams) WithContext(ctx context.Context) *SettingResourceServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the setting resource service create params
func (o *SettingResourceServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the setting resource service create params
func (o *SettingResourceServiceCreateParams) WithHTTPClient(client *http.Client) *SettingResourceServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the setting resource service create params
func (o *SettingResourceServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the setting resource service create params
func (o *SettingResourceServiceCreateParams) WithBody(body *models.OrganizationSettingCreateSettingRequest) *SettingResourceServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the setting resource service create params
func (o *SettingResourceServiceCreateParams) SetBody(body *models.OrganizationSettingCreateSettingRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SettingResourceServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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