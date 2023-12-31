// Code generated by go-swagger; DO NOT EDIT.

package integration_resource_service

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

// NewIntegrationResourceServiceDeleteParams creates a new IntegrationResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIntegrationResourceServiceDeleteParams() *IntegrationResourceServiceDeleteParams {
	return &IntegrationResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationResourceServiceDeleteParamsWithTimeout creates a new IntegrationResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewIntegrationResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *IntegrationResourceServiceDeleteParams {
	return &IntegrationResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewIntegrationResourceServiceDeleteParamsWithContext creates a new IntegrationResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewIntegrationResourceServiceDeleteParamsWithContext(ctx context.Context) *IntegrationResourceServiceDeleteParams {
	return &IntegrationResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewIntegrationResourceServiceDeleteParamsWithHTTPClient creates a new IntegrationResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIntegrationResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *IntegrationResourceServiceDeleteParams {
	return &IntegrationResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
IntegrationResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the integration resource service delete operation.

	Typically these are written to a http.Request.
*/
type IntegrationResourceServiceDeleteParams struct {

	/* FullNameName.

	   Name of the integration.
	*/
	FullNameName string

	/* FullNameOrgID.

	   Org ID of the organization.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the integration resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationResourceServiceDeleteParams) WithDefaults() *IntegrationResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the integration resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *IntegrationResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) WithContext(ctx context.Context) *IntegrationResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *IntegrationResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) WithFullNameName(fullNameName string) *IntegrationResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *IntegrationResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the integration resource service delete params
func (o *IntegrationResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
