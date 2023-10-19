// Code generated by go-swagger; DO NOT EDIT.

package subscriptions_service

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

// NewSubscriptionsServiceListParams creates a new SubscriptionsServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionsServiceListParams() *SubscriptionsServiceListParams {
	return &SubscriptionsServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionsServiceListParamsWithTimeout creates a new SubscriptionsServiceListParams object
// with the ability to set a timeout on a request.
func NewSubscriptionsServiceListParamsWithTimeout(timeout time.Duration) *SubscriptionsServiceListParams {
	return &SubscriptionsServiceListParams{
		timeout: timeout,
	}
}

// NewSubscriptionsServiceListParamsWithContext creates a new SubscriptionsServiceListParams object
// with the ability to set a context for a request.
func NewSubscriptionsServiceListParamsWithContext(ctx context.Context) *SubscriptionsServiceListParams {
	return &SubscriptionsServiceListParams{
		Context: ctx,
	}
}

// NewSubscriptionsServiceListParamsWithHTTPClient creates a new SubscriptionsServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionsServiceListParamsWithHTTPClient(client *http.Client) *SubscriptionsServiceListParams {
	return &SubscriptionsServiceListParams{
		HTTPClient: client,
	}
}

/*
SubscriptionsServiceListParams contains all the parameters to send to the API endpoint

	for the subscriptions service list operation.

	Typically these are written to a http.Request.
*/
type SubscriptionsServiceListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscriptions service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionsServiceListParams) WithDefaults() *SubscriptionsServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscriptions service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionsServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscriptions service list params
func (o *SubscriptionsServiceListParams) WithTimeout(timeout time.Duration) *SubscriptionsServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscriptions service list params
func (o *SubscriptionsServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscriptions service list params
func (o *SubscriptionsServiceListParams) WithContext(ctx context.Context) *SubscriptionsServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscriptions service list params
func (o *SubscriptionsServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscriptions service list params
func (o *SubscriptionsServiceListParams) WithHTTPClient(client *http.Client) *SubscriptionsServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscriptions service list params
func (o *SubscriptionsServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionsServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
