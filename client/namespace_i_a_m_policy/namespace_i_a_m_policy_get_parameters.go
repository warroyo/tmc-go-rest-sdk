// Code generated by go-swagger; DO NOT EDIT.

package namespace_i_a_m_policy

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

// NewNamespaceIAMPolicyGetParams creates a new NamespaceIAMPolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNamespaceIAMPolicyGetParams() *NamespaceIAMPolicyGetParams {
	return &NamespaceIAMPolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNamespaceIAMPolicyGetParamsWithTimeout creates a new NamespaceIAMPolicyGetParams object
// with the ability to set a timeout on a request.
func NewNamespaceIAMPolicyGetParamsWithTimeout(timeout time.Duration) *NamespaceIAMPolicyGetParams {
	return &NamespaceIAMPolicyGetParams{
		timeout: timeout,
	}
}

// NewNamespaceIAMPolicyGetParamsWithContext creates a new NamespaceIAMPolicyGetParams object
// with the ability to set a context for a request.
func NewNamespaceIAMPolicyGetParamsWithContext(ctx context.Context) *NamespaceIAMPolicyGetParams {
	return &NamespaceIAMPolicyGetParams{
		Context: ctx,
	}
}

// NewNamespaceIAMPolicyGetParamsWithHTTPClient creates a new NamespaceIAMPolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNamespaceIAMPolicyGetParamsWithHTTPClient(client *http.Client) *NamespaceIAMPolicyGetParams {
	return &NamespaceIAMPolicyGetParams{
		HTTPClient: client,
	}
}

/*
NamespaceIAMPolicyGetParams contains all the parameters to send to the API endpoint

	for the namespace i a m policy get operation.

	Typically these are written to a http.Request.
*/
type NamespaceIAMPolicyGetParams struct {

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

// WithDefaults hydrates default values in the namespace i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceIAMPolicyGetParams) WithDefaults() *NamespaceIAMPolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the namespace i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceIAMPolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithTimeout(timeout time.Duration) *NamespaceIAMPolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithContext(ctx context.Context) *NamespaceIAMPolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithHTTPClient(client *http.Client) *NamespaceIAMPolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameClusterName adds the fullNameClusterName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithFullNameClusterName(fullNameClusterName string) *NamespaceIAMPolicyGetParams {
	o.SetFullNameClusterName(fullNameClusterName)
	return o
}

// SetFullNameClusterName adds the fullNameClusterName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetFullNameClusterName(fullNameClusterName string) {
	o.FullNameClusterName = fullNameClusterName
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *NamespaceIAMPolicyGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithFullNameName(fullNameName string) *NamespaceIAMPolicyGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithFullNameOrgID(fullNameOrgID *string) *NamespaceIAMPolicyGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *NamespaceIAMPolicyGetParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the namespace i a m policy get params
func (o *NamespaceIAMPolicyGetParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *NamespaceIAMPolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
