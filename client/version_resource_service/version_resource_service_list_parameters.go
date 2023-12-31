// Code generated by go-swagger; DO NOT EDIT.

package version_resource_service

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

// NewVersionResourceServiceListParams creates a new VersionResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersionResourceServiceListParams() *VersionResourceServiceListParams {
	return &VersionResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersionResourceServiceListParamsWithTimeout creates a new VersionResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewVersionResourceServiceListParamsWithTimeout(timeout time.Duration) *VersionResourceServiceListParams {
	return &VersionResourceServiceListParams{
		timeout: timeout,
	}
}

// NewVersionResourceServiceListParamsWithContext creates a new VersionResourceServiceListParams object
// with the ability to set a context for a request.
func NewVersionResourceServiceListParamsWithContext(ctx context.Context) *VersionResourceServiceListParams {
	return &VersionResourceServiceListParams{
		Context: ctx,
	}
}

// NewVersionResourceServiceListParamsWithHTTPClient creates a new VersionResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersionResourceServiceListParamsWithHTTPClient(client *http.Client) *VersionResourceServiceListParams {
	return &VersionResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
VersionResourceServiceListParams contains all the parameters to send to the API endpoint

	for the version resource service list operation.

	Typically these are written to a http.Request.
*/
type VersionResourceServiceListParams struct {

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

	/* SearchScopeRecipeName.

	   Scope search to the specified recipe_name; supports globbing; default (*).
	*/
	SearchScopeRecipeName string

	/* SearchScopeTypeName.

	   Scope search to the specified type_name; supports globbing; default (*).
	*/
	SearchScopeTypeName string

	/* SortBy.

	   Sort Order.
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the version resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionResourceServiceListParams) WithDefaults() *VersionResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the version resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the version resource service list params
func (o *VersionResourceServiceListParams) WithTimeout(timeout time.Duration) *VersionResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version resource service list params
func (o *VersionResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version resource service list params
func (o *VersionResourceServiceListParams) WithContext(ctx context.Context) *VersionResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version resource service list params
func (o *VersionResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version resource service list params
func (o *VersionResourceServiceListParams) WithHTTPClient(client *http.Client) *VersionResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version resource service list params
func (o *VersionResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the version resource service list params
func (o *VersionResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *VersionResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the version resource service list params
func (o *VersionResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the version resource service list params
func (o *VersionResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *VersionResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the version resource service list params
func (o *VersionResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the version resource service list params
func (o *VersionResourceServiceListParams) WithPaginationSize(paginationSize *string) *VersionResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the version resource service list params
func (o *VersionResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the version resource service list params
func (o *VersionResourceServiceListParams) WithQuery(query *string) *VersionResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the version resource service list params
func (o *VersionResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeName adds the searchScopeName to the version resource service list params
func (o *VersionResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *VersionResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the version resource service list params
func (o *VersionResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSearchScopeRecipeName adds the searchScopeRecipeName to the version resource service list params
func (o *VersionResourceServiceListParams) WithSearchScopeRecipeName(searchScopeRecipeName string) *VersionResourceServiceListParams {
	o.SetSearchScopeRecipeName(searchScopeRecipeName)
	return o
}

// SetSearchScopeRecipeName adds the searchScopeRecipeName to the version resource service list params
func (o *VersionResourceServiceListParams) SetSearchScopeRecipeName(searchScopeRecipeName string) {
	o.SearchScopeRecipeName = searchScopeRecipeName
}

// WithSearchScopeTypeName adds the searchScopeTypeName to the version resource service list params
func (o *VersionResourceServiceListParams) WithSearchScopeTypeName(searchScopeTypeName string) *VersionResourceServiceListParams {
	o.SetSearchScopeTypeName(searchScopeTypeName)
	return o
}

// SetSearchScopeTypeName adds the searchScopeTypeName to the version resource service list params
func (o *VersionResourceServiceListParams) SetSearchScopeTypeName(searchScopeTypeName string) {
	o.SearchScopeTypeName = searchScopeTypeName
}

// WithSortBy adds the sortBy to the version resource service list params
func (o *VersionResourceServiceListParams) WithSortBy(sortBy *string) *VersionResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the version resource service list params
func (o *VersionResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *VersionResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchScope.recipeName
	if err := r.SetPathParam("searchScope.recipeName", o.SearchScopeRecipeName); err != nil {
		return err
	}

	// path param searchScope.typeName
	if err := r.SetPathParam("searchScope.typeName", o.SearchScopeTypeName); err != nil {
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
