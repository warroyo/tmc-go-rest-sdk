// Code generated by go-swagger; DO NOT EDIT.

package cluster_i_a_m_policy

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

// NewClusterIAMPolicyUpdateParams creates a new ClusterIAMPolicyUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterIAMPolicyUpdateParams() *ClusterIAMPolicyUpdateParams {
	return &ClusterIAMPolicyUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterIAMPolicyUpdateParamsWithTimeout creates a new ClusterIAMPolicyUpdateParams object
// with the ability to set a timeout on a request.
func NewClusterIAMPolicyUpdateParamsWithTimeout(timeout time.Duration) *ClusterIAMPolicyUpdateParams {
	return &ClusterIAMPolicyUpdateParams{
		timeout: timeout,
	}
}

// NewClusterIAMPolicyUpdateParamsWithContext creates a new ClusterIAMPolicyUpdateParams object
// with the ability to set a context for a request.
func NewClusterIAMPolicyUpdateParamsWithContext(ctx context.Context) *ClusterIAMPolicyUpdateParams {
	return &ClusterIAMPolicyUpdateParams{
		Context: ctx,
	}
}

// NewClusterIAMPolicyUpdateParamsWithHTTPClient creates a new ClusterIAMPolicyUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterIAMPolicyUpdateParamsWithHTTPClient(client *http.Client) *ClusterIAMPolicyUpdateParams {
	return &ClusterIAMPolicyUpdateParams{
		HTTPClient: client,
	}
}

/*
ClusterIAMPolicyUpdateParams contains all the parameters to send to the API endpoint

	for the cluster i a m policy update operation.

	Typically these are written to a http.Request.
*/
type ClusterIAMPolicyUpdateParams struct {

	/* Body.

	   Cluster policy.
	*/
	Body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy

	/* FullNameManagementClusterName.

	   Name of the management cluster.
	*/
	FullNameManagementClusterName *string

	/* FullNameName.

	   Name of this cluster.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameProvisionerName.

	   Provisioner of the cluster.
	*/
	FullNameProvisionerName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterIAMPolicyUpdateParams) WithDefaults() *ClusterIAMPolicyUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterIAMPolicyUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithTimeout(timeout time.Duration) *ClusterIAMPolicyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithContext(ctx context.Context) *ClusterIAMPolicyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithHTTPClient(client *http.Client) *ClusterIAMPolicyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) *ClusterIAMPolicyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) {
	o.Body = body
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *ClusterIAMPolicyUpdateParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithFullNameName(fullNameName string) *ClusterIAMPolicyUpdateParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithFullNameOrgID(fullNameOrgID *string) *ClusterIAMPolicyUpdateParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *ClusterIAMPolicyUpdateParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the cluster i a m policy update params
func (o *ClusterIAMPolicyUpdateParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterIAMPolicyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.FullNameManagementClusterName != nil {

		// query param fullName.managementClusterName
		var qrFullNameManagementClusterName string

		if o.FullNameManagementClusterName != nil {
			qrFullNameManagementClusterName = *o.FullNameManagementClusterName
		}
		qFullNameManagementClusterName := qrFullNameManagementClusterName
		if qFullNameManagementClusterName != "" {

			if err := r.SetQueryParam("fullName.managementClusterName", qFullNameManagementClusterName); err != nil {
				return err
			}
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

	if o.FullNameProvisionerName != nil {

		// query param fullName.provisionerName
		var qrFullNameProvisionerName string

		if o.FullNameProvisionerName != nil {
			qrFullNameProvisionerName = *o.FullNameProvisionerName
		}
		qFullNameProvisionerName := qrFullNameProvisionerName
		if qFullNameProvisionerName != "" {

			if err := r.SetQueryParam("fullName.provisionerName", qFullNameProvisionerName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}