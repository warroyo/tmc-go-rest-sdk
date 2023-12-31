// Code generated by go-swagger; DO NOT EDIT.

package secret_export_resource_service

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

// NewSecretExportResourceServiceCreateParams creates a new SecretExportResourceServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretExportResourceServiceCreateParams() *SecretExportResourceServiceCreateParams {
	return &SecretExportResourceServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretExportResourceServiceCreateParamsWithTimeout creates a new SecretExportResourceServiceCreateParams object
// with the ability to set a timeout on a request.
func NewSecretExportResourceServiceCreateParamsWithTimeout(timeout time.Duration) *SecretExportResourceServiceCreateParams {
	return &SecretExportResourceServiceCreateParams{
		timeout: timeout,
	}
}

// NewSecretExportResourceServiceCreateParamsWithContext creates a new SecretExportResourceServiceCreateParams object
// with the ability to set a context for a request.
func NewSecretExportResourceServiceCreateParamsWithContext(ctx context.Context) *SecretExportResourceServiceCreateParams {
	return &SecretExportResourceServiceCreateParams{
		Context: ctx,
	}
}

// NewSecretExportResourceServiceCreateParamsWithHTTPClient creates a new SecretExportResourceServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretExportResourceServiceCreateParamsWithHTTPClient(client *http.Client) *SecretExportResourceServiceCreateParams {
	return &SecretExportResourceServiceCreateParams{
		HTTPClient: client,
	}
}

/*
SecretExportResourceServiceCreateParams contains all the parameters to send to the API endpoint

	for the secret export resource service create operation.

	Typically these are written to a http.Request.
*/
type SecretExportResourceServiceCreateParams struct {

	// Body.
	Body *models.ClusterNamespaceSecretexportCreateSecretExportRequest

	/* SecretExportFullNameClusterName.

	   Name of Cluster.
	*/
	SecretExportFullNameClusterName string

	/* SecretExportFullNameNamespaceName.

	   Name of Namespace.
	*/
	SecretExportFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret export resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretExportResourceServiceCreateParams) WithDefaults() *SecretExportResourceServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret export resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretExportResourceServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) WithTimeout(timeout time.Duration) *SecretExportResourceServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) WithContext(ctx context.Context) *SecretExportResourceServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) WithHTTPClient(client *http.Client) *SecretExportResourceServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) WithBody(body *models.ClusterNamespaceSecretexportCreateSecretExportRequest) *SecretExportResourceServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) SetBody(body *models.ClusterNamespaceSecretexportCreateSecretExportRequest) {
	o.Body = body
}

// WithSecretExportFullNameClusterName adds the secretExportFullNameClusterName to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) WithSecretExportFullNameClusterName(secretExportFullNameClusterName string) *SecretExportResourceServiceCreateParams {
	o.SetSecretExportFullNameClusterName(secretExportFullNameClusterName)
	return o
}

// SetSecretExportFullNameClusterName adds the secretExportFullNameClusterName to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) SetSecretExportFullNameClusterName(secretExportFullNameClusterName string) {
	o.SecretExportFullNameClusterName = secretExportFullNameClusterName
}

// WithSecretExportFullNameNamespaceName adds the secretExportFullNameNamespaceName to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) WithSecretExportFullNameNamespaceName(secretExportFullNameNamespaceName string) *SecretExportResourceServiceCreateParams {
	o.SetSecretExportFullNameNamespaceName(secretExportFullNameNamespaceName)
	return o
}

// SetSecretExportFullNameNamespaceName adds the secretExportFullNameNamespaceName to the secret export resource service create params
func (o *SecretExportResourceServiceCreateParams) SetSecretExportFullNameNamespaceName(secretExportFullNameNamespaceName string) {
	o.SecretExportFullNameNamespaceName = secretExportFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *SecretExportResourceServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param secretExport.fullName.clusterName
	if err := r.SetPathParam("secretExport.fullName.clusterName", o.SecretExportFullNameClusterName); err != nil {
		return err
	}

	// path param secretExport.fullName.namespaceName
	if err := r.SetPathParam("secretExport.fullName.namespaceName", o.SecretExportFullNameNamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
