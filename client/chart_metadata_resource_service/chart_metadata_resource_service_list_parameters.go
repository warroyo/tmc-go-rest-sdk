// Code generated by go-swagger; DO NOT EDIT.

package chart_metadata_resource_service

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

// NewChartMetadataResourceServiceListParams creates a new ChartMetadataResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChartMetadataResourceServiceListParams() *ChartMetadataResourceServiceListParams {
	return &ChartMetadataResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChartMetadataResourceServiceListParamsWithTimeout creates a new ChartMetadataResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewChartMetadataResourceServiceListParamsWithTimeout(timeout time.Duration) *ChartMetadataResourceServiceListParams {
	return &ChartMetadataResourceServiceListParams{
		timeout: timeout,
	}
}

// NewChartMetadataResourceServiceListParamsWithContext creates a new ChartMetadataResourceServiceListParams object
// with the ability to set a context for a request.
func NewChartMetadataResourceServiceListParamsWithContext(ctx context.Context) *ChartMetadataResourceServiceListParams {
	return &ChartMetadataResourceServiceListParams{
		Context: ctx,
	}
}

// NewChartMetadataResourceServiceListParamsWithHTTPClient creates a new ChartMetadataResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewChartMetadataResourceServiceListParamsWithHTTPClient(client *http.Client) *ChartMetadataResourceServiceListParams {
	return &ChartMetadataResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
ChartMetadataResourceServiceListParams contains all the parameters to send to the API endpoint

	for the chart metadata resource service list operation.

	Typically these are written to a http.Request.
*/
type ChartMetadataResourceServiceListParams struct {

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

	/* SearchScopeRepositoryName.

	   Scope search to the specified repository_name; supports globbing; default (*).
	*/
	SearchScopeRepositoryName string

	/* SortBy.

	   Sort Order.
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chart metadata resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChartMetadataResourceServiceListParams) WithDefaults() *ChartMetadataResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chart metadata resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChartMetadataResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithTimeout(timeout time.Duration) *ChartMetadataResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithContext(ctx context.Context) *ChartMetadataResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithHTTPClient(client *http.Client) *ChartMetadataResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *ChartMetadataResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *ChartMetadataResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithPaginationSize(paginationSize *string) *ChartMetadataResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithQuery(query *string) *ChartMetadataResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeName adds the searchScopeName to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *ChartMetadataResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSearchScopeRepositoryName adds the searchScopeRepositoryName to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithSearchScopeRepositoryName(searchScopeRepositoryName string) *ChartMetadataResourceServiceListParams {
	o.SetSearchScopeRepositoryName(searchScopeRepositoryName)
	return o
}

// SetSearchScopeRepositoryName adds the searchScopeRepositoryName to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetSearchScopeRepositoryName(searchScopeRepositoryName string) {
	o.SearchScopeRepositoryName = searchScopeRepositoryName
}

// WithSortBy adds the sortBy to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) WithSortBy(sortBy *string) *ChartMetadataResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the chart metadata resource service list params
func (o *ChartMetadataResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *ChartMetadataResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchScope.repositoryName
	if err := r.SetPathParam("searchScope.repositoryName", o.SearchScopeRepositoryName); err != nil {
		return err
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
