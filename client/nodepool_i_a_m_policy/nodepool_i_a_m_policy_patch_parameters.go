// Code generated by go-swagger; DO NOT EDIT.

package nodepool_i_a_m_policy

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

// NewNodepoolIAMPolicyPatchParams creates a new NodepoolIAMPolicyPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodepoolIAMPolicyPatchParams() *NodepoolIAMPolicyPatchParams {
	return &NodepoolIAMPolicyPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodepoolIAMPolicyPatchParamsWithTimeout creates a new NodepoolIAMPolicyPatchParams object
// with the ability to set a timeout on a request.
func NewNodepoolIAMPolicyPatchParamsWithTimeout(timeout time.Duration) *NodepoolIAMPolicyPatchParams {
	return &NodepoolIAMPolicyPatchParams{
		timeout: timeout,
	}
}

// NewNodepoolIAMPolicyPatchParamsWithContext creates a new NodepoolIAMPolicyPatchParams object
// with the ability to set a context for a request.
func NewNodepoolIAMPolicyPatchParamsWithContext(ctx context.Context) *NodepoolIAMPolicyPatchParams {
	return &NodepoolIAMPolicyPatchParams{
		Context: ctx,
	}
}

// NewNodepoolIAMPolicyPatchParamsWithHTTPClient creates a new NodepoolIAMPolicyPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodepoolIAMPolicyPatchParamsWithHTTPClient(client *http.Client) *NodepoolIAMPolicyPatchParams {
	return &NodepoolIAMPolicyPatchParams{
		HTTPClient: client,
	}
}

/*
NodepoolIAMPolicyPatchParams contains all the parameters to send to the API endpoint

	for the nodepool i a m policy patch operation.

	Typically these are written to a http.Request.
*/
type NodepoolIAMPolicyPatchParams struct {

	// Body.
	Body *models.EksclusterNodepoolPatchNodepoolIAMPolicyRequest

	/* FullNameEksClusterName.

	   Name of the cluster.
	*/
	FullNameEksClusterName string

	/* FullNameName.

	   Name of the nodepool.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nodepool i a m policy patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolIAMPolicyPatchParams) WithDefaults() *NodepoolIAMPolicyPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodepool i a m policy patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolIAMPolicyPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) WithTimeout(timeout time.Duration) *NodepoolIAMPolicyPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) WithContext(ctx context.Context) *NodepoolIAMPolicyPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) WithHTTPClient(client *http.Client) *NodepoolIAMPolicyPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) WithBody(body *models.EksclusterNodepoolPatchNodepoolIAMPolicyRequest) *NodepoolIAMPolicyPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) SetBody(body *models.EksclusterNodepoolPatchNodepoolIAMPolicyRequest) {
	o.Body = body
}

// WithFullNameEksClusterName adds the fullNameEksClusterName to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) WithFullNameEksClusterName(fullNameEksClusterName string) *NodepoolIAMPolicyPatchParams {
	o.SetFullNameEksClusterName(fullNameEksClusterName)
	return o
}

// SetFullNameEksClusterName adds the fullNameEksClusterName to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) SetFullNameEksClusterName(fullNameEksClusterName string) {
	o.FullNameEksClusterName = fullNameEksClusterName
}

// WithFullNameName adds the fullNameName to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) WithFullNameName(fullNameName string) *NodepoolIAMPolicyPatchParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the nodepool i a m policy patch params
func (o *NodepoolIAMPolicyPatchParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *NodepoolIAMPolicyPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param fullName.eksClusterName
	if err := r.SetPathParam("fullName.eksClusterName", o.FullNameEksClusterName); err != nil {
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