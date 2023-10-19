// Code generated by go-swagger; DO NOT EDIT.

package workspace_i_a_m_policy

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

// NewWorkspaceIAMPolicyPatchParams creates a new WorkspaceIAMPolicyPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkspaceIAMPolicyPatchParams() *WorkspaceIAMPolicyPatchParams {
	return &WorkspaceIAMPolicyPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkspaceIAMPolicyPatchParamsWithTimeout creates a new WorkspaceIAMPolicyPatchParams object
// with the ability to set a timeout on a request.
func NewWorkspaceIAMPolicyPatchParamsWithTimeout(timeout time.Duration) *WorkspaceIAMPolicyPatchParams {
	return &WorkspaceIAMPolicyPatchParams{
		timeout: timeout,
	}
}

// NewWorkspaceIAMPolicyPatchParamsWithContext creates a new WorkspaceIAMPolicyPatchParams object
// with the ability to set a context for a request.
func NewWorkspaceIAMPolicyPatchParamsWithContext(ctx context.Context) *WorkspaceIAMPolicyPatchParams {
	return &WorkspaceIAMPolicyPatchParams{
		Context: ctx,
	}
}

// NewWorkspaceIAMPolicyPatchParamsWithHTTPClient creates a new WorkspaceIAMPolicyPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkspaceIAMPolicyPatchParamsWithHTTPClient(client *http.Client) *WorkspaceIAMPolicyPatchParams {
	return &WorkspaceIAMPolicyPatchParams{
		HTTPClient: client,
	}
}

/*
WorkspaceIAMPolicyPatchParams contains all the parameters to send to the API endpoint

	for the workspace i a m policy patch operation.

	Typically these are written to a http.Request.
*/
type WorkspaceIAMPolicyPatchParams struct {

	// Body.
	Body *models.WorkspacePatchWorkspaceIAMPolicyRequest

	/* FullNameName.

	   Name of this Workspace.
	*/
	FullNameName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workspace i a m policy patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceIAMPolicyPatchParams) WithDefaults() *WorkspaceIAMPolicyPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workspace i a m policy patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceIAMPolicyPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) WithTimeout(timeout time.Duration) *WorkspaceIAMPolicyPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) WithContext(ctx context.Context) *WorkspaceIAMPolicyPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) WithHTTPClient(client *http.Client) *WorkspaceIAMPolicyPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) WithBody(body *models.WorkspacePatchWorkspaceIAMPolicyRequest) *WorkspaceIAMPolicyPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) SetBody(body *models.WorkspacePatchWorkspaceIAMPolicyRequest) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) WithFullNameName(fullNameName string) *WorkspaceIAMPolicyPatchParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the workspace i a m policy patch params
func (o *WorkspaceIAMPolicyPatchParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WriteToRequest writes these params to a swagger request
func (o *WorkspaceIAMPolicyPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
