// Code generated by go-swagger; DO NOT EDIT.

package provider_aks_cluster_resource_service

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

// NewProviderAksClusterResourceServiceGetParams creates a new ProviderAksClusterResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProviderAksClusterResourceServiceGetParams() *ProviderAksClusterResourceServiceGetParams {
	return &ProviderAksClusterResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProviderAksClusterResourceServiceGetParamsWithTimeout creates a new ProviderAksClusterResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewProviderAksClusterResourceServiceGetParamsWithTimeout(timeout time.Duration) *ProviderAksClusterResourceServiceGetParams {
	return &ProviderAksClusterResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewProviderAksClusterResourceServiceGetParamsWithContext creates a new ProviderAksClusterResourceServiceGetParams object
// with the ability to set a context for a request.
func NewProviderAksClusterResourceServiceGetParamsWithContext(ctx context.Context) *ProviderAksClusterResourceServiceGetParams {
	return &ProviderAksClusterResourceServiceGetParams{
		Context: ctx,
	}
}

// NewProviderAksClusterResourceServiceGetParamsWithHTTPClient creates a new ProviderAksClusterResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewProviderAksClusterResourceServiceGetParamsWithHTTPClient(client *http.Client) *ProviderAksClusterResourceServiceGetParams {
	return &ProviderAksClusterResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
ProviderAksClusterResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the provider aks cluster resource service get operation.

	Typically these are written to a http.Request.
*/
type ProviderAksClusterResourceServiceGetParams struct {

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

// WithDefaults hydrates default values in the provider aks cluster resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderAksClusterResourceServiceGetParams) WithDefaults() *ProviderAksClusterResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provider aks cluster resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderAksClusterResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithTimeout(timeout time.Duration) *ProviderAksClusterResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithContext(ctx context.Context) *ProviderAksClusterResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithHTTPClient(client *http.Client) *ProviderAksClusterResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameCredentialName adds the fullNameCredentialName to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithFullNameCredentialName(fullNameCredentialName *string) *ProviderAksClusterResourceServiceGetParams {
	o.SetFullNameCredentialName(fullNameCredentialName)
	return o
}

// SetFullNameCredentialName adds the fullNameCredentialName to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetFullNameCredentialName(fullNameCredentialName *string) {
	o.FullNameCredentialName = fullNameCredentialName
}

// WithFullNameName adds the fullNameName to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithFullNameName(fullNameName string) *ProviderAksClusterResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *ProviderAksClusterResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameResourceGroupName adds the fullNameResourceGroupName to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithFullNameResourceGroupName(fullNameResourceGroupName *string) *ProviderAksClusterResourceServiceGetParams {
	o.SetFullNameResourceGroupName(fullNameResourceGroupName)
	return o
}

// SetFullNameResourceGroupName adds the fullNameResourceGroupName to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetFullNameResourceGroupName(fullNameResourceGroupName *string) {
	o.FullNameResourceGroupName = fullNameResourceGroupName
}

// WithFullNameSubscriptionID adds the fullNameSubscriptionID to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) WithFullNameSubscriptionID(fullNameSubscriptionID *string) *ProviderAksClusterResourceServiceGetParams {
	o.SetFullNameSubscriptionID(fullNameSubscriptionID)
	return o
}

// SetFullNameSubscriptionID adds the fullNameSubscriptionId to the provider aks cluster resource service get params
func (o *ProviderAksClusterResourceServiceGetParams) SetFullNameSubscriptionID(fullNameSubscriptionID *string) {
	o.FullNameSubscriptionID = fullNameSubscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ProviderAksClusterResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
