// Code generated by go-swagger; DO NOT EDIT.

package preference_resource_service

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

// NewPreferenceResourceServiceUpdateParams creates a new PreferenceResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPreferenceResourceServiceUpdateParams() *PreferenceResourceServiceUpdateParams {
	return &PreferenceResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPreferenceResourceServiceUpdateParamsWithTimeout creates a new PreferenceResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewPreferenceResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *PreferenceResourceServiceUpdateParams {
	return &PreferenceResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewPreferenceResourceServiceUpdateParamsWithContext creates a new PreferenceResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewPreferenceResourceServiceUpdateParamsWithContext(ctx context.Context) *PreferenceResourceServiceUpdateParams {
	return &PreferenceResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewPreferenceResourceServiceUpdateParamsWithHTTPClient creates a new PreferenceResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPreferenceResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *PreferenceResourceServiceUpdateParams {
	return &PreferenceResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
PreferenceResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the preference resource service update operation.

	Typically these are written to a http.Request.
*/
type PreferenceResourceServiceUpdateParams struct {

	// Body.
	Body *models.IamCurrentuserPreferenceUpdatePreferenceRequest

	/* PreferenceFullNameName.

	   Name of this preference.
	*/
	PreferenceFullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the preference resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PreferenceResourceServiceUpdateParams) WithDefaults() *PreferenceResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the preference resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PreferenceResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *PreferenceResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) WithContext(ctx context.Context) *PreferenceResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *PreferenceResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) WithBody(body *models.IamCurrentuserPreferenceUpdatePreferenceRequest) *PreferenceResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) SetBody(body *models.IamCurrentuserPreferenceUpdatePreferenceRequest) {
	o.Body = body
}

// WithPreferenceFullNameName adds the preferenceFullNameName to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) WithPreferenceFullNameName(preferenceFullNameName string) *PreferenceResourceServiceUpdateParams {
	o.SetPreferenceFullNameName(preferenceFullNameName)
	return o
}

// SetPreferenceFullNameName adds the preferenceFullNameName to the preference resource service update params
func (o *PreferenceResourceServiceUpdateParams) SetPreferenceFullNameName(preferenceFullNameName string) {
	o.PreferenceFullNameName = preferenceFullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *PreferenceResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param preference.fullName.name
	if err := r.SetPathParam("preference.fullName.name", o.PreferenceFullNameName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}