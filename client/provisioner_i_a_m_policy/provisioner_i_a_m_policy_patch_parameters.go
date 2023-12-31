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

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// NewProvisionerIAMPolicyPatchParams creates a new ProvisionerIAMPolicyPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionerIAMPolicyPatchParams() *ProvisionerIAMPolicyPatchParams {
	return &ProvisionerIAMPolicyPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionerIAMPolicyPatchParamsWithTimeout creates a new ProvisionerIAMPolicyPatchParams object
// with the ability to set a timeout on a request.
func NewProvisionerIAMPolicyPatchParamsWithTimeout(timeout time.Duration) *ProvisionerIAMPolicyPatchParams {
	return &ProvisionerIAMPolicyPatchParams{
		timeout: timeout,
	}
}

// NewProvisionerIAMPolicyPatchParamsWithContext creates a new ProvisionerIAMPolicyPatchParams object
// with the ability to set a context for a request.
func NewProvisionerIAMPolicyPatchParamsWithContext(ctx context.Context) *ProvisionerIAMPolicyPatchParams {
	return &ProvisionerIAMPolicyPatchParams{
		Context: ctx,
	}
}

// NewProvisionerIAMPolicyPatchParamsWithHTTPClient creates a new ProvisionerIAMPolicyPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionerIAMPolicyPatchParamsWithHTTPClient(client *http.Client) *ProvisionerIAMPolicyPatchParams {
	return &ProvisionerIAMPolicyPatchParams{
		HTTPClient: client,
	}
}

/*
ProvisionerIAMPolicyPatchParams contains all the parameters to send to the API endpoint

	for the provisioner i a m policy patch operation.

	Typically these are written to a http.Request.
*/
type ProvisionerIAMPolicyPatchParams struct {

	// Body.
	Body *models.ManagementclusterProvisionerPatchProvisionerIAMPolicyRequest

	/* FullNameManagementClusterName.

	   Name of the ManagementCluster.
	*/
	FullNameManagementClusterName string

	/* FullNameName.

	   Name of the provisioner. It must be unique within a management cluster.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provisioner i a m policy patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerIAMPolicyPatchParams) WithDefaults() *ProvisionerIAMPolicyPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provisioner i a m policy patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerIAMPolicyPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) WithTimeout(timeout time.Duration) *ProvisionerIAMPolicyPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) WithContext(ctx context.Context) *ProvisionerIAMPolicyPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) WithHTTPClient(client *http.Client) *ProvisionerIAMPolicyPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) WithBody(body *models.ManagementclusterProvisionerPatchProvisionerIAMPolicyRequest) *ProvisionerIAMPolicyPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) SetBody(body *models.ManagementclusterProvisionerPatchProvisionerIAMPolicyRequest) {
	o.Body = body
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) WithFullNameManagementClusterName(fullNameManagementClusterName string) *ProvisionerIAMPolicyPatchParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) SetFullNameManagementClusterName(fullNameManagementClusterName string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) WithFullNameName(fullNameName string) *ProvisionerIAMPolicyPatchParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the provisioner i a m policy patch params
func (o *ProvisionerIAMPolicyPatchParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionerIAMPolicyPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param fullName.managementClusterName
	if err := r.SetPathParam("fullName.managementClusterName", o.FullNameManagementClusterName); err != nil {
		return err
	}

	// path param fullName.name
	if err := r.SetPathParam("fullName.name", o.FullNameName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
