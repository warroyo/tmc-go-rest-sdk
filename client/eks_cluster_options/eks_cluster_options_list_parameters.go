// Code generated by go-swagger; DO NOT EDIT.

package eks_cluster_options

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

// NewEksClusterOptionsListParams creates a new EksClusterOptionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEksClusterOptionsListParams() *EksClusterOptionsListParams {
	return &EksClusterOptionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEksClusterOptionsListParamsWithTimeout creates a new EksClusterOptionsListParams object
// with the ability to set a timeout on a request.
func NewEksClusterOptionsListParamsWithTimeout(timeout time.Duration) *EksClusterOptionsListParams {
	return &EksClusterOptionsListParams{
		timeout: timeout,
	}
}

// NewEksClusterOptionsListParamsWithContext creates a new EksClusterOptionsListParams object
// with the ability to set a context for a request.
func NewEksClusterOptionsListParamsWithContext(ctx context.Context) *EksClusterOptionsListParams {
	return &EksClusterOptionsListParams{
		Context: ctx,
	}
}

// NewEksClusterOptionsListParamsWithHTTPClient creates a new EksClusterOptionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewEksClusterOptionsListParamsWithHTTPClient(client *http.Client) *EksClusterOptionsListParams {
	return &EksClusterOptionsListParams{
		HTTPClient: client,
	}
}

/*
EksClusterOptionsListParams contains all the parameters to send to the API endpoint

	for the eks cluster options list operation.

	Typically these are written to a http.Request.
*/
type EksClusterOptionsListParams struct {

	/* CredentialName.

	   Name of this credential.
	*/
	CredentialName *string

	/* OrgID.

	   ID of Organization.
	*/
	OrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eks cluster options list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterOptionsListParams) WithDefaults() *EksClusterOptionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eks cluster options list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterOptionsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eks cluster options list params
func (o *EksClusterOptionsListParams) WithTimeout(timeout time.Duration) *EksClusterOptionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eks cluster options list params
func (o *EksClusterOptionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eks cluster options list params
func (o *EksClusterOptionsListParams) WithContext(ctx context.Context) *EksClusterOptionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eks cluster options list params
func (o *EksClusterOptionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eks cluster options list params
func (o *EksClusterOptionsListParams) WithHTTPClient(client *http.Client) *EksClusterOptionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eks cluster options list params
func (o *EksClusterOptionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialName adds the credentialName to the eks cluster options list params
func (o *EksClusterOptionsListParams) WithCredentialName(credentialName *string) *EksClusterOptionsListParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the eks cluster options list params
func (o *EksClusterOptionsListParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithOrgID adds the orgID to the eks cluster options list params
func (o *EksClusterOptionsListParams) WithOrgID(orgID *string) *EksClusterOptionsListParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the eks cluster options list params
func (o *EksClusterOptionsListParams) SetOrgID(orgID *string) {
	o.OrgID = orgID
}

// WriteToRequest writes these params to a swagger request
func (o *EksClusterOptionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}