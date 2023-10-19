// Code generated by go-swagger; DO NOT EDIT.

package eks_cluster_apply_helper

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

// NewEksClusterApplyHelperApplyParams creates a new EksClusterApplyHelperApplyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEksClusterApplyHelperApplyParams() *EksClusterApplyHelperApplyParams {
	return &EksClusterApplyHelperApplyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEksClusterApplyHelperApplyParamsWithTimeout creates a new EksClusterApplyHelperApplyParams object
// with the ability to set a timeout on a request.
func NewEksClusterApplyHelperApplyParamsWithTimeout(timeout time.Duration) *EksClusterApplyHelperApplyParams {
	return &EksClusterApplyHelperApplyParams{
		timeout: timeout,
	}
}

// NewEksClusterApplyHelperApplyParamsWithContext creates a new EksClusterApplyHelperApplyParams object
// with the ability to set a context for a request.
func NewEksClusterApplyHelperApplyParamsWithContext(ctx context.Context) *EksClusterApplyHelperApplyParams {
	return &EksClusterApplyHelperApplyParams{
		Context: ctx,
	}
}

// NewEksClusterApplyHelperApplyParamsWithHTTPClient creates a new EksClusterApplyHelperApplyParams object
// with the ability to set a custom HTTPClient for a request.
func NewEksClusterApplyHelperApplyParamsWithHTTPClient(client *http.Client) *EksClusterApplyHelperApplyParams {
	return &EksClusterApplyHelperApplyParams{
		HTTPClient: client,
	}
}

/*
EksClusterApplyHelperApplyParams contains all the parameters to send to the API endpoint

	for the eks cluster apply helper apply operation.

	Typically these are written to a http.Request.
*/
type EksClusterApplyHelperApplyParams struct {

	// Body.
	Body *models.EksclusterApplyEksClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eks cluster apply helper apply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterApplyHelperApplyParams) WithDefaults() *EksClusterApplyHelperApplyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eks cluster apply helper apply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterApplyHelperApplyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) WithTimeout(timeout time.Duration) *EksClusterApplyHelperApplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) WithContext(ctx context.Context) *EksClusterApplyHelperApplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) WithHTTPClient(client *http.Client) *EksClusterApplyHelperApplyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) WithBody(body *models.EksclusterApplyEksClusterRequest) *EksClusterApplyHelperApplyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the eks cluster apply helper apply params
func (o *EksClusterApplyHelperApplyParams) SetBody(body *models.EksclusterApplyEksClusterRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EksClusterApplyHelperApplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
