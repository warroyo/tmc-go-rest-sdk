// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_resource_service

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
	"github.com/go-openapi/swag"
)

// NewManagementClusterResourceServiceDeleteParams creates a new ManagementClusterResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagementClusterResourceServiceDeleteParams() *ManagementClusterResourceServiceDeleteParams {
	return &ManagementClusterResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagementClusterResourceServiceDeleteParamsWithTimeout creates a new ManagementClusterResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewManagementClusterResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *ManagementClusterResourceServiceDeleteParams {
	return &ManagementClusterResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewManagementClusterResourceServiceDeleteParamsWithContext creates a new ManagementClusterResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewManagementClusterResourceServiceDeleteParamsWithContext(ctx context.Context) *ManagementClusterResourceServiceDeleteParams {
	return &ManagementClusterResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewManagementClusterResourceServiceDeleteParamsWithHTTPClient creates a new ManagementClusterResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagementClusterResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *ManagementClusterResourceServiceDeleteParams {
	return &ManagementClusterResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
ManagementClusterResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the management cluster resource service delete operation.

	Typically these are written to a http.Request.
*/
type ManagementClusterResourceServiceDeleteParams struct {

	/* Force.

	   Force delete.
	*/
	Force *bool

	/* FullNameName.

	   Unique identifier of the ManagementCluster.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization. Generally a GUID.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the management cluster resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterResourceServiceDeleteParams) WithDefaults() *ManagementClusterResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the management cluster resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *ManagementClusterResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) WithContext(ctx context.Context) *ManagementClusterResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *ManagementClusterResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) WithForce(force *bool) *ManagementClusterResourceServiceDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) SetForce(force *bool) {
	o.Force = force
}

// WithFullNameName adds the fullNameName to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) WithFullNameName(fullNameName string) *ManagementClusterResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *ManagementClusterResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the management cluster resource service delete params
func (o *ManagementClusterResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ManagementClusterResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}