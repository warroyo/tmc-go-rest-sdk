// Code generated by go-swagger; DO NOT EDIT.

package eks_cluster_i_a_m_policy

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

// NewEksClusterIAMPolicyTestPermissionsParams creates a new EksClusterIAMPolicyTestPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEksClusterIAMPolicyTestPermissionsParams() *EksClusterIAMPolicyTestPermissionsParams {
	return &EksClusterIAMPolicyTestPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEksClusterIAMPolicyTestPermissionsParamsWithTimeout creates a new EksClusterIAMPolicyTestPermissionsParams object
// with the ability to set a timeout on a request.
func NewEksClusterIAMPolicyTestPermissionsParamsWithTimeout(timeout time.Duration) *EksClusterIAMPolicyTestPermissionsParams {
	return &EksClusterIAMPolicyTestPermissionsParams{
		timeout: timeout,
	}
}

// NewEksClusterIAMPolicyTestPermissionsParamsWithContext creates a new EksClusterIAMPolicyTestPermissionsParams object
// with the ability to set a context for a request.
func NewEksClusterIAMPolicyTestPermissionsParamsWithContext(ctx context.Context) *EksClusterIAMPolicyTestPermissionsParams {
	return &EksClusterIAMPolicyTestPermissionsParams{
		Context: ctx,
	}
}

// NewEksClusterIAMPolicyTestPermissionsParamsWithHTTPClient creates a new EksClusterIAMPolicyTestPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewEksClusterIAMPolicyTestPermissionsParamsWithHTTPClient(client *http.Client) *EksClusterIAMPolicyTestPermissionsParams {
	return &EksClusterIAMPolicyTestPermissionsParams{
		HTTPClient: client,
	}
}

/*
EksClusterIAMPolicyTestPermissionsParams contains all the parameters to send to the API endpoint

	for the eks cluster i a m policy test permissions operation.

	Typically these are written to a http.Request.
*/
type EksClusterIAMPolicyTestPermissionsParams struct {

	// Body.
	Body *models.EksclusterTestEksClusterIAMPermissionsRequest

	/* FullNameName.

	   Name of this cluster.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eks cluster i a m policy test permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterIAMPolicyTestPermissionsParams) WithDefaults() *EksClusterIAMPolicyTestPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eks cluster i a m policy test permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterIAMPolicyTestPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) WithTimeout(timeout time.Duration) *EksClusterIAMPolicyTestPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) WithContext(ctx context.Context) *EksClusterIAMPolicyTestPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) WithHTTPClient(client *http.Client) *EksClusterIAMPolicyTestPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) WithBody(body *models.EksclusterTestEksClusterIAMPermissionsRequest) *EksClusterIAMPolicyTestPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) SetBody(body *models.EksclusterTestEksClusterIAMPermissionsRequest) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) WithFullNameName(fullNameName string) *EksClusterIAMPolicyTestPermissionsParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the eks cluster i a m policy test permissions params
func (o *EksClusterIAMPolicyTestPermissionsParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *EksClusterIAMPolicyTestPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
