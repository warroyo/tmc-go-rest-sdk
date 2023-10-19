// Code generated by go-swagger; DO NOT EDIT.

package secret_resource_service

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

// NewSecretResourceServiceCreateParams creates a new SecretResourceServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretResourceServiceCreateParams() *SecretResourceServiceCreateParams {
	return &SecretResourceServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretResourceServiceCreateParamsWithTimeout creates a new SecretResourceServiceCreateParams object
// with the ability to set a timeout on a request.
func NewSecretResourceServiceCreateParamsWithTimeout(timeout time.Duration) *SecretResourceServiceCreateParams {
	return &SecretResourceServiceCreateParams{
		timeout: timeout,
	}
}

// NewSecretResourceServiceCreateParamsWithContext creates a new SecretResourceServiceCreateParams object
// with the ability to set a context for a request.
func NewSecretResourceServiceCreateParamsWithContext(ctx context.Context) *SecretResourceServiceCreateParams {
	return &SecretResourceServiceCreateParams{
		Context: ctx,
	}
}

// NewSecretResourceServiceCreateParamsWithHTTPClient creates a new SecretResourceServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretResourceServiceCreateParamsWithHTTPClient(client *http.Client) *SecretResourceServiceCreateParams {
	return &SecretResourceServiceCreateParams{
		HTTPClient: client,
	}
}

/*
SecretResourceServiceCreateParams contains all the parameters to send to the API endpoint

	for the secret resource service create operation.

	Typically these are written to a http.Request.
*/
type SecretResourceServiceCreateParams struct {

	// Body.
	Body *models.ClusterNamespaceSecretCreateSecretRequest

	/* SecretFullNameClusterName.

	   Name of Cluster.
	*/
	SecretFullNameClusterName string

	/* SecretFullNameNamespaceName.

	   Name of Namespace.
	*/
	SecretFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretResourceServiceCreateParams) WithDefaults() *SecretResourceServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret resource service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretResourceServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret resource service create params
func (o *SecretResourceServiceCreateParams) WithTimeout(timeout time.Duration) *SecretResourceServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret resource service create params
func (o *SecretResourceServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret resource service create params
func (o *SecretResourceServiceCreateParams) WithContext(ctx context.Context) *SecretResourceServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret resource service create params
func (o *SecretResourceServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret resource service create params
func (o *SecretResourceServiceCreateParams) WithHTTPClient(client *http.Client) *SecretResourceServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret resource service create params
func (o *SecretResourceServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the secret resource service create params
func (o *SecretResourceServiceCreateParams) WithBody(body *models.ClusterNamespaceSecretCreateSecretRequest) *SecretResourceServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the secret resource service create params
func (o *SecretResourceServiceCreateParams) SetBody(body *models.ClusterNamespaceSecretCreateSecretRequest) {
	o.Body = body
}

// WithSecretFullNameClusterName adds the secretFullNameClusterName to the secret resource service create params
func (o *SecretResourceServiceCreateParams) WithSecretFullNameClusterName(secretFullNameClusterName string) *SecretResourceServiceCreateParams {
	o.SetSecretFullNameClusterName(secretFullNameClusterName)
	return o
}

// SetSecretFullNameClusterName adds the secretFullNameClusterName to the secret resource service create params
func (o *SecretResourceServiceCreateParams) SetSecretFullNameClusterName(secretFullNameClusterName string) {
	o.SecretFullNameClusterName = secretFullNameClusterName
}

// WithSecretFullNameNamespaceName adds the secretFullNameNamespaceName to the secret resource service create params
func (o *SecretResourceServiceCreateParams) WithSecretFullNameNamespaceName(secretFullNameNamespaceName string) *SecretResourceServiceCreateParams {
	o.SetSecretFullNameNamespaceName(secretFullNameNamespaceName)
	return o
}

// SetSecretFullNameNamespaceName adds the secretFullNameNamespaceName to the secret resource service create params
func (o *SecretResourceServiceCreateParams) SetSecretFullNameNamespaceName(secretFullNameNamespaceName string) {
	o.SecretFullNameNamespaceName = secretFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *SecretResourceServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param secret.fullName.clusterName
	if err := r.SetPathParam("secret.fullName.clusterName", o.SecretFullNameClusterName); err != nil {
		return err
	}

	// path param secret.fullName.namespaceName
	if err := r.SetPathParam("secret.fullName.namespaceName", o.SecretFullNameNamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
