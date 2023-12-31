// Code generated by go-swagger; DO NOT EDIT.

package provisioner_resource_service

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

// NewProvisionerResourceServiceUpdateParams creates a new ProvisionerResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionerResourceServiceUpdateParams() *ProvisionerResourceServiceUpdateParams {
	return &ProvisionerResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionerResourceServiceUpdateParamsWithTimeout creates a new ProvisionerResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewProvisionerResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *ProvisionerResourceServiceUpdateParams {
	return &ProvisionerResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewProvisionerResourceServiceUpdateParamsWithContext creates a new ProvisionerResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewProvisionerResourceServiceUpdateParamsWithContext(ctx context.Context) *ProvisionerResourceServiceUpdateParams {
	return &ProvisionerResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewProvisionerResourceServiceUpdateParamsWithHTTPClient creates a new ProvisionerResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionerResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *ProvisionerResourceServiceUpdateParams {
	return &ProvisionerResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
ProvisionerResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the provisioner resource service update operation.

	Typically these are written to a http.Request.
*/
type ProvisionerResourceServiceUpdateParams struct {

	// Body.
	Body *models.ManagementclusterProvisionerUpdateProvisionerRequest

	/* ProvisionerFullNameManagementClusterName.

	   Name of the ManagementCluster.
	*/
	ProvisionerFullNameManagementClusterName string

	/* ProvisionerFullNameName.

	   Name of the provisioner. It must be unique within a management cluster.
	*/
	ProvisionerFullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provisioner resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerResourceServiceUpdateParams) WithDefaults() *ProvisionerResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provisioner resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionerResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *ProvisionerResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) WithContext(ctx context.Context) *ProvisionerResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *ProvisionerResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) WithBody(body *models.ManagementclusterProvisionerUpdateProvisionerRequest) *ProvisionerResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) SetBody(body *models.ManagementclusterProvisionerUpdateProvisionerRequest) {
	o.Body = body
}

// WithProvisionerFullNameManagementClusterName adds the provisionerFullNameManagementClusterName to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) WithProvisionerFullNameManagementClusterName(provisionerFullNameManagementClusterName string) *ProvisionerResourceServiceUpdateParams {
	o.SetProvisionerFullNameManagementClusterName(provisionerFullNameManagementClusterName)
	return o
}

// SetProvisionerFullNameManagementClusterName adds the provisionerFullNameManagementClusterName to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) SetProvisionerFullNameManagementClusterName(provisionerFullNameManagementClusterName string) {
	o.ProvisionerFullNameManagementClusterName = provisionerFullNameManagementClusterName
}

// WithProvisionerFullNameName adds the provisionerFullNameName to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) WithProvisionerFullNameName(provisionerFullNameName string) *ProvisionerResourceServiceUpdateParams {
	o.SetProvisionerFullNameName(provisionerFullNameName)
	return o
}

// SetProvisionerFullNameName adds the provisionerFullNameName to the provisioner resource service update params
func (o *ProvisionerResourceServiceUpdateParams) SetProvisionerFullNameName(provisionerFullNameName string) {
	o.ProvisionerFullNameName = provisionerFullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionerResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param provisioner.fullName.managementClusterName
	if err := r.SetPathParam("provisioner.fullName.managementClusterName", o.ProvisionerFullNameManagementClusterName); err != nil {
		return err
	}

	// path param provisioner.fullName.name
	if err := r.SetPathParam("provisioner.fullName.name", o.ProvisionerFullNameName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
