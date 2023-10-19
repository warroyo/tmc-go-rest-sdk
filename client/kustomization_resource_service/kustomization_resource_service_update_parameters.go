// Code generated by go-swagger; DO NOT EDIT.

package kustomization_resource_service

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

// NewKustomizationResourceServiceUpdateParams creates a new KustomizationResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKustomizationResourceServiceUpdateParams() *KustomizationResourceServiceUpdateParams {
	return &KustomizationResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKustomizationResourceServiceUpdateParamsWithTimeout creates a new KustomizationResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewKustomizationResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *KustomizationResourceServiceUpdateParams {
	return &KustomizationResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewKustomizationResourceServiceUpdateParamsWithContext creates a new KustomizationResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewKustomizationResourceServiceUpdateParamsWithContext(ctx context.Context) *KustomizationResourceServiceUpdateParams {
	return &KustomizationResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewKustomizationResourceServiceUpdateParamsWithHTTPClient creates a new KustomizationResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewKustomizationResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *KustomizationResourceServiceUpdateParams {
	return &KustomizationResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
KustomizationResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the kustomization resource service update operation.

	Typically these are written to a http.Request.
*/
type KustomizationResourceServiceUpdateParams struct {

	// Body.
	Body *models.ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest

	/* KustomizationFullNameClusterName.

	   Name of Cluster.
	*/
	KustomizationFullNameClusterName string

	/* KustomizationFullNameName.

	   Name of the Kustomization.
	*/
	KustomizationFullNameName string

	/* KustomizationFullNameNamespaceName.

	   Name of Namespace.
	*/
	KustomizationFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kustomization resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KustomizationResourceServiceUpdateParams) WithDefaults() *KustomizationResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kustomization resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KustomizationResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *KustomizationResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithContext(ctx context.Context) *KustomizationResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *KustomizationResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithBody(body *models.ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) *KustomizationResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetBody(body *models.ClusterNamespaceFluxcdKustomizationUpdateKustomizationRequest) {
	o.Body = body
}

// WithKustomizationFullNameClusterName adds the kustomizationFullNameClusterName to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithKustomizationFullNameClusterName(kustomizationFullNameClusterName string) *KustomizationResourceServiceUpdateParams {
	o.SetKustomizationFullNameClusterName(kustomizationFullNameClusterName)
	return o
}

// SetKustomizationFullNameClusterName adds the kustomizationFullNameClusterName to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetKustomizationFullNameClusterName(kustomizationFullNameClusterName string) {
	o.KustomizationFullNameClusterName = kustomizationFullNameClusterName
}

// WithKustomizationFullNameName adds the kustomizationFullNameName to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithKustomizationFullNameName(kustomizationFullNameName string) *KustomizationResourceServiceUpdateParams {
	o.SetKustomizationFullNameName(kustomizationFullNameName)
	return o
}

// SetKustomizationFullNameName adds the kustomizationFullNameName to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetKustomizationFullNameName(kustomizationFullNameName string) {
	o.KustomizationFullNameName = kustomizationFullNameName
}

// WithKustomizationFullNameNamespaceName adds the kustomizationFullNameNamespaceName to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) WithKustomizationFullNameNamespaceName(kustomizationFullNameNamespaceName string) *KustomizationResourceServiceUpdateParams {
	o.SetKustomizationFullNameNamespaceName(kustomizationFullNameNamespaceName)
	return o
}

// SetKustomizationFullNameNamespaceName adds the kustomizationFullNameNamespaceName to the kustomization resource service update params
func (o *KustomizationResourceServiceUpdateParams) SetKustomizationFullNameNamespaceName(kustomizationFullNameNamespaceName string) {
	o.KustomizationFullNameNamespaceName = kustomizationFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *KustomizationResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param kustomization.fullName.clusterName
	if err := r.SetPathParam("kustomization.fullName.clusterName", o.KustomizationFullNameClusterName); err != nil {
		return err
	}

	// path param kustomization.fullName.name
	if err := r.SetPathParam("kustomization.fullName.name", o.KustomizationFullNameName); err != nil {
		return err
	}

	// path param kustomization.fullName.namespaceName
	if err := r.SetPathParam("kustomization.fullName.namespaceName", o.KustomizationFullNameNamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}