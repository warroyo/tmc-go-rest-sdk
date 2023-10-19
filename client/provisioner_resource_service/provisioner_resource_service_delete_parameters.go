// Code generated by go-swagger; DO NOT EDIT.

package provisioner_resource_service

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

// NewProvisionerResourceServiceDeleteParams creates a new ProvisionerResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionerResourceServiceDeleteParams() *ProvisionerResourceServiceDeleteParams {
	return &ProvisionerResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionerResourceServiceDeleteParamsWithTimeout creates a new ProvisionerResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewProvisionerResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *ProvisionerResourceServiceDeleteParams {
	return &ProvisionerResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewProvisionerResourceServiceDeleteParamsWithContext creates a new ProvisionerResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewProvisionerResourceServiceDeleteParamsWithContext(ctx context.Context) *ProvisionerResourceServiceDeleteParams {
	return &ProvisionerResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewProvisionerResourceServiceDeleteParamsWithHTTPClient creates a new ProvisionerResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionerResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *ProvisionerResourceServiceDeleteParams {
	return &ProvisionerResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
ProvisionerResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the provisioner resource service delete operation.

	Typically these are written to a http.Request.
*/
type ProvisionerResourceServiceDeleteParams struct {

	/* FullNameManagementClusterName.

	   Name of the ManagementCluster.
	*/
	FullNameManagementClusterName string

	/* FullNameName.

	   Name of the provisioner. It must be unique within a management cluster.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provisioner resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerResourceServiceDeleteParams) WithDefaults() *ProvisionerResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provisioner resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *ProvisionerResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) WithContext(ctx context.Context) *ProvisionerResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *ProvisionerResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *ProvisionerResourceServiceDeleteParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) WithFullNameName(fullNameName string) *ProvisionerResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *ProvisionerResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the provisioner resource service delete params
func (o *ProvisionerResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionerResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
