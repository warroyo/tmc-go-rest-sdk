// Code generated by go-swagger; DO NOT EDIT.

package release_resource_service

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

// NewReleaseResourceServiceUpdateParams creates a new ReleaseResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReleaseResourceServiceUpdateParams() *ReleaseResourceServiceUpdateParams {
	return &ReleaseResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseResourceServiceUpdateParamsWithTimeout creates a new ReleaseResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewReleaseResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *ReleaseResourceServiceUpdateParams {
	return &ReleaseResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewReleaseResourceServiceUpdateParamsWithContext creates a new ReleaseResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewReleaseResourceServiceUpdateParamsWithContext(ctx context.Context) *ReleaseResourceServiceUpdateParams {
	return &ReleaseResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewReleaseResourceServiceUpdateParamsWithHTTPClient creates a new ReleaseResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewReleaseResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *ReleaseResourceServiceUpdateParams {
	return &ReleaseResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
ReleaseResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the release resource service update operation.

	Typically these are written to a http.Request.
*/
type ReleaseResourceServiceUpdateParams struct {

	// Body.
	Body *models.ClusterNamespaceFluxcdHelmReleaseUpdateReleaseRequest

	/* ReleaseFullNameClusterName.

	   Name of Cluster.
	*/
	ReleaseFullNameClusterName string

	/* ReleaseFullNameName.

	   Name of the Release.
	*/
	ReleaseFullNameName string

	/* ReleaseFullNameNamespaceName.

	   Name of Namespace.
	*/
	ReleaseFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the release resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseResourceServiceUpdateParams) WithDefaults() *ReleaseResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the release resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *ReleaseResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithContext(ctx context.Context) *ReleaseResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *ReleaseResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithBody(body *models.ClusterNamespaceFluxcdHelmReleaseUpdateReleaseRequest) *ReleaseResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetBody(body *models.ClusterNamespaceFluxcdHelmReleaseUpdateReleaseRequest) {
	o.Body = body
}

// WithReleaseFullNameClusterName adds the releaseFullNameClusterName to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithReleaseFullNameClusterName(releaseFullNameClusterName string) *ReleaseResourceServiceUpdateParams {
	o.SetReleaseFullNameClusterName(releaseFullNameClusterName)
	return o
}

// SetReleaseFullNameClusterName adds the releaseFullNameClusterName to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetReleaseFullNameClusterName(releaseFullNameClusterName string) {
	o.ReleaseFullNameClusterName = releaseFullNameClusterName
}

// WithReleaseFullNameName adds the releaseFullNameName to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithReleaseFullNameName(releaseFullNameName string) *ReleaseResourceServiceUpdateParams {
	o.SetReleaseFullNameName(releaseFullNameName)
	return o
}

// SetReleaseFullNameName adds the releaseFullNameName to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetReleaseFullNameName(releaseFullNameName string) {
	o.ReleaseFullNameName = releaseFullNameName
}

// WithReleaseFullNameNamespaceName adds the releaseFullNameNamespaceName to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) WithReleaseFullNameNamespaceName(releaseFullNameNamespaceName string) *ReleaseResourceServiceUpdateParams {
	o.SetReleaseFullNameNamespaceName(releaseFullNameNamespaceName)
	return o
}

// SetReleaseFullNameNamespaceName adds the releaseFullNameNamespaceName to the release resource service update params
func (o *ReleaseResourceServiceUpdateParams) SetReleaseFullNameNamespaceName(releaseFullNameNamespaceName string) {
	o.ReleaseFullNameNamespaceName = releaseFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param release.fullName.clusterName
	if err := r.SetPathParam("release.fullName.clusterName", o.ReleaseFullNameClusterName); err != nil {
		return err
	}

	// path param release.fullName.name
	if err := r.SetPathParam("release.fullName.name", o.ReleaseFullNameName); err != nil {
		return err
	}

	// path param release.fullName.namespaceName
	if err := r.SetPathParam("release.fullName.namespaceName", o.ReleaseFullNameNamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
