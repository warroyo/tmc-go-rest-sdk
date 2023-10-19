// Code generated by go-swagger; DO NOT EDIT.

package workspace_resource_service

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

// NewWorkspaceResourceServiceDeleteParams creates a new WorkspaceResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkspaceResourceServiceDeleteParams() *WorkspaceResourceServiceDeleteParams {
	return &WorkspaceResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkspaceResourceServiceDeleteParamsWithTimeout creates a new WorkspaceResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewWorkspaceResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *WorkspaceResourceServiceDeleteParams {
	return &WorkspaceResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewWorkspaceResourceServiceDeleteParamsWithContext creates a new WorkspaceResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewWorkspaceResourceServiceDeleteParamsWithContext(ctx context.Context) *WorkspaceResourceServiceDeleteParams {
	return &WorkspaceResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewWorkspaceResourceServiceDeleteParamsWithHTTPClient creates a new WorkspaceResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkspaceResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *WorkspaceResourceServiceDeleteParams {
	return &WorkspaceResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
WorkspaceResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the workspace resource service delete operation.

	Typically these are written to a http.Request.
*/
type WorkspaceResourceServiceDeleteParams struct {

	/* FullNameName.

	   Name of this Workspace.
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

// WithDefaults hydrates default values in the workspace resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceResourceServiceDeleteParams) WithDefaults() *WorkspaceResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workspace resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *WorkspaceResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) WithContext(ctx context.Context) *WorkspaceResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *WorkspaceResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) WithFullNameName(fullNameName string) *WorkspaceResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *WorkspaceResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the workspace resource service delete params
func (o *WorkspaceResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkspaceResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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