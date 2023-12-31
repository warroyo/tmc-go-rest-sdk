// Code generated by go-swagger; DO NOT EDIT.

package eula_resource_service

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

// NewEulaResourceServiceValidateParams creates a new EulaResourceServiceValidateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEulaResourceServiceValidateParams() *EulaResourceServiceValidateParams {
	return &EulaResourceServiceValidateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEulaResourceServiceValidateParamsWithTimeout creates a new EulaResourceServiceValidateParams object
// with the ability to set a timeout on a request.
func NewEulaResourceServiceValidateParamsWithTimeout(timeout time.Duration) *EulaResourceServiceValidateParams {
	return &EulaResourceServiceValidateParams{
		timeout: timeout,
	}
}

// NewEulaResourceServiceValidateParamsWithContext creates a new EulaResourceServiceValidateParams object
// with the ability to set a context for a request.
func NewEulaResourceServiceValidateParamsWithContext(ctx context.Context) *EulaResourceServiceValidateParams {
	return &EulaResourceServiceValidateParams{
		Context: ctx,
	}
}

// NewEulaResourceServiceValidateParamsWithHTTPClient creates a new EulaResourceServiceValidateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEulaResourceServiceValidateParamsWithHTTPClient(client *http.Client) *EulaResourceServiceValidateParams {
	return &EulaResourceServiceValidateParams{
		HTTPClient: client,
	}
}

/*
EulaResourceServiceValidateParams contains all the parameters to send to the API endpoint

	for the eula resource service validate operation.

	Typically these are written to a http.Request.
*/
type EulaResourceServiceValidateParams struct {

	/* OrgID.

	   Eula to validate.
	*/
	OrgID *string

	// TapVersion.
	TapVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eula resource service validate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EulaResourceServiceValidateParams) WithDefaults() *EulaResourceServiceValidateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eula resource service validate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EulaResourceServiceValidateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) WithTimeout(timeout time.Duration) *EulaResourceServiceValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) WithContext(ctx context.Context) *EulaResourceServiceValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) WithHTTPClient(client *http.Client) *EulaResourceServiceValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgID adds the orgID to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) WithOrgID(orgID *string) *EulaResourceServiceValidateParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) SetOrgID(orgID *string) {
	o.OrgID = orgID
}

// WithTapVersion adds the tapVersion to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) WithTapVersion(tapVersion *string) *EulaResourceServiceValidateParams {
	o.SetTapVersion(tapVersion)
	return o
}

// SetTapVersion adds the tapVersion to the eula resource service validate params
func (o *EulaResourceServiceValidateParams) SetTapVersion(tapVersion *string) {
	o.TapVersion = tapVersion
}

// WriteToRequest writes these params to a swagger request
func (o *EulaResourceServiceValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrgID != nil {

		// query param orgId
		var qrOrgID string

		if o.OrgID != nil {
			qrOrgID = *o.OrgID
		}
		qOrgID := qrOrgID
		if qOrgID != "" {

			if err := r.SetQueryParam("orgId", qOrgID); err != nil {
				return err
			}
		}
	}

	if o.TapVersion != nil {

		// query param tapVersion
		var qrTapVersion string

		if o.TapVersion != nil {
			qrTapVersion = *o.TapVersion
		}
		qTapVersion := qrTapVersion
		if qTapVersion != "" {

			if err := r.SetQueryParam("tapVersion", qTapVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
