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

// NewSecretResourceServiceUpdateParams creates a new SecretResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretResourceServiceUpdateParams() *SecretResourceServiceUpdateParams {
	return &SecretResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretResourceServiceUpdateParamsWithTimeout creates a new SecretResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewSecretResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *SecretResourceServiceUpdateParams {
	return &SecretResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewSecretResourceServiceUpdateParamsWithContext creates a new SecretResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewSecretResourceServiceUpdateParamsWithContext(ctx context.Context) *SecretResourceServiceUpdateParams {
	return &SecretResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewSecretResourceServiceUpdateParamsWithHTTPClient creates a new SecretResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *SecretResourceServiceUpdateParams {
	return &SecretResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
SecretResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the secret resource service update operation.

	Typically these are written to a http.Request.
*/
type SecretResourceServiceUpdateParams struct {

	// Body.
	Body *models.ClusterNamespaceSecretUpdateSecretRequest

	/* SecretFullNameClusterName.

	   Name of Cluster.
	*/
	SecretFullNameClusterName string

	/* SecretFullNameName.

	   Name of the Secret.
	*/
	SecretFullNameName string

	/* SecretFullNameNamespaceName.

	   Name of Namespace.
	*/
	SecretFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretResourceServiceUpdateParams) WithDefaults() *SecretResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *SecretResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithContext(ctx context.Context) *SecretResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *SecretResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithBody(body *models.ClusterNamespaceSecretUpdateSecretRequest) *SecretResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetBody(body *models.ClusterNamespaceSecretUpdateSecretRequest) {
	o.Body = body
}

// WithSecretFullNameClusterName adds the secretFullNameClusterName to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithSecretFullNameClusterName(secretFullNameClusterName string) *SecretResourceServiceUpdateParams {
	o.SetSecretFullNameClusterName(secretFullNameClusterName)
	return o
}

// SetSecretFullNameClusterName adds the secretFullNameClusterName to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetSecretFullNameClusterName(secretFullNameClusterName string) {
	o.SecretFullNameClusterName = secretFullNameClusterName
}

// WithSecretFullNameName adds the secretFullNameName to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithSecretFullNameName(secretFullNameName string) *SecretResourceServiceUpdateParams {
	o.SetSecretFullNameName(secretFullNameName)
	return o
}

// SetSecretFullNameName adds the secretFullNameName to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetSecretFullNameName(secretFullNameName string) {
	o.SecretFullNameName = secretFullNameName
}

// WithSecretFullNameNamespaceName adds the secretFullNameNamespaceName to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) WithSecretFullNameNamespaceName(secretFullNameNamespaceName string) *SecretResourceServiceUpdateParams {
	o.SetSecretFullNameNamespaceName(secretFullNameNamespaceName)
	return o
}

// SetSecretFullNameNamespaceName adds the secretFullNameNamespaceName to the secret resource service update params
func (o *SecretResourceServiceUpdateParams) SetSecretFullNameNamespaceName(secretFullNameNamespaceName string) {
	o.SecretFullNameNamespaceName = secretFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *SecretResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param secret.fullName.name
	if err := r.SetPathParam("secret.fullName.name", o.SecretFullNameName); err != nil {
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