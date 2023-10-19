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

// NewWorkspaceIAMPolicyUpdateParams creates a new WorkspaceIAMPolicyUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkspaceIAMPolicyUpdateParams() *WorkspaceIAMPolicyUpdateParams {
	return &WorkspaceIAMPolicyUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkspaceIAMPolicyUpdateParamsWithTimeout creates a new WorkspaceIAMPolicyUpdateParams object
// with the ability to set a timeout on a request.
func NewWorkspaceIAMPolicyUpdateParamsWithTimeout(timeout time.Duration) *WorkspaceIAMPolicyUpdateParams {
	return &WorkspaceIAMPolicyUpdateParams{
		timeout: timeout,
	}
}

// NewWorkspaceIAMPolicyUpdateParamsWithContext creates a new WorkspaceIAMPolicyUpdateParams object
// with the ability to set a context for a request.
func NewWorkspaceIAMPolicyUpdateParamsWithContext(ctx context.Context) *WorkspaceIAMPolicyUpdateParams {
	return &WorkspaceIAMPolicyUpdateParams{
		Context: ctx,
	}
}

// NewWorkspaceIAMPolicyUpdateParamsWithHTTPClient creates a new WorkspaceIAMPolicyUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkspaceIAMPolicyUpdateParamsWithHTTPClient(client *http.Client) *WorkspaceIAMPolicyUpdateParams {
	return &WorkspaceIAMPolicyUpdateParams{
		HTTPClient: client,
	}
}

/*
WorkspaceIAMPolicyUpdateParams contains all the parameters to send to the API endpoint

	for the workspace i a m policy update operation.

	Typically these are written to a http.Request.
*/
type WorkspaceIAMPolicyUpdateParams struct {

	/* Body.

	   Workspace policy.
	*/
	Body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy

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

// WithDefaults hydrates default values in the workspace i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceIAMPolicyUpdateParams) WithDefaults() *WorkspaceIAMPolicyUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workspace i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceIAMPolicyUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) WithTimeout(timeout time.Duration) *WorkspaceIAMPolicyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) WithContext(ctx context.Context) *WorkspaceIAMPolicyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) WithHTTPClient(client *http.Client) *WorkspaceIAMPolicyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) WithBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) *WorkspaceIAMPolicyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) SetBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) WithFullNameName(fullNameName string) *WorkspaceIAMPolicyUpdateParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) WithFullNameOrgID(fullNameOrgID *string) *WorkspaceIAMPolicyUpdateParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the workspace i a m policy update params
func (o *WorkspaceIAMPolicyUpdateParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkspaceIAMPolicyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
