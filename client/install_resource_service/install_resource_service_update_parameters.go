// Code generated by go-swagger; DO NOT EDIT.

package install_resource_service

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

// NewInstallResourceServiceUpdateParams creates a new InstallResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstallResourceServiceUpdateParams() *InstallResourceServiceUpdateParams {
	return &InstallResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstallResourceServiceUpdateParamsWithTimeout creates a new InstallResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewInstallResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *InstallResourceServiceUpdateParams {
	return &InstallResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewInstallResourceServiceUpdateParamsWithContext creates a new InstallResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewInstallResourceServiceUpdateParamsWithContext(ctx context.Context) *InstallResourceServiceUpdateParams {
	return &InstallResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewInstallResourceServiceUpdateParamsWithHTTPClient creates a new InstallResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstallResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *InstallResourceServiceUpdateParams {
	return &InstallResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
InstallResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the install resource service update operation.

	Typically these are written to a http.Request.
*/
type InstallResourceServiceUpdateParams struct {

	// Body.
	Body *models.ClusterNamespaceTanzupackageInstallUpdateInstallRequest

	/* InstallFullNameClusterName.

	   Name of Cluster.
	*/
	InstallFullNameClusterName string

	/* InstallFullNameName.

	   Name of the Package Install.
	*/
	InstallFullNameName string

	/* InstallFullNameNamespaceName.

	   Name of Namespace.
	*/
	InstallFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the install resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallResourceServiceUpdateParams) WithDefaults() *InstallResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the install resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *InstallResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithContext(ctx context.Context) *InstallResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *InstallResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithBody(body *models.ClusterNamespaceTanzupackageInstallUpdateInstallRequest) *InstallResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetBody(body *models.ClusterNamespaceTanzupackageInstallUpdateInstallRequest) {
	o.Body = body
}

// WithInstallFullNameClusterName adds the installFullNameClusterName to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithInstallFullNameClusterName(installFullNameClusterName string) *InstallResourceServiceUpdateParams {
	o.SetInstallFullNameClusterName(installFullNameClusterName)
	return o
}

// SetInstallFullNameClusterName adds the installFullNameClusterName to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetInstallFullNameClusterName(installFullNameClusterName string) {
	o.InstallFullNameClusterName = installFullNameClusterName
}

// WithInstallFullNameName adds the installFullNameName to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithInstallFullNameName(installFullNameName string) *InstallResourceServiceUpdateParams {
	o.SetInstallFullNameName(installFullNameName)
	return o
}

// SetInstallFullNameName adds the installFullNameName to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetInstallFullNameName(installFullNameName string) {
	o.InstallFullNameName = installFullNameName
}

// WithInstallFullNameNamespaceName adds the installFullNameNamespaceName to the install resource service update params
func (o *InstallResourceServiceUpdateParams) WithInstallFullNameNamespaceName(installFullNameNamespaceName string) *InstallResourceServiceUpdateParams {
	o.SetInstallFullNameNamespaceName(installFullNameNamespaceName)
	return o
}

// SetInstallFullNameNamespaceName adds the installFullNameNamespaceName to the install resource service update params
func (o *InstallResourceServiceUpdateParams) SetInstallFullNameNamespaceName(installFullNameNamespaceName string) {
	o.InstallFullNameNamespaceName = installFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *InstallResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param install.fullName.clusterName
	if err := r.SetPathParam("install.fullName.clusterName", o.InstallFullNameClusterName); err != nil {
		return err
	}

	// path param install.fullName.name
	if err := r.SetPathParam("install.fullName.name", o.InstallFullNameName); err != nil {
		return err
	}

	// path param install.fullName.namespaceName
	if err := r.SetPathParam("install.fullName.namespaceName", o.InstallFullNameNamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}