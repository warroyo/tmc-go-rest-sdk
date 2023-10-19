// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_manifest_helper

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

// NewManagementClusterManifestHelperGetManifestParams creates a new ManagementClusterManifestHelperGetManifestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagementClusterManifestHelperGetManifestParams() *ManagementClusterManifestHelperGetManifestParams {
	return &ManagementClusterManifestHelperGetManifestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagementClusterManifestHelperGetManifestParamsWithTimeout creates a new ManagementClusterManifestHelperGetManifestParams object
// with the ability to set a timeout on a request.
func NewManagementClusterManifestHelperGetManifestParamsWithTimeout(timeout time.Duration) *ManagementClusterManifestHelperGetManifestParams {
	return &ManagementClusterManifestHelperGetManifestParams{
		timeout: timeout,
	}
}

// NewManagementClusterManifestHelperGetManifestParamsWithContext creates a new ManagementClusterManifestHelperGetManifestParams object
// with the ability to set a context for a request.
func NewManagementClusterManifestHelperGetManifestParamsWithContext(ctx context.Context) *ManagementClusterManifestHelperGetManifestParams {
	return &ManagementClusterManifestHelperGetManifestParams{
		Context: ctx,
	}
}

// NewManagementClusterManifestHelperGetManifestParamsWithHTTPClient creates a new ManagementClusterManifestHelperGetManifestParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagementClusterManifestHelperGetManifestParamsWithHTTPClient(client *http.Client) *ManagementClusterManifestHelperGetManifestParams {
	return &ManagementClusterManifestHelperGetManifestParams{
		HTTPClient: client,
	}
}

/*
ManagementClusterManifestHelperGetManifestParams contains all the parameters to send to the API endpoint

	for the management cluster manifest helper get manifest operation.

	Typically these are written to a http.Request.
*/
type ManagementClusterManifestHelperGetManifestParams struct {

	/* FullNameName.

	   Unique identifier of the ManagementCluster.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization. Generally a GUID.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the management cluster manifest helper get manifest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterManifestHelperGetManifestParams) WithDefaults() *ManagementClusterManifestHelperGetManifestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the management cluster manifest helper get manifest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterManifestHelperGetManifestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) WithTimeout(timeout time.Duration) *ManagementClusterManifestHelperGetManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) WithContext(ctx context.Context) *ManagementClusterManifestHelperGetManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) WithHTTPClient(client *http.Client) *ManagementClusterManifestHelperGetManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) WithFullNameName(fullNameName string) *ManagementClusterManifestHelperGetManifestParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) WithFullNameOrgID(fullNameOrgID *string) *ManagementClusterManifestHelperGetManifestParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the management cluster manifest helper get manifest params
func (o *ManagementClusterManifestHelperGetManifestParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ManagementClusterManifestHelperGetManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
