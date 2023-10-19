// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_options

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

// NewAksClusterOptionsListParams creates a new AksClusterOptionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAksClusterOptionsListParams() *AksClusterOptionsListParams {
	return &AksClusterOptionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAksClusterOptionsListParamsWithTimeout creates a new AksClusterOptionsListParams object
// with the ability to set a timeout on a request.
func NewAksClusterOptionsListParamsWithTimeout(timeout time.Duration) *AksClusterOptionsListParams {
	return &AksClusterOptionsListParams{
		timeout: timeout,
	}
}

// NewAksClusterOptionsListParamsWithContext creates a new AksClusterOptionsListParams object
// with the ability to set a context for a request.
func NewAksClusterOptionsListParamsWithContext(ctx context.Context) *AksClusterOptionsListParams {
	return &AksClusterOptionsListParams{
		Context: ctx,
	}
}

// NewAksClusterOptionsListParamsWithHTTPClient creates a new AksClusterOptionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAksClusterOptionsListParamsWithHTTPClient(client *http.Client) *AksClusterOptionsListParams {
	return &AksClusterOptionsListParams{
		HTTPClient: client,
	}
}

/*
AksClusterOptionsListParams contains all the parameters to send to the API endpoint

	for the aks cluster options list operation.

	Typically these are written to a http.Request.
*/
type AksClusterOptionsListParams struct {

	/* CredentialName.

	   Name of this credential.
	*/
	CredentialName *string

	/* OrgID.

	   ID of Organization.
	*/
	OrgID *string

	/* SubscriptionID.

	   ID of Azure subscription for the cluster options.
	*/
	SubscriptionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aks cluster options list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AksClusterOptionsListParams) WithDefaults() *AksClusterOptionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aks cluster options list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AksClusterOptionsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aks cluster options list params
func (o *AksClusterOptionsListParams) WithTimeout(timeout time.Duration) *AksClusterOptionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aks cluster options list params
func (o *AksClusterOptionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aks cluster options list params
func (o *AksClusterOptionsListParams) WithContext(ctx context.Context) *AksClusterOptionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aks cluster options list params
func (o *AksClusterOptionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aks cluster options list params
func (o *AksClusterOptionsListParams) WithHTTPClient(client *http.Client) *AksClusterOptionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aks cluster options list params
func (o *AksClusterOptionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialName adds the credentialName to the aks cluster options list params
func (o *AksClusterOptionsListParams) WithCredentialName(credentialName *string) *AksClusterOptionsListParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the aks cluster options list params
func (o *AksClusterOptionsListParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithOrgID adds the orgID to the aks cluster options list params
func (o *AksClusterOptionsListParams) WithOrgID(orgID *string) *AksClusterOptionsListParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the aks cluster options list params
func (o *AksClusterOptionsListParams) SetOrgID(orgID *string) {
	o.OrgID = orgID
}

// WithSubscriptionID adds the subscriptionID to the aks cluster options list params
func (o *AksClusterOptionsListParams) WithSubscriptionID(subscriptionID *string) *AksClusterOptionsListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the aks cluster options list params
func (o *AksClusterOptionsListParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *AksClusterOptionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CredentialName != nil {

		// query param credentialName
		var qrCredentialName string

		if o.CredentialName != nil {
			qrCredentialName = *o.CredentialName
		}
		qCredentialName := qrCredentialName
		if qCredentialName != "" {

			if err := r.SetQueryParam("credentialName", qCredentialName); err != nil {
				return err
			}
		}
	}

	if o.OrgID != nil {

		// query param orgId
		var qrOrgID string

		if o.OrgID != nil {
			qrOrgID = *o.OrgID
		}
		qOrgID := qrOrgID
		if qOrgID != "" {

			if err := r.SetQueryParam("orgId", qOrgID); err != nil {
				return err
			}
		}
	}

	if o.SubscriptionID != nil {

		// query param subscriptionId
		var qrSubscriptionID string

		if o.SubscriptionID != nil {
			qrSubscriptionID = *o.SubscriptionID
		}
		qSubscriptionID := qrSubscriptionID
		if qSubscriptionID != "" {

			if err := r.SetQueryParam("subscriptionId", qSubscriptionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
