// Code generated by go-swagger; DO NOT EDIT.

package nodepool_resource_service

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

// NewNodepoolResourceServiceGetParams creates a new NodepoolResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodepoolResourceServiceGetParams() *NodepoolResourceServiceGetParams {
	return &NodepoolResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodepoolResourceServiceGetParamsWithTimeout creates a new NodepoolResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewNodepoolResourceServiceGetParamsWithTimeout(timeout time.Duration) *NodepoolResourceServiceGetParams {
	return &NodepoolResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewNodepoolResourceServiceGetParamsWithContext creates a new NodepoolResourceServiceGetParams object
// with the ability to set a context for a request.
func NewNodepoolResourceServiceGetParamsWithContext(ctx context.Context) *NodepoolResourceServiceGetParams {
	return &NodepoolResourceServiceGetParams{
		Context: ctx,
	}
}

// NewNodepoolResourceServiceGetParamsWithHTTPClient creates a new NodepoolResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodepoolResourceServiceGetParamsWithHTTPClient(client *http.Client) *NodepoolResourceServiceGetParams {
	return &NodepoolResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
NodepoolResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the nodepool resource service get operation.

	Typically these are written to a http.Request.
*/
type NodepoolResourceServiceGetParams struct {

	/* FullNameManagementClusterName.

	   Name of the management cluster.
	*/
	FullNameManagementClusterName string

	/* FullNameName.

	   Name of this nodepool.
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

	/* FullNameTanzuKubernetesClusterName.

	   Name of the cluster.
	*/
	FullNameTanzuKubernetesClusterName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nodepool resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolResourceServiceGetParams) WithDefaults() *NodepoolResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodepool resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithTimeout(timeout time.Duration) *NodepoolResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithContext(ctx context.Context) *NodepoolResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithHTTPClient(client *http.Client) *NodepoolResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *NodepoolResourceServiceGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithFullNameName(fullNameName string) *NodepoolResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *NodepoolResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithFullNameProvisionerName(fullNameProvisionerName string) *NodepoolResourceServiceGetParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetFullNameProvisionerName(fullNameProvisionerName string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WithFullNameTanzuKubernetesClusterName adds the fullNameTanzuKubernetesClusterName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) WithFullNameTanzuKubernetesClusterName(fullNameTanzuKubernetesClusterName string) *NodepoolResourceServiceGetParams {
	o.SetFullNameTanzuKubernetesClusterName(fullNameTanzuKubernetesClusterName)
	return o
}

// SetFullNameTanzuKubernetesClusterName adds the fullNameTanzuKubernetesClusterName to the nodepool resource service get params
func (o *NodepoolResourceServiceGetParams) SetFullNameTanzuKubernetesClusterName(fullNameTanzuKubernetesClusterName string) {
	o.FullNameTanzuKubernetesClusterName = fullNameTanzuKubernetesClusterName
}

// WriteToRequest writes these params to a swagger request
func (o *NodepoolResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param fullName.tanzuKubernetesClusterName
	if err := r.SetPathParam("fullName.tanzuKubernetesClusterName", o.FullNameTanzuKubernetesClusterName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}