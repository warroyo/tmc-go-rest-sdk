// Code generated by go-swagger; DO NOT EDIT.

package cluster_group_resource_service

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

// NewClusterGroupResourceServicePatchParams creates a new ClusterGroupResourceServicePatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterGroupResourceServicePatchParams() *ClusterGroupResourceServicePatchParams {
	return &ClusterGroupResourceServicePatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterGroupResourceServicePatchParamsWithTimeout creates a new ClusterGroupResourceServicePatchParams object
// with the ability to set a timeout on a request.
func NewClusterGroupResourceServicePatchParamsWithTimeout(timeout time.Duration) *ClusterGroupResourceServicePatchParams {
	return &ClusterGroupResourceServicePatchParams{
		timeout: timeout,
	}
}

// NewClusterGroupResourceServicePatchParamsWithContext creates a new ClusterGroupResourceServicePatchParams object
// with the ability to set a context for a request.
func NewClusterGroupResourceServicePatchParamsWithContext(ctx context.Context) *ClusterGroupResourceServicePatchParams {
	return &ClusterGroupResourceServicePatchParams{
		Context: ctx,
	}
}

// NewClusterGroupResourceServicePatchParamsWithHTTPClient creates a new ClusterGroupResourceServicePatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterGroupResourceServicePatchParamsWithHTTPClient(client *http.Client) *ClusterGroupResourceServicePatchParams {
	return &ClusterGroupResourceServicePatchParams{
		HTTPClient: client,
	}
}

/*
ClusterGroupResourceServicePatchParams contains all the parameters to send to the API endpoint

	for the cluster group resource service patch operation.

	Typically these are written to a http.Request.
*/
type ClusterGroupResourceServicePatchParams struct {

	// Body.
	Body *models.ClustergroupPatchClusterGroupRequest

	/* FullNameName.

	   Name of this ClusterGroup.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster group resource service patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGroupResourceServicePatchParams) WithDefaults() *ClusterGroupResourceServicePatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster group resource service patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGroupResourceServicePatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) WithTimeout(timeout time.Duration) *ClusterGroupResourceServicePatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) WithContext(ctx context.Context) *ClusterGroupResourceServicePatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) WithHTTPClient(client *http.Client) *ClusterGroupResourceServicePatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) WithBody(body *models.ClustergroupPatchClusterGroupRequest) *ClusterGroupResourceServicePatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) SetBody(body *models.ClustergroupPatchClusterGroupRequest) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) WithFullNameName(fullNameName string) *ClusterGroupResourceServicePatchParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the cluster group resource service patch params
func (o *ClusterGroupResourceServicePatchParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterGroupResourceServicePatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
