// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_i_a_m_policy

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

// NewAksClusterIAMPolicyUpdateParams creates a new AksClusterIAMPolicyUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAksClusterIAMPolicyUpdateParams() *AksClusterIAMPolicyUpdateParams {
	return &AksClusterIAMPolicyUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAksClusterIAMPolicyUpdateParamsWithTimeout creates a new AksClusterIAMPolicyUpdateParams object
// with the ability to set a timeout on a request.
func NewAksClusterIAMPolicyUpdateParamsWithTimeout(timeout time.Duration) *AksClusterIAMPolicyUpdateParams {
	return &AksClusterIAMPolicyUpdateParams{
		timeout: timeout,
	}
}

// NewAksClusterIAMPolicyUpdateParamsWithContext creates a new AksClusterIAMPolicyUpdateParams object
// with the ability to set a context for a request.
func NewAksClusterIAMPolicyUpdateParamsWithContext(ctx context.Context) *AksClusterIAMPolicyUpdateParams {
	return &AksClusterIAMPolicyUpdateParams{
		Context: ctx,
	}
}

// NewAksClusterIAMPolicyUpdateParamsWithHTTPClient creates a new AksClusterIAMPolicyUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAksClusterIAMPolicyUpdateParamsWithHTTPClient(client *http.Client) *AksClusterIAMPolicyUpdateParams {
	return &AksClusterIAMPolicyUpdateParams{
		HTTPClient: client,
	}
}

/*
AksClusterIAMPolicyUpdateParams contains all the parameters to send to the API endpoint

	for the aks cluster i a m policy update operation.

	Typically these are written to a http.Request.
*/
type AksClusterIAMPolicyUpdateParams struct {

	/* Body.

	   AksCluster policy.
	*/
	Body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy

	/* FullNameCredentialName.

	   Name of the credential.
	*/
	FullNameCredentialName *string

	/* FullNameName.

	   Name of this cluster.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameResourceGroupName.

	   Name of the resource group.
	*/
	FullNameResourceGroupName *string

	/* FullNameSubscriptionID.

	   ID of Azure subscription of the cluster.
	*/
	FullNameSubscriptionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aks cluster i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AksClusterIAMPolicyUpdateParams) WithDefaults() *AksClusterIAMPolicyUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aks cluster i a m policy update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AksClusterIAMPolicyUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithTimeout(timeout time.Duration) *AksClusterIAMPolicyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithContext(ctx context.Context) *AksClusterIAMPolicyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithHTTPClient(client *http.Client) *AksClusterIAMPolicyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) *AksClusterIAMPolicyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetBody(body *models.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) {
	o.Body = body
}

// WithFullNameCredentialName adds the fullNameCredentialName to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithFullNameCredentialName(fullNameCredentialName *string) *AksClusterIAMPolicyUpdateParams {
	o.SetFullNameCredentialName(fullNameCredentialName)
	return o
}

// SetFullNameCredentialName adds the fullNameCredentialName to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetFullNameCredentialName(fullNameCredentialName *string) {
	o.FullNameCredentialName = fullNameCredentialName
}

// WithFullNameName adds the fullNameName to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithFullNameName(fullNameName string) *AksClusterIAMPolicyUpdateParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithFullNameOrgID(fullNameOrgID *string) *AksClusterIAMPolicyUpdateParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameResourceGroupName adds the fullNameResourceGroupName to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithFullNameResourceGroupName(fullNameResourceGroupName *string) *AksClusterIAMPolicyUpdateParams {
	o.SetFullNameResourceGroupName(fullNameResourceGroupName)
	return o
}

// SetFullNameResourceGroupName adds the fullNameResourceGroupName to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetFullNameResourceGroupName(fullNameResourceGroupName *string) {
	o.FullNameResourceGroupName = fullNameResourceGroupName
}

// WithFullNameSubscriptionID adds the fullNameSubscriptionID to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) WithFullNameSubscriptionID(fullNameSubscriptionID *string) *AksClusterIAMPolicyUpdateParams {
	o.SetFullNameSubscriptionID(fullNameSubscriptionID)
	return o
}

// SetFullNameSubscriptionID adds the fullNameSubscriptionId to the aks cluster i a m policy update params
func (o *AksClusterIAMPolicyUpdateParams) SetFullNameSubscriptionID(fullNameSubscriptionID *string) {
	o.FullNameSubscriptionID = fullNameSubscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *AksClusterIAMPolicyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.FullNameCredentialName != nil {

		// query param fullName.credentialName
		var qrFullNameCredentialName string

		if o.FullNameCredentialName != nil {
			qrFullNameCredentialName = *o.FullNameCredentialName
		}
		qFullNameCredentialName := qrFullNameCredentialName
		if qFullNameCredentialName != "" {

			if err := r.SetQueryParam("fullName.credentialName", qFullNameCredentialName); err != nil {
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

	if o.FullNameResourceGroupName != nil {

		// query param fullName.resourceGroupName
		var qrFullNameResourceGroupName string

		if o.FullNameResourceGroupName != nil {
			qrFullNameResourceGroupName = *o.FullNameResourceGroupName
		}
		qFullNameResourceGroupName := qrFullNameResourceGroupName
		if qFullNameResourceGroupName != "" {

			if err := r.SetQueryParam("fullName.resourceGroupName", qFullNameResourceGroupName); err != nil {
				return err
			}
		}
	}

	if o.FullNameSubscriptionID != nil {

		// query param fullName.subscriptionId
		var qrFullNameSubscriptionID string

		if o.FullNameSubscriptionID != nil {
			qrFullNameSubscriptionID = *o.FullNameSubscriptionID
		}
		qFullNameSubscriptionID := qrFullNameSubscriptionID
		if qFullNameSubscriptionID != "" {

			if err := r.SetQueryParam("fullName.subscriptionId", qFullNameSubscriptionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
