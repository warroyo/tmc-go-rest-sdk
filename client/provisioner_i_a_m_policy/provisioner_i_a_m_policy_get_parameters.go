// Code generated by go-swagger; DO NOT EDIT.

package provisioner_i_a_m_policy

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

// NewProvisionerIAMPolicyGetParams creates a new ProvisionerIAMPolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionerIAMPolicyGetParams() *ProvisionerIAMPolicyGetParams {
	return &ProvisionerIAMPolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionerIAMPolicyGetParamsWithTimeout creates a new ProvisionerIAMPolicyGetParams object
// with the ability to set a timeout on a request.
func NewProvisionerIAMPolicyGetParamsWithTimeout(timeout time.Duration) *ProvisionerIAMPolicyGetParams {
	return &ProvisionerIAMPolicyGetParams{
		timeout: timeout,
	}
}

// NewProvisionerIAMPolicyGetParamsWithContext creates a new ProvisionerIAMPolicyGetParams object
// with the ability to set a context for a request.
func NewProvisionerIAMPolicyGetParamsWithContext(ctx context.Context) *ProvisionerIAMPolicyGetParams {
	return &ProvisionerIAMPolicyGetParams{
		Context: ctx,
	}
}

// NewProvisionerIAMPolicyGetParamsWithHTTPClient creates a new ProvisionerIAMPolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionerIAMPolicyGetParamsWithHTTPClient(client *http.Client) *ProvisionerIAMPolicyGetParams {
	return &ProvisionerIAMPolicyGetParams{
		HTTPClient: client,
	}
}

/*
ProvisionerIAMPolicyGetParams contains all the parameters to send to the API endpoint

	for the provisioner i a m policy get operation.

	Typically these are written to a http.Request.
*/
type ProvisionerIAMPolicyGetParams struct {

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

// WithDefaults hydrates default values in the provisioner i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerIAMPolicyGetParams) WithDefaults() *ProvisionerIAMPolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provisioner i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerIAMPolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) WithTimeout(timeout time.Duration) *ProvisionerIAMPolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) WithContext(ctx context.Context) *ProvisionerIAMPolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) WithHTTPClient(client *http.Client) *ProvisionerIAMPolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *ProvisionerIAMPolicyGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) WithFullNameName(fullNameName string) *ProvisionerIAMPolicyGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) WithFullNameOrgID(fullNameOrgID *string) *ProvisionerIAMPolicyGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the provisioner i a m policy get params
func (o *ProvisionerIAMPolicyGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionerIAMPolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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