// Code generated by go-swagger; DO NOT EDIT.

package recipe_resource_service

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

// NewRecipeResourceServiceListParams creates a new RecipeResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecipeResourceServiceListParams() *RecipeResourceServiceListParams {
	return &RecipeResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecipeResourceServiceListParamsWithTimeout creates a new RecipeResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewRecipeResourceServiceListParamsWithTimeout(timeout time.Duration) *RecipeResourceServiceListParams {
	return &RecipeResourceServiceListParams{
		timeout: timeout,
	}
}

// NewRecipeResourceServiceListParamsWithContext creates a new RecipeResourceServiceListParams object
// with the ability to set a context for a request.
func NewRecipeResourceServiceListParamsWithContext(ctx context.Context) *RecipeResourceServiceListParams {
	return &RecipeResourceServiceListParams{
		Context: ctx,
	}
}

// NewRecipeResourceServiceListParamsWithHTTPClient creates a new RecipeResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecipeResourceServiceListParamsWithHTTPClient(client *http.Client) *RecipeResourceServiceListParams {
	return &RecipeResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
RecipeResourceServiceListParams contains all the parameters to send to the API endpoint

	for the recipe resource service list operation.

	Typically these are written to a http.Request.
*/
type RecipeResourceServiceListParams struct {

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

// WithDefaults hydrates default values in the recipe resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecipeResourceServiceListParams) WithDefaults() *RecipeResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the recipe resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecipeResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithTimeout(timeout time.Duration) *RecipeResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithContext(ctx context.Context) *RecipeResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithHTTPClient(client *http.Client) *RecipeResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *RecipeResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *RecipeResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithPaginationSize(paginationSize *string) *RecipeResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithQuery(query *string) *RecipeResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeName adds the searchScopeName to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *RecipeResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSearchScopeTypeName adds the searchScopeTypeName to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithSearchScopeTypeName(searchScopeTypeName string) *RecipeResourceServiceListParams {
	o.SetSearchScopeTypeName(searchScopeTypeName)
	return o
}

// SetSearchScopeTypeName adds the searchScopeTypeName to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetSearchScopeTypeName(searchScopeTypeName string) {
	o.SearchScopeTypeName = searchScopeTypeName
}

// WithSortBy adds the sortBy to the recipe resource service list params
func (o *RecipeResourceServiceListParams) WithSortBy(sortBy *string) *RecipeResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the recipe resource service list params
func (o *RecipeResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *RecipeResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
