// Code generated by go-swagger; DO NOT EDIT.

package tanzu_kubernetes_cluster_resource_service

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
)

// NewTanzuKubernetesClusterResourceServiceGetParams creates a new TanzuKubernetesClusterResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTanzuKubernetesClusterResourceServiceGetParams() *TanzuKubernetesClusterResourceServiceGetParams {
	return &TanzuKubernetesClusterResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTanzuKubernetesClusterResourceServiceGetParamsWithTimeout creates a new TanzuKubernetesClusterResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewTanzuKubernetesClusterResourceServiceGetParamsWithTimeout(timeout time.Duration) *TanzuKubernetesClusterResourceServiceGetParams {
	return &TanzuKubernetesClusterResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewTanzuKubernetesClusterResourceServiceGetParamsWithContext creates a new TanzuKubernetesClusterResourceServiceGetParams object
// with the ability to set a context for a request.
func NewTanzuKubernetesClusterResourceServiceGetParamsWithContext(ctx context.Context) *TanzuKubernetesClusterResourceServiceGetParams {
	return &TanzuKubernetesClusterResourceServiceGetParams{
		Context: ctx,
	}
}

// NewTanzuKubernetesClusterResourceServiceGetParamsWithHTTPClient creates a new TanzuKubernetesClusterResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTanzuKubernetesClusterResourceServiceGetParamsWithHTTPClient(client *http.Client) *TanzuKubernetesClusterResourceServiceGetParams {
	return &TanzuKubernetesClusterResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
TanzuKubernetesClusterResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the tanzu kubernetes cluster resource service get operation.

	Typically these are written to a http.Request.
*/
type TanzuKubernetesClusterResourceServiceGetParams struct {

	/* FullNameManagementClusterName.

	   Name of the management cluster.
	*/
	FullNameManagementClusterName string

	/* FullNameName.

	   Name of this cluster.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameProvisionerName.

	   Provisioner of the cluster.
	*/
	FullNameProvisionerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tanzu kubernetes cluster resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithDefaults() *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tanzu kubernetes cluster resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithTimeout(timeout time.Duration) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithContext(ctx context.Context) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithHTTPClient(client *http.Client) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithFullNameName(fullNameName string) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) WithFullNameProvisionerName(fullNameProvisionerName string) *TanzuKubernetesClusterResourceServiceGetParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the tanzu kubernetes cluster resource service get params
func (o *TanzuKubernetesClusterResourceServiceGetParams) SetFullNameProvisionerName(fullNameProvisionerName string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *TanzuKubernetesClusterResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fullName.managementClusterName
	if err := r.SetPathParam("fullName.managementClusterName", o.FullNameManagementClusterName); err != nil {
		return err
	}

	// path param fullName.name
	if err := r.SetPathParam("fullName.name", o.FullNameName); err != nil {
		return err
	}

	if o.FullNameOrgID != nil {

		// query param fullName.orgId
		var qrFullNameOrgID string

		if o.FullNameOrgID != nil {
			qrFullNameOrgID = *o.FullNameOrgID
		}
		qFullNameOrgID := qrFullNameOrgID
		if qFullNameOrgID != "" {

			if err := r.SetQueryParam("fullName.orgId", qFullNameOrgID); err != nil {
				return err
			}
		}
	}

	// path param fullName.provisionerName
	if err := r.SetPathParam("fullName.provisionerName", o.FullNameProvisionerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
