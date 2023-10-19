// Code generated by go-swagger; DO NOT EDIT.

package organization_i_a_m_policy

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

// NewOrganizationIAMPolicyGetParams creates a new OrganizationIAMPolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationIAMPolicyGetParams() *OrganizationIAMPolicyGetParams {
	return &OrganizationIAMPolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationIAMPolicyGetParamsWithTimeout creates a new OrganizationIAMPolicyGetParams object
// with the ability to set a timeout on a request.
func NewOrganizationIAMPolicyGetParamsWithTimeout(timeout time.Duration) *OrganizationIAMPolicyGetParams {
	return &OrganizationIAMPolicyGetParams{
		timeout: timeout,
	}
}

// NewOrganizationIAMPolicyGetParamsWithContext creates a new OrganizationIAMPolicyGetParams object
// with the ability to set a context for a request.
func NewOrganizationIAMPolicyGetParamsWithContext(ctx context.Context) *OrganizationIAMPolicyGetParams {
	return &OrganizationIAMPolicyGetParams{
		Context: ctx,
	}
}

// NewOrganizationIAMPolicyGetParamsWithHTTPClient creates a new OrganizationIAMPolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationIAMPolicyGetParamsWithHTTPClient(client *http.Client) *OrganizationIAMPolicyGetParams {
	return &OrganizationIAMPolicyGetParams{
		HTTPClient: client,
	}
}

/*
OrganizationIAMPolicyGetParams contains all the parameters to send to the API endpoint

	for the organization i a m policy get operation.

	Typically these are written to a http.Request.
*/
type OrganizationIAMPolicyGetParams struct {

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationIAMPolicyGetParams) WithDefaults() *OrganizationIAMPolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization i a m policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationIAMPolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) WithTimeout(timeout time.Duration) *OrganizationIAMPolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) WithContext(ctx context.Context) *OrganizationIAMPolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) WithHTTPClient(client *http.Client) *OrganizationIAMPolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameOrgID adds the fullNameOrgID to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) WithFullNameOrgID(fullNameOrgID *string) *OrganizationIAMPolicyGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the organization i a m policy get params
func (o *OrganizationIAMPolicyGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationIAMPolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
