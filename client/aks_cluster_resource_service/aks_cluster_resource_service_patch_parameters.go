// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_resource_service

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

// NewAksClusterResourceServicePatchParams creates a new AksClusterResourceServicePatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAksClusterResourceServicePatchParams() *AksClusterResourceServicePatchParams {
	return &AksClusterResourceServicePatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAksClusterResourceServicePatchParamsWithTimeout creates a new AksClusterResourceServicePatchParams object
// with the ability to set a timeout on a request.
func NewAksClusterResourceServicePatchParamsWithTimeout(timeout time.Duration) *AksClusterResourceServicePatchParams {
	return &AksClusterResourceServicePatchParams{
		timeout: timeout,
	}
}

// NewAksClusterResourceServicePatchParamsWithContext creates a new AksClusterResourceServicePatchParams object
// with the ability to set a context for a request.
func NewAksClusterResourceServicePatchParamsWithContext(ctx context.Context) *AksClusterResourceServicePatchParams {
	return &AksClusterResourceServicePatchParams{
		Context: ctx,
	}
}

// NewAksClusterResourceServicePatchParamsWithHTTPClient creates a new AksClusterResourceServicePatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewAksClusterResourceServicePatchParamsWithHTTPClient(client *http.Client) *AksClusterResourceServicePatchParams {
	return &AksClusterResourceServicePatchParams{
		HTTPClient: client,
	}
}

/*
AksClusterResourceServicePatchParams contains all the parameters to send to the API endpoint

	for the aks cluster resource service patch operation.

	Typically these are written to a http.Request.
*/
type AksClusterResourceServicePatchParams struct {

	// Body.
	Body *models.AksclusterPatchAksClusterRequest

	/* FullNameName.

	   Name of this cluster.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aks cluster resource service patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AksClusterResourceServicePatchParams) WithDefaults() *AksClusterResourceServicePatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aks cluster resource service patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AksClusterResourceServicePatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) WithTimeout(timeout time.Duration) *AksClusterResourceServicePatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) WithContext(ctx context.Context) *AksClusterResourceServicePatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) WithHTTPClient(client *http.Client) *AksClusterResourceServicePatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) WithBody(body *models.AksclusterPatchAksClusterRequest) *AksClusterResourceServicePatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) SetBody(body *models.AksclusterPatchAksClusterRequest) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) WithFullNameName(fullNameName string) *AksClusterResourceServicePatchParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the aks cluster resource service patch params
func (o *AksClusterResourceServicePatchParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *AksClusterResourceServicePatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param fullName.name
	if err := r.SetPathParam("fullName.name", o.FullNameName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}