// Code generated by go-swagger; DO NOT EDIT.

package provider_eks_cluster_i_a_m_policy

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

// NewProviderEksClusterIAMPolicyGetParams creates a new ProviderEksClusterIAMPolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProviderEksClusterIAMPolicyGetParams() *ProviderEksClusterIAMPolicyGetParams {
	return &ProviderEksClusterIAMPolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProviderEksClusterIAMPolicyGetParamsWithTimeout creates a new ProviderEksClusterIAMPolicyGetParams object
// with the ability to set a timeout on a request.
func NewProviderEksClusterIAMPolicyGetParamsWithTimeout(timeout time.Duration) *ProviderEksClusterIAMPolicyGetParams {
	return &ProviderEksClusterIAMPolicyGetParams{
		timeout: timeout,
	}
}

// NewProviderEksClusterIAMPolicyGetParamsWithContext creates a new ProviderEksClusterIAMPolicyGetParams object
// with the ability to set a context for a request.
func NewProviderEksClusterIAMPolicyGetParamsWithContext(ctx context.Context) *ProviderEksClusterIAMPolicyGetParams {
	return &ProviderEksClusterIAMPolicyGetParams{
		Context: ctx,
	}
}

// NewProviderEksClusterIAMPolicyGetParamsWithHTTPClient creates a new ProviderEksClusterIAMPolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewProviderEksClusterIAMPolicyGetParamsWithHTTPClient(client *http.Client) *ProviderEksClusterIAMPolicyGetParams {
	return &ProviderEksClusterIAMPolicyGetParams{
		HTTPClient: client,
	}
}

/*
ProviderEksClusterIAMPolicyGetParams contains all the parameters to send to the API endpoint

	for the provider eks cluster i a m policy get operation.

	Typically these are written to a http.Request.
*/
type ProviderEksClusterIAMPolicyGetParams struct {

	/* FullNameCredentialName.

	   Name of the credential.
	*/
	FullNameCredentialName *string

	/* FullNameName.

	   Name of this cluster in EKS.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameRegion.

	   Name of the region.
	*/
	FullNameRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provider eks cluster i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderEksClusterIAMPolicyGetParams) WithDefaults() *ProviderEksClusterIAMPolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provider eks cluster i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderEksClusterIAMPolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithTimeout(timeout time.Duration) *ProviderEksClusterIAMPolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithContext(ctx context.Context) *ProviderEksClusterIAMPolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithHTTPClient(client *http.Client) *ProviderEksClusterIAMPolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameCredentialName adds the fullNameCredentialName to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithFullNameCredentialName(fullNameCredentialName *string) *ProviderEksClusterIAMPolicyGetParams {
	o.SetFullNameCredentialName(fullNameCredentialName)
	return o
}

// SetFullNameCredentialName adds the fullNameCredentialName to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetFullNameCredentialName(fullNameCredentialName *string) {
	o.FullNameCredentialName = fullNameCredentialName
}

// WithFullNameName adds the fullNameName to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithFullNameName(fullNameName string) *ProviderEksClusterIAMPolicyGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithFullNameOrgID(fullNameOrgID *string) *ProviderEksClusterIAMPolicyGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameRegion adds the fullNameRegion to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) WithFullNameRegion(fullNameRegion *string) *ProviderEksClusterIAMPolicyGetParams {
	o.SetFullNameRegion(fullNameRegion)
	return o
}

// SetFullNameRegion adds the fullNameRegion to the provider eks cluster i a m policy get params
func (o *ProviderEksClusterIAMPolicyGetParams) SetFullNameRegion(fullNameRegion *string) {
	o.FullNameRegion = fullNameRegion
}

// WriteToRequest writes these params to a swagger request
func (o *ProviderEksClusterIAMPolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FullNameRegion != nil {

		// query param fullName.region
		var qrFullNameRegion string

		if o.FullNameRegion != nil {
			qrFullNameRegion = *o.FullNameRegion
		}
		qFullNameRegion := qrFullNameRegion
		if qFullNameRegion != "" {

			if err := r.SetQueryParam("fullName.region", qFullNameRegion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
