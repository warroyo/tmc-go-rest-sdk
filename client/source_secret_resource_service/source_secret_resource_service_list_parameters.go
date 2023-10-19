// Code generated by go-swagger; DO NOT EDIT.

package source_secret_resource_service

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

// NewSourceSecretResourceServiceListParams creates a new SourceSecretResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSourceSecretResourceServiceListParams() *SourceSecretResourceServiceListParams {
	return &SourceSecretResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSourceSecretResourceServiceListParamsWithTimeout creates a new SourceSecretResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewSourceSecretResourceServiceListParamsWithTimeout(timeout time.Duration) *SourceSecretResourceServiceListParams {
	return &SourceSecretResourceServiceListParams{
		timeout: timeout,
	}
}

// NewSourceSecretResourceServiceListParamsWithContext creates a new SourceSecretResourceServiceListParams object
// with the ability to set a context for a request.
func NewSourceSecretResourceServiceListParamsWithContext(ctx context.Context) *SourceSecretResourceServiceListParams {
	return &SourceSecretResourceServiceListParams{
		Context: ctx,
	}
}

// NewSourceSecretResourceServiceListParamsWithHTTPClient creates a new SourceSecretResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSourceSecretResourceServiceListParamsWithHTTPClient(client *http.Client) *SourceSecretResourceServiceListParams {
	return &SourceSecretResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
SourceSecretResourceServiceListParams contains all the parameters to send to the API endpoint

	for the source secret resource service list operation.

	Typically these are written to a http.Request.
*/
type SourceSecretResourceServiceListParams struct {

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

	/* SearchScopeClusterName.

	   Scope search to the specified cluster_name; supports globbing; default (*).
	*/
	SearchScopeClusterName string

	/* SearchScopeManagementClusterName.

	   Scope search to the specified management_cluster_name; supports globbing; default (*).
	*/
	SearchScopeManagementClusterName *string

	/* SearchScopeName.

	   Scope search to the specified name; supports globbing; default (*).
	*/
	SearchScopeName *string

	/* SearchScopeProvisionerName.

	   Scope search to the specified provisioner_name; supports globbing; default (*).
	*/
	SearchScopeProvisionerName *string

	/* SortBy.

	   Sort Order.
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the source secret resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceSecretResourceServiceListParams) WithDefaults() *SourceSecretResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the source secret resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceSecretResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithTimeout(timeout time.Duration) *SourceSecretResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithContext(ctx context.Context) *SourceSecretResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithHTTPClient(client *http.Client) *SourceSecretResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *SourceSecretResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *SourceSecretResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithPaginationSize(paginationSize *string) *SourceSecretResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithQuery(query *string) *SourceSecretResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeClusterName adds the searchScopeClusterName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithSearchScopeClusterName(searchScopeClusterName string) *SourceSecretResourceServiceListParams {
	o.SetSearchScopeClusterName(searchScopeClusterName)
	return o
}

// SetSearchScopeClusterName adds the searchScopeClusterName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetSearchScopeClusterName(searchScopeClusterName string) {
	o.SearchScopeClusterName = searchScopeClusterName
}

// WithSearchScopeManagementClusterName adds the searchScopeManagementClusterName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithSearchScopeManagementClusterName(searchScopeManagementClusterName *string) *SourceSecretResourceServiceListParams {
	o.SetSearchScopeManagementClusterName(searchScopeManagementClusterName)
	return o
}

// SetSearchScopeManagementClusterName adds the searchScopeManagementClusterName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetSearchScopeManagementClusterName(searchScopeManagementClusterName *string) {
	o.SearchScopeManagementClusterName = searchScopeManagementClusterName
}

// WithSearchScopeName adds the searchScopeName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *SourceSecretResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSearchScopeProvisionerName adds the searchScopeProvisionerName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithSearchScopeProvisionerName(searchScopeProvisionerName *string) *SourceSecretResourceServiceListParams {
	o.SetSearchScopeProvisionerName(searchScopeProvisionerName)
	return o
}

// SetSearchScopeProvisionerName adds the searchScopeProvisionerName to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetSearchScopeProvisionerName(searchScopeProvisionerName *string) {
	o.SearchScopeProvisionerName = searchScopeProvisionerName
}

// WithSortBy adds the sortBy to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) WithSortBy(sortBy *string) *SourceSecretResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the source secret resource service list params
func (o *SourceSecretResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *SourceSecretResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchScope.clusterName
	if err := r.SetPathParam("searchScope.clusterName", o.SearchScopeClusterName); err != nil {
		return err
	}

	if o.SearchScopeManagementClusterName != nil {

		// query param searchScope.managementClusterName
		var qrSearchScopeManagementClusterName string

		if o.SearchScopeManagementClusterName != nil {
			qrSearchScopeManagementClusterName = *o.SearchScopeManagementClusterName
		}
		qSearchScopeManagementClusterName := qrSearchScopeManagementClusterName
		if qSearchScopeManagementClusterName != "" {

			if err := r.SetQueryParam("searchScope.managementClusterName", qSearchScopeManagementClusterName); err != nil {
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

	if o.SearchScopeProvisionerName != nil {

		// query param searchScope.provisionerName
		var qrSearchScopeProvisionerName string

		if o.SearchScopeProvisionerName != nil {
			qrSearchScopeProvisionerName = *o.SearchScopeProvisionerName
		}
		qSearchScopeProvisionerName := qrSearchScopeProvisionerName
		if qSearchScopeProvisionerName != "" {

			if err := r.SetQueryParam("searchScope.provisionerName", qSearchScopeProvisionerName); err != nil {
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
