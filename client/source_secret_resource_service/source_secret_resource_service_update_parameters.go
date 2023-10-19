// Code generated by go-swagger; DO NOT EDIT.

package source_secret_resource_service

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

// NewSourceSecretResourceServiceUpdateParams creates a new SourceSecretResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSourceSecretResourceServiceUpdateParams() *SourceSecretResourceServiceUpdateParams {
	return &SourceSecretResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSourceSecretResourceServiceUpdateParamsWithTimeout creates a new SourceSecretResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewSourceSecretResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *SourceSecretResourceServiceUpdateParams {
	return &SourceSecretResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewSourceSecretResourceServiceUpdateParamsWithContext creates a new SourceSecretResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewSourceSecretResourceServiceUpdateParamsWithContext(ctx context.Context) *SourceSecretResourceServiceUpdateParams {
	return &SourceSecretResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewSourceSecretResourceServiceUpdateParamsWithHTTPClient creates a new SourceSecretResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSourceSecretResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *SourceSecretResourceServiceUpdateParams {
	return &SourceSecretResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
SourceSecretResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the source secret resource service update operation.

	Typically these are written to a http.Request.
*/
type SourceSecretResourceServiceUpdateParams struct {

	// Body.
	Body *models.ClusterFluxcdSourcesecretUpdateSourceSecretRequest

	/* SourceSecretFullNameClusterName.

	   Name of Cluster.
	*/
	SourceSecretFullNameClusterName string

	/* SourceSecretFullNameName.

	   Name of Source Secret.
	*/
	SourceSecretFullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the source secret resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceSecretResourceServiceUpdateParams) WithDefaults() *SourceSecretResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the source secret resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceSecretResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *SourceSecretResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) WithContext(ctx context.Context) *SourceSecretResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *SourceSecretResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) WithBody(body *models.ClusterFluxcdSourcesecretUpdateSourceSecretRequest) *SourceSecretResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) SetBody(body *models.ClusterFluxcdSourcesecretUpdateSourceSecretRequest) {
	o.Body = body
}

// WithSourceSecretFullNameClusterName adds the sourceSecretFullNameClusterName to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) WithSourceSecretFullNameClusterName(sourceSecretFullNameClusterName string) *SourceSecretResourceServiceUpdateParams {
	o.SetSourceSecretFullNameClusterName(sourceSecretFullNameClusterName)
	return o
}

// SetSourceSecretFullNameClusterName adds the sourceSecretFullNameClusterName to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) SetSourceSecretFullNameClusterName(sourceSecretFullNameClusterName string) {
	o.SourceSecretFullNameClusterName = sourceSecretFullNameClusterName
}

// WithSourceSecretFullNameName adds the sourceSecretFullNameName to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) WithSourceSecretFullNameName(sourceSecretFullNameName string) *SourceSecretResourceServiceUpdateParams {
	o.SetSourceSecretFullNameName(sourceSecretFullNameName)
	return o
}

// SetSourceSecretFullNameName adds the sourceSecretFullNameName to the source secret resource service update params
func (o *SourceSecretResourceServiceUpdateParams) SetSourceSecretFullNameName(sourceSecretFullNameName string) {
	o.SourceSecretFullNameName = sourceSecretFullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *SourceSecretResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sourceSecret.fullName.clusterName
	if err := r.SetPathParam("sourceSecret.fullName.clusterName", o.SourceSecretFullNameClusterName); err != nil {
		return err
	}

	// path param sourceSecret.fullName.name
	if err := r.SetPathParam("sourceSecret.fullName.name", o.SourceSecretFullNameName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
