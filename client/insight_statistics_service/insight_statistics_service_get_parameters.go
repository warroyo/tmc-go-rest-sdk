// Code generated by go-swagger; DO NOT EDIT.

package insight_statistics_service

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

// NewInsightStatisticsServiceGetParams creates a new InsightStatisticsServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsightStatisticsServiceGetParams() *InsightStatisticsServiceGetParams {
	return &InsightStatisticsServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsightStatisticsServiceGetParamsWithTimeout creates a new InsightStatisticsServiceGetParams object
// with the ability to set a timeout on a request.
func NewInsightStatisticsServiceGetParamsWithTimeout(timeout time.Duration) *InsightStatisticsServiceGetParams {
	return &InsightStatisticsServiceGetParams{
		timeout: timeout,
	}
}

// NewInsightStatisticsServiceGetParamsWithContext creates a new InsightStatisticsServiceGetParams object
// with the ability to set a context for a request.
func NewInsightStatisticsServiceGetParamsWithContext(ctx context.Context) *InsightStatisticsServiceGetParams {
	return &InsightStatisticsServiceGetParams{
		Context: ctx,
	}
}

// NewInsightStatisticsServiceGetParamsWithHTTPClient creates a new InsightStatisticsServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsightStatisticsServiceGetParamsWithHTTPClient(client *http.Client) *InsightStatisticsServiceGetParams {
	return &InsightStatisticsServiceGetParams{
		HTTPClient: client,
	}
}

/*
InsightStatisticsServiceGetParams contains all the parameters to send to the API endpoint

	for the insight statistics service get operation.

	Typically these are written to a http.Request.
*/
type InsightStatisticsServiceGetParams struct {

	/* OrgID.

	   ID of Organization.
	*/
	OrgID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the insight statistics service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsightStatisticsServiceGetParams) WithDefaults() *InsightStatisticsServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insight statistics service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsightStatisticsServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) WithTimeout(timeout time.Duration) *InsightStatisticsServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) WithContext(ctx context.Context) *InsightStatisticsServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) WithHTTPClient(client *http.Client) *InsightStatisticsServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgID adds the orgID to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) WithOrgID(orgID *string) *InsightStatisticsServiceGetParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the insight statistics service get params
func (o *InsightStatisticsServiceGetParams) SetOrgID(orgID *string) {
	o.OrgID = orgID
}

// WriteToRequest writes these params to a swagger request
func (o *InsightStatisticsServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
