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

// NewSettingResourceServiceUpdateParams creates a new SettingResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingResourceServiceUpdateParams() *SettingResourceServiceUpdateParams {
	return &SettingResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingResourceServiceUpdateParamsWithTimeout creates a new SettingResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewSettingResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *SettingResourceServiceUpdateParams {
	return &SettingResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewSettingResourceServiceUpdateParamsWithContext creates a new SettingResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewSettingResourceServiceUpdateParamsWithContext(ctx context.Context) *SettingResourceServiceUpdateParams {
	return &SettingResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewSettingResourceServiceUpdateParamsWithHTTPClient creates a new SettingResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *SettingResourceServiceUpdateParams {
	return &SettingResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
SettingResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the setting resource service update operation.

	Typically these are written to a http.Request.
*/
type SettingResourceServiceUpdateParams struct {

	// Body.
	Body *models.OrganizationSettingUpdateSettingRequest

	/* SettingFullNameName.

	     Name of the setting.
	This should be same as the name of setting type.
	*/
	SettingFullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the setting resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingResourceServiceUpdateParams) WithDefaults() *SettingResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the setting resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *SettingResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) WithContext(ctx context.Context) *SettingResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *SettingResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) WithBody(body *models.OrganizationSettingUpdateSettingRequest) *SettingResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) SetBody(body *models.OrganizationSettingUpdateSettingRequest) {
	o.Body = body
}

// WithSettingFullNameName adds the settingFullNameName to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) WithSettingFullNameName(settingFullNameName string) *SettingResourceServiceUpdateParams {
	o.SetSettingFullNameName(settingFullNameName)
	return o
}

// SetSettingFullNameName adds the settingFullNameName to the setting resource service update params
func (o *SettingResourceServiceUpdateParams) SetSettingFullNameName(settingFullNameName string) {
	o.SettingFullNameName = settingFullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *SettingResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param setting.fullName.name
	if err := r.SetPathParam("setting.fullName.name", o.SettingFullNameName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
