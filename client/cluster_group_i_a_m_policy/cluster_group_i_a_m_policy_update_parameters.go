// Code generated by go-swagger; DO NOT EDIT.

package cluster_group_i_a_m_policy

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

// NewClusterGroupIAMPolicyUpdateParams creates a new ClusterGroupIAMPolicyUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterGroupIAMPolicyUpdateParams() *ClusterGroupIAMPolicyUpdateParams {
	return &ClusterGroupIAMPolicyUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterGroupIAMPolicyUpdateParamsWithTimeout creates a new ClusterGroupIAMPolicyUpdateParams object
// with the ability to set a timeout on a request.
func NewClusterGroupIAMPolicyUpdateParamsWithTimeout(timeout time.Duration) *ClusterGroupIAMPolicyUpdateParams {
	return &ClusterGroupIAMPolicyUpdateParams{
		timeout: timeout,
	}
}

// NewClusterGroupIAMPolicyUpdateParamsWithContext creates a new ClusterGroupIAMPolicyUpdateParams object
// with the ability to set a context for a request.
func NewClusterGroupIAMPolicyUpdateParamsWithContext(ctx context.Context) *ClusterGroupIAMPolicyUpdateParams {
	return &ClusterGroupIAMPolicyUpdateParams{
		Context: ctx,
	}
}

// NewClusterGroupIAMPolicyUpdateParamsWithHTTPClient creates a new ClusterGroupIAMPolicyUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterGroupIAMPolicyUpdateParamsWithHTTPClient(client *http.Client) *ClusterGroupIAMPolicyUpdateParams {
	return &ClusterGroupIAMPolicyUpdateParams{
		HTTPClient: client,
	}
}

/*
ClusterGroupIAMPolicyUpdateParams contains all the parameters to send to the API endpoint

	for the cluster group i a m policy update operation.

	Typically these are written to a http.Request.
*/
type ClusterGroupIAMPolicyUpdateParams struct {

	/* Body.

	   ClusterGroup policy.
	*/
	Body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy

	/* FullNameName.

	   Name of this ClusterGroup.
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

// WithDefaults hydrates default values in the cluster group i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGroupIAMPolicyUpdateParams) WithDefaults() *ClusterGroupIAMPolicyUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster group i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGroupIAMPolicyUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) WithTimeout(timeout time.Duration) *ClusterGroupIAMPolicyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) WithContext(ctx context.Context) *ClusterGroupIAMPolicyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) WithHTTPClient(client *http.Client) *ClusterGroupIAMPolicyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) WithBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) *ClusterGroupIAMPolicyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) SetBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) {
	o.Body = body
}

// WithFullNameName adds the fullNameName to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) WithFullNameName(fullNameName string) *ClusterGroupIAMPolicyUpdateParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) WithFullNameOrgID(fullNameOrgID *string) *ClusterGroupIAMPolicyUpdateParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the cluster group i a m policy update params
func (o *ClusterGroupIAMPolicyUpdateParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterGroupIAMPolicyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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