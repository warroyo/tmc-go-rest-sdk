// Code generated by go-swagger; DO NOT EDIT.

package eks_cluster_resource_service

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

// NewEksClusterResourceServiceListParams creates a new EksClusterResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEksClusterResourceServiceListParams() *EksClusterResourceServiceListParams {
	return &EksClusterResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEksClusterResourceServiceListParamsWithTimeout creates a new EksClusterResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewEksClusterResourceServiceListParamsWithTimeout(timeout time.Duration) *EksClusterResourceServiceListParams {
	return &EksClusterResourceServiceListParams{
		timeout: timeout,
	}
}

// NewEksClusterResourceServiceListParamsWithContext creates a new EksClusterResourceServiceListParams object
// with the ability to set a context for a request.
func NewEksClusterResourceServiceListParamsWithContext(ctx context.Context) *EksClusterResourceServiceListParams {
	return &EksClusterResourceServiceListParams{
		Context: ctx,
	}
}

// NewEksClusterResourceServiceListParamsWithHTTPClient creates a new EksClusterResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewEksClusterResourceServiceListParamsWithHTTPClient(client *http.Client) *EksClusterResourceServiceListParams {
	return &EksClusterResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
EksClusterResourceServiceListParams contains all the parameters to send to the API endpoint

	for the eks cluster resource service list operation.

	Typically these are written to a http.Request.
*/
type EksClusterResourceServiceListParams struct {

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

	/* SearchScopeCredentialName.

	   Scope search to the specified credential_name; supports globbing; default (*).
	*/
	SearchScopeCredentialName *string

	/* SearchScopeName.

	   Scope search to the specified name; supports globbing; default (*).
	*/
	SearchScopeName *string

	/* SearchScopeRegion.

	   Scope search to the specified region; supports globbing; default (*).
	*/
	SearchScopeRegion *string

	/* SortBy.

	   Sort Order.
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eks cluster resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterResourceServiceListParams) WithDefaults() *EksClusterResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eks cluster resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EksClusterResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithTimeout(timeout time.Duration) *EksClusterResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithContext(ctx context.Context) *EksClusterResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithHTTPClient(client *http.Client) *EksClusterResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *EksClusterResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *EksClusterResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithPaginationSize(paginationSize *string) *EksClusterResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithQuery(query *string) *EksClusterResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeCredentialName adds the searchScopeCredentialName to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithSearchScopeCredentialName(searchScopeCredentialName *string) *EksClusterResourceServiceListParams {
	o.SetSearchScopeCredentialName(searchScopeCredentialName)
	return o
}

// SetSearchScopeCredentialName adds the searchScopeCredentialName to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetSearchScopeCredentialName(searchScopeCredentialName *string) {
	o.SearchScopeCredentialName = searchScopeCredentialName
}

// WithSearchScopeName adds the searchScopeName to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *EksClusterResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSearchScopeRegion adds the searchScopeRegion to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithSearchScopeRegion(searchScopeRegion *string) *EksClusterResourceServiceListParams {
	o.SetSearchScopeRegion(searchScopeRegion)
	return o
}

// SetSearchScopeRegion adds the searchScopeRegion to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetSearchScopeRegion(searchScopeRegion *string) {
	o.SearchScopeRegion = searchScopeRegion
}

// WithSortBy adds the sortBy to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) WithSortBy(sortBy *string) *EksClusterResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the eks cluster resource service list params
func (o *EksClusterResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *EksClusterResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SearchScopeCredentialName != nil {

		// query param searchScope.credentialName
		var qrSearchScopeCredentialName string

		if o.SearchScopeCredentialName != nil {
			qrSearchScopeCredentialName = *o.SearchScopeCredentialName
		}
		qSearchScopeCredentialName := qrSearchScopeCredentialName
		if qSearchScopeCredentialName != "" {

			if err := r.SetQueryParam("searchScope.credentialName", qSearchScopeCredentialName); err != nil {
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

	if o.SearchScopeRegion != nil {

		// query param searchScope.region
		var qrSearchScopeRegion string

		if o.SearchScopeRegion != nil {
			qrSearchScopeRegion = *o.SearchScopeRegion
		}
		qSearchScopeRegion := qrSearchScopeRegion
		if qSearchScopeRegion != "" {

			if err := r.SetQueryParam("searchScope.region", qSearchScopeRegion); err != nil {
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