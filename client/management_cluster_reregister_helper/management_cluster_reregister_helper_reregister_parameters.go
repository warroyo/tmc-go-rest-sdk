// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_reregister_helper

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

// NewManagementClusterReregisterHelperReregisterParams creates a new ManagementClusterReregisterHelperReregisterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagementClusterReregisterHelperReregisterParams() *ManagementClusterReregisterHelperReregisterParams {
	return &ManagementClusterReregisterHelperReregisterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagementClusterReregisterHelperReregisterParamsWithTimeout creates a new ManagementClusterReregisterHelperReregisterParams object
// with the ability to set a timeout on a request.
func NewManagementClusterReregisterHelperReregisterParamsWithTimeout(timeout time.Duration) *ManagementClusterReregisterHelperReregisterParams {
	return &ManagementClusterReregisterHelperReregisterParams{
		timeout: timeout,
	}
}

// NewManagementClusterReregisterHelperReregisterParamsWithContext creates a new ManagementClusterReregisterHelperReregisterParams object
// with the ability to set a context for a request.
func NewManagementClusterReregisterHelperReregisterParamsWithContext(ctx context.Context) *ManagementClusterReregisterHelperReregisterParams {
	return &ManagementClusterReregisterHelperReregisterParams{
		Context: ctx,
	}
}

// NewManagementClusterReregisterHelperReregisterParamsWithHTTPClient creates a new ManagementClusterReregisterHelperReregisterParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagementClusterReregisterHelperReregisterParamsWithHTTPClient(client *http.Client) *ManagementClusterReregisterHelperReregisterParams {
	return &ManagementClusterReregisterHelperReregisterParams{
		HTTPClient: client,
	}
}

/*
ManagementClusterReregisterHelperReregisterParams contains all the parameters to send to the API endpoint

	for the management cluster reregister helper reregister operation.

	Typically these are written to a http.Request.
*/
type ManagementClusterReregisterHelperReregisterParams struct {

	// Body.
	Body *models.ManagementclusterManagementClusterReregisterRequest

	/* FullNameName.

	   Unique identifier of the ManagementCluster.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the management cluster reregister helper reregister params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterReregisterHelperReregisterParams) WithDefaults() *ManagementClusterReregisterHelperReregisterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the management cluster reregister helper reregister params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterReregisterHelperReregisterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) WithTimeout(timeout time.Duration) *ManagementClusterReregisterHelperReregisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) WithContext(ctx context.Context) *ManagementClusterReregisterHelperReregisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) WithHTTPClient(client *http.Client) *ManagementClusterReregisterHelperReregisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) WithBody(body *models.ManagementclusterManagementClusterReregisterRequest) *ManagementClusterReregisterHelperReregisterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) SetBody(body *models.ManagementclusterManagementClusterReregisterRequest) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) WithFullNameName(fullNameName string) *ManagementClusterReregisterHelperReregisterParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the management cluster reregister helper reregister params
func (o *ManagementClusterReregisterHelperReregisterParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *ManagementClusterReregisterHelperReregisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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