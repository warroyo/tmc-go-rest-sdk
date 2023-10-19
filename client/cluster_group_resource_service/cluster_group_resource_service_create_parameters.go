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

// NewClusterGroupResourceServiceCreateParams creates a new ClusterGroupResourceServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterGroupResourceServiceCreateParams() *ClusterGroupResourceServiceCreateParams {
	return &ClusterGroupResourceServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterGroupResourceServiceCreateParamsWithTimeout creates a new ClusterGroupResourceServiceCreateParams object
// with the ability to set a timeout on a request.
func NewClusterGroupResourceServiceCreateParamsWithTimeout(timeout time.Duration) *ClusterGroupResourceServiceCreateParams {
	return &ClusterGroupResourceServiceCreateParams{
		timeout: timeout,
	}
}

// NewClusterGroupResourceServiceCreateParamsWithContext creates a new ClusterGroupResourceServiceCreateParams object
// with the ability to set a context for a request.
func NewClusterGroupResourceServiceCreateParamsWithContext(ctx context.Context) *ClusterGroupResourceServiceCreateParams {
	return &ClusterGroupResourceServiceCreateParams{
		Context: ctx,
	}
}

// NewClusterGroupResourceServiceCreateParamsWithHTTPClient creates a new ClusterGroupResourceServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterGroupResourceServiceCreateParamsWithHTTPClient(client *http.Client) *ClusterGroupResourceServiceCreateParams {
	return &ClusterGroupResourceServiceCreateParams{
		HTTPClient: client,
	}
}

/*
ClusterGroupResourceServiceCreateParams contains all the parameters to send to the API endpoint

	for the cluster group resource service create operation.

	Typically these are written to a http.Request.
*/
type ClusterGroupResourceServiceCreateParams struct {

	// Body.
	Body *models.ClustergroupCreateClusterGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster group resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGroupResourceServiceCreateParams) WithDefaults() *ClusterGroupResourceServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster group resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGroupResourceServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) WithTimeout(timeout time.Duration) *ClusterGroupResourceServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) WithContext(ctx context.Context) *ClusterGroupResourceServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) WithHTTPClient(client *http.Client) *ClusterGroupResourceServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) WithBody(body *models.ClustergroupCreateClusterGroupRequest) *ClusterGroupResourceServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cluster group resource service create params
func (o *ClusterGroupResourceServiceCreateParams) SetBody(body *models.ClustergroupCreateClusterGroupRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterGroupResourceServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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