// Code generated by go-swagger; DO NOT EDIT.

package policy_resource_service

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

// NewPolicyResourceServiceDeleteParams creates a new PolicyResourceServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPolicyResourceServiceDeleteParams() *PolicyResourceServiceDeleteParams {
	return &PolicyResourceServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPolicyResourceServiceDeleteParamsWithTimeout creates a new PolicyResourceServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewPolicyResourceServiceDeleteParamsWithTimeout(timeout time.Duration) *PolicyResourceServiceDeleteParams {
	return &PolicyResourceServiceDeleteParams{
		timeout: timeout,
	}
}

// NewPolicyResourceServiceDeleteParamsWithContext creates a new PolicyResourceServiceDeleteParams object
// with the ability to set a context for a request.
func NewPolicyResourceServiceDeleteParamsWithContext(ctx context.Context) *PolicyResourceServiceDeleteParams {
	return &PolicyResourceServiceDeleteParams{
		Context: ctx,
	}
}

// NewPolicyResourceServiceDeleteParamsWithHTTPClient creates a new PolicyResourceServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPolicyResourceServiceDeleteParamsWithHTTPClient(client *http.Client) *PolicyResourceServiceDeleteParams {
	return &PolicyResourceServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
PolicyResourceServiceDeleteParams contains all the parameters to send to the API endpoint

	for the policy resource service delete operation.

	Typically these are written to a http.Request.
*/
type PolicyResourceServiceDeleteParams struct {

	/* FullNameName.

	   Name of the policy.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameWorkspaceName.

	   Name of the workspace.
	*/
	FullNameWorkspaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the policy resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyResourceServiceDeleteParams) WithDefaults() *PolicyResourceServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the policy resource service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyResourceServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) WithTimeout(timeout time.Duration) *PolicyResourceServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) WithContext(ctx context.Context) *PolicyResourceServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) WithHTTPClient(client *http.Client) *PolicyResourceServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) WithFullNameName(fullNameName string) *PolicyResourceServiceDeleteParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) WithFullNameOrgID(fullNameOrgID *string) *PolicyResourceServiceDeleteParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameWorkspaceName adds the fullNameWorkspaceName to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) WithFullNameWorkspaceName(fullNameWorkspaceName string) *PolicyResourceServiceDeleteParams {
	o.SetFullNameWorkspaceName(fullNameWorkspaceName)
	return o
}

// SetFullNameWorkspaceName adds the fullNameWorkspaceName to the policy resource service delete params
func (o *PolicyResourceServiceDeleteParams) SetFullNameWorkspaceName(fullNameWorkspaceName string) {
	o.FullNameWorkspaceName = fullNameWorkspaceName
}

// WriteToRequest writes these params to a swagger request
func (o *PolicyResourceServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param fullName.workspaceName
	if err := r.SetPathParam("fullName.workspaceName", o.FullNameWorkspaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
