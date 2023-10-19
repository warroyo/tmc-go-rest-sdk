// Code generated by go-swagger; DO NOT EDIT.

package namespace_resource_service

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

// NewNamespaceResourceServiceGetParams creates a new NamespaceResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNamespaceResourceServiceGetParams() *NamespaceResourceServiceGetParams {
	return &NamespaceResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNamespaceResourceServiceGetParamsWithTimeout creates a new NamespaceResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewNamespaceResourceServiceGetParamsWithTimeout(timeout time.Duration) *NamespaceResourceServiceGetParams {
	return &NamespaceResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewNamespaceResourceServiceGetParamsWithContext creates a new NamespaceResourceServiceGetParams object
// with the ability to set a context for a request.
func NewNamespaceResourceServiceGetParamsWithContext(ctx context.Context) *NamespaceResourceServiceGetParams {
	return &NamespaceResourceServiceGetParams{
		Context: ctx,
	}
}

// NewNamespaceResourceServiceGetParamsWithHTTPClient creates a new NamespaceResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNamespaceResourceServiceGetParamsWithHTTPClient(client *http.Client) *NamespaceResourceServiceGetParams {
	return &NamespaceResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
NamespaceResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the namespace resource service get operation.

	Typically these are written to a http.Request.
*/
type NamespaceResourceServiceGetParams struct {

	/* FullNameClusterName.

	   Name of Cluster.
	*/
	FullNameClusterName string

	/* FullNameManagementClusterName.

	   Name of ManagementCluster.
	*/
	FullNameManagementClusterName *string

	/* FullNameName.

	   Name of Namespace.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameProvisionerName.

	   Name of Provisioner.
	*/
	FullNameProvisionerName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the namespace resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceResourceServiceGetParams) WithDefaults() *NamespaceResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the namespace resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithTimeout(timeout time.Duration) *NamespaceResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithContext(ctx context.Context) *NamespaceResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithHTTPClient(client *http.Client) *NamespaceResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameClusterName adds the fullNameClusterName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithFullNameClusterName(fullNameClusterName string) *NamespaceResourceServiceGetParams {
	o.SetFullNameClusterName(fullNameClusterName)
	return o
}

// SetFullNameClusterName adds the fullNameClusterName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetFullNameClusterName(fullNameClusterName string) {
	o.FullNameClusterName = fullNameClusterName
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *NamespaceResourceServiceGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithFullNameName(fullNameName string) *NamespaceResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *NamespaceResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *NamespaceResourceServiceGetParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the namespace resource service get params
func (o *NamespaceResourceServiceGetParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *NamespaceResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fullName.clusterName
	if err := r.SetPathParam("fullName.clusterName", o.FullNameClusterName); err != nil {
		return err
	}

	if o.FullNameManagementClusterName != nil {

		// query param fullName.managementClusterName
		var qrFullNameManagementClusterName string

		if o.FullNameManagementClusterName != nil {
			qrFullNameManagementClusterName = *o.FullNameManagementClusterName
		}
		qFullNameManagementClusterName := qrFullNameManagementClusterName
		if qFullNameManagementClusterName != "" {

			if err := r.SetQueryParam("fullName.managementClusterName", qFullNameManagementClusterName); err != nil {
				return err
			}
		}
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

	if o.FullNameProvisionerName != nil {

		// query param fullName.provisionerName
		var qrFullNameProvisionerName string

		if o.FullNameProvisionerName != nil {
			qrFullNameProvisionerName = *o.FullNameProvisionerName
		}
		qFullNameProvisionerName := qrFullNameProvisionerName
		if qFullNameProvisionerName != "" {

			if err := r.SetQueryParam("fullName.provisionerName", qFullNameProvisionerName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
