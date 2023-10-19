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

// NewNodepoolResourceServiceDeleteParams creates a new NodepoolResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodepoolResourceServiceDeleteParams() *NodepoolResourceServiceDeleteParams {
	return &NodepoolResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodepoolResourceServiceDeleteParamsWithTimeout creates a new NodepoolResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewNodepoolResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *NodepoolResourceServiceDeleteParams {
	return &NodepoolResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewNodepoolResourceServiceDeleteParamsWithContext creates a new NodepoolResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewNodepoolResourceServiceDeleteParamsWithContext(ctx context.Context) *NodepoolResourceServiceDeleteParams {
	return &NodepoolResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewNodepoolResourceServiceDeleteParamsWithHTTPClient creates a new NodepoolResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodepoolResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *NodepoolResourceServiceDeleteParams {
	return &NodepoolResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
NodepoolResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the nodepool resource service delete operation.

	Typically these are written to a http.Request.
*/
type NodepoolResourceServiceDeleteParams struct {

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

// WithDefaults hydrates default values in the nodepool resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolResourceServiceDeleteParams) WithDefaults() *NodepoolResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodepool resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *NodepoolResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithContext(ctx context.Context) *NodepoolResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *NodepoolResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *NodepoolResourceServiceDeleteParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithFullNameName(fullNameName string) *NodepoolResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *NodepoolResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithFullNameProvisionerName(fullNameProvisionerName string) *NodepoolResourceServiceDeleteParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetFullNameProvisionerName(fullNameProvisionerName string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WithFullNameTanzuKubernetesClusterName adds the fullNameTanzuKubernetesClusterName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) WithFullNameTanzuKubernetesClusterName(fullNameTanzuKubernetesClusterName string) *NodepoolResourceServiceDeleteParams {
	o.SetFullNameTanzuKubernetesClusterName(fullNameTanzuKubernetesClusterName)
	return o
}

// SetFullNameTanzuKubernetesClusterName adds the fullNameTanzuKubernetesClusterName to the nodepool resource service delete params
func (o *NodepoolResourceServiceDeleteParams) SetFullNameTanzuKubernetesClusterName(fullNameTanzuKubernetesClusterName string) {
	o.FullNameTanzuKubernetesClusterName = fullNameTanzuKubernetesClusterName
}

// WriteToRequest writes these params to a swagger request
func (o *NodepoolResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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