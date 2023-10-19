// Code generated by go-swagger; DO NOT EDIT.

package provider_resource_service

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
	"github.com/go-openapi/swag"
)

// NewProviderResourceServiceListParams creates a new ProviderResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProviderResourceServiceListParams() *ProviderResourceServiceListParams {
	return &ProviderResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProviderResourceServiceListParamsWithTimeout creates a new ProviderResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewProviderResourceServiceListParamsWithTimeout(timeout time.Duration) *ProviderResourceServiceListParams {
	return &ProviderResourceServiceListParams{
		timeout: timeout,
	}
}

// NewProviderResourceServiceListParamsWithContext creates a new ProviderResourceServiceListParams object
// with the ability to set a context for a request.
func NewProviderResourceServiceListParamsWithContext(ctx context.Context) *ProviderResourceServiceListParams {
	return &ProviderResourceServiceListParams{
		Context: ctx,
	}
}

// NewProviderResourceServiceListParamsWithHTTPClient creates a new ProviderResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewProviderResourceServiceListParamsWithHTTPClient(client *http.Client) *ProviderResourceServiceListParams {
	return &ProviderResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
ProviderResourceServiceListParams contains all the parameters to send to the API endpoint

	for the provider resource service list operation.

	Typically these are written to a http.Request.
*/
type ProviderResourceServiceListParams struct {

	/* IncludeTotalCount.

	   Include total count.
	*/
	IncludeTotalCount *bool

	/* PaginationOffset.

	   Offset at which to start returning records.

	   Format: uint64
	*/
	PaginationOffset *string

	/* PaginationSize.

	   Number of records to return.

	   Format: uint64
	*/
	PaginationSize *string

	/* Query.

	   TQL query string.
	*/
	Query *string

	/* SearchScopeName.

	   Scope search to the specified name; supports globbing; default (*).
	*/
	SearchScopeName *string

	/* SortBy.

	   Sort Order.
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provider resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderResourceServiceListParams) WithDefaults() *ProviderResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provider resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithTimeout(timeout time.Duration) *ProviderResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithContext(ctx context.Context) *ProviderResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithHTTPClient(client *http.Client) *ProviderResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *ProviderResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *ProviderResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithPaginationSize(paginationSize *string) *ProviderResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithQuery(query *string) *ProviderResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeName adds the searchScopeName to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *ProviderResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSortBy adds the sortBy to the provider resource service list params
func (o *ProviderResourceServiceListParams) WithSortBy(sortBy *string) *ProviderResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the provider resource service list params
func (o *ProviderResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *ProviderResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeTotalCount != nil {

		// query param includeTotalCount
		var qrIncludeTotalCount bool

		if o.IncludeTotalCount != nil {
			qrIncludeTotalCount = *o.IncludeTotalCount
		}
		qIncludeTotalCount := swag.FormatBool(qrIncludeTotalCount)
		if qIncludeTotalCount != "" {

			if err := r.SetQueryParam("includeTotalCount", qIncludeTotalCount); err != nil {
				return err
			}
		}
	}

	if o.PaginationOffset != nil {

		// query param pagination.offset
		var qrPaginationOffset string

		if o.PaginationOffset != nil {
			qrPaginationOffset = *o.PaginationOffset
		}
		qPaginationOffset := qrPaginationOffset
		if qPaginationOffset != "" {

			if err := r.SetQueryParam("pagination.offset", qPaginationOffset); err != nil {
				return err
			}
		}
	}

	if o.PaginationSize != nil {

		// query param pagination.size
		var qrPaginationSize string

		if o.PaginationSize != nil {
			qrPaginationSize = *o.PaginationSize
		}
		qPaginationSize := qrPaginationSize
		if qPaginationSize != "" {

			if err := r.SetQueryParam("pagination.size", qPaginationSize); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.SearchScopeName != nil {

		// query param searchScope.name
		var qrSearchScopeName string

		if o.SearchScopeName != nil {
			qrSearchScopeName = *o.SearchScopeName
		}
		qSearchScopeName := qrSearchScopeName
		if qSearchScopeName != "" {

			if err := r.SetQueryParam("searchScope.name", qSearchScopeName); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
