// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_i_a_m_policy

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

// NewManagementClusterIAMPolicyUpdateParams creates a new ManagementClusterIAMPolicyUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagementClusterIAMPolicyUpdateParams() *ManagementClusterIAMPolicyUpdateParams {
	return &ManagementClusterIAMPolicyUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagementClusterIAMPolicyUpdateParamsWithTimeout creates a new ManagementClusterIAMPolicyUpdateParams object
// with the ability to set a timeout on a request.
func NewManagementClusterIAMPolicyUpdateParamsWithTimeout(timeout time.Duration) *ManagementClusterIAMPolicyUpdateParams {
	return &ManagementClusterIAMPolicyUpdateParams{
		timeout: timeout,
	}
}

// NewManagementClusterIAMPolicyUpdateParamsWithContext creates a new ManagementClusterIAMPolicyUpdateParams object
// with the ability to set a context for a request.
func NewManagementClusterIAMPolicyUpdateParamsWithContext(ctx context.Context) *ManagementClusterIAMPolicyUpdateParams {
	return &ManagementClusterIAMPolicyUpdateParams{
		Context: ctx,
	}
}

// NewManagementClusterIAMPolicyUpdateParamsWithHTTPClient creates a new ManagementClusterIAMPolicyUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagementClusterIAMPolicyUpdateParamsWithHTTPClient(client *http.Client) *ManagementClusterIAMPolicyUpdateParams {
	return &ManagementClusterIAMPolicyUpdateParams{
		HTTPClient: client,
	}
}

/*
ManagementClusterIAMPolicyUpdateParams contains all the parameters to send to the API endpoint

	for the management cluster i a m policy update operation.

	Typically these are written to a http.Request.
*/
type ManagementClusterIAMPolicyUpdateParams struct {

	/* Body.

	   ManagementCluster policy.
	*/
	Body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy

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

// WithDefaults hydrates default values in the management cluster i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterIAMPolicyUpdateParams) WithDefaults() *ManagementClusterIAMPolicyUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the management cluster i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagementClusterIAMPolicyUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) WithTimeout(timeout time.Duration) *ManagementClusterIAMPolicyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) WithContext(ctx context.Context) *ManagementClusterIAMPolicyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) WithHTTPClient(client *http.Client) *ManagementClusterIAMPolicyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) WithBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) *ManagementClusterIAMPolicyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) SetBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) WithFullNameName(fullNameName string) *ManagementClusterIAMPolicyUpdateParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) WithFullNameOrgID(fullNameOrgID *string) *ManagementClusterIAMPolicyUpdateParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the management cluster i a m policy update params
func (o *ManagementClusterIAMPolicyUpdateParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ManagementClusterIAMPolicyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
