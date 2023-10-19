// Code generated by go-swagger; DO NOT EDIT.

package credential_resource_service

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

// NewCredentialResourceServiceListParams creates a new CredentialResourceServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialResourceServiceListParams() *CredentialResourceServiceListParams {
	return &CredentialResourceServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialResourceServiceListParamsWithTimeout creates a new CredentialResourceServiceListParams object
// with the ability to set a timeout on a request.
func NewCredentialResourceServiceListParamsWithTimeout(timeout time.Duration) *CredentialResourceServiceListParams {
	return &CredentialResourceServiceListParams{
		timeout: timeout,
	}
}

// NewCredentialResourceServiceListParamsWithContext creates a new CredentialResourceServiceListParams object
// with the ability to set a context for a request.
func NewCredentialResourceServiceListParamsWithContext(ctx context.Context) *CredentialResourceServiceListParams {
	return &CredentialResourceServiceListParams{
		Context: ctx,
	}
}

// NewCredentialResourceServiceListParamsWithHTTPClient creates a new CredentialResourceServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialResourceServiceListParamsWithHTTPClient(client *http.Client) *CredentialResourceServiceListParams {
	return &CredentialResourceServiceListParams{
		HTTPClient: client,
	}
}

/*
CredentialResourceServiceListParams contains all the parameters to send to the API endpoint

	for the credential resource service list operation.

	Typically these are written to a http.Request.
*/
type CredentialResourceServiceListParams struct {

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

	/* SearchScopeManagementClusterName.

	   Scope search to the specified management_cluster_name; supports globbing; default (*).
	*/
	SearchScopeManagementClusterName *string

	/* SearchScopeName.

	   Scope search to the specified name; supports globbing; default (*).
	*/
	SearchScopeName *string

	/* SearchScopeProvider.

	   Filter credentials by a provider; supports globbing; default (*).
	*/
	SearchScopeProvider *string

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

// WithDefaults hydrates default values in the credential resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialResourceServiceListParams) WithDefaults() *CredentialResourceServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credential resource service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialResourceServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithTimeout(timeout time.Duration) *CredentialResourceServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithContext(ctx context.Context) *CredentialResourceServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithHTTPClient(client *http.Client) *CredentialResourceServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTotalCount adds the includeTotalCount to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithIncludeTotalCount(includeTotalCount *bool) *CredentialResourceServiceListParams {
	o.SetIncludeTotalCount(includeTotalCount)
	return o
}

// SetIncludeTotalCount adds the includeTotalCount to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetIncludeTotalCount(includeTotalCount *bool) {
	o.IncludeTotalCount = includeTotalCount
}

// WithPaginationOffset adds the paginationOffset to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithPaginationOffset(paginationOffset *string) *CredentialResourceServiceListParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationSize adds the paginationSize to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithPaginationSize(paginationSize *string) *CredentialResourceServiceListParams {
	o.SetPaginationSize(paginationSize)
	return o
}

// SetPaginationSize adds the paginationSize to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetPaginationSize(paginationSize *string) {
	o.PaginationSize = paginationSize
}

// WithQuery adds the query to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithQuery(query *string) *CredentialResourceServiceListParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetQuery(query *string) {
	o.Query = query
}

// WithSearchScopeManagementClusterName adds the searchScopeManagementClusterName to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithSearchScopeManagementClusterName(searchScopeManagementClusterName *string) *CredentialResourceServiceListParams {
	o.SetSearchScopeManagementClusterName(searchScopeManagementClusterName)
	return o
}

// SetSearchScopeManagementClusterName adds the searchScopeManagementClusterName to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetSearchScopeManagementClusterName(searchScopeManagementClusterName *string) {
	o.SearchScopeManagementClusterName = searchScopeManagementClusterName
}

// WithSearchScopeName adds the searchScopeName to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithSearchScopeName(searchScopeName *string) *CredentialResourceServiceListParams {
	o.SetSearchScopeName(searchScopeName)
	return o
}

// SetSearchScopeName adds the searchScopeName to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetSearchScopeName(searchScopeName *string) {
	o.SearchScopeName = searchScopeName
}

// WithSearchScopeProvider adds the searchScopeProvider to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithSearchScopeProvider(searchScopeProvider *string) *CredentialResourceServiceListParams {
	o.SetSearchScopeProvider(searchScopeProvider)
	return o
}

// SetSearchScopeProvider adds the searchScopeProvider to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetSearchScopeProvider(searchScopeProvider *string) {
	o.SearchScopeProvider = searchScopeProvider
}

// WithSearchScopeProvisionerName adds the searchScopeProvisionerName to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithSearchScopeProvisionerName(searchScopeProvisionerName *string) *CredentialResourceServiceListParams {
	o.SetSearchScopeProvisionerName(searchScopeProvisionerName)
	return o
}

// SetSearchScopeProvisionerName adds the searchScopeProvisionerName to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetSearchScopeProvisionerName(searchScopeProvisionerName *string) {
	o.SearchScopeProvisionerName = searchScopeProvisionerName
}

// WithSortBy adds the sortBy to the credential resource service list params
func (o *CredentialResourceServiceListParams) WithSortBy(sortBy *string) *CredentialResourceServiceListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the credential resource service list params
func (o *CredentialResourceServiceListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialResourceServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SearchScopeProvider != nil {

		// query param searchScope.provider
		var qrSearchScopeProvider string

		if o.SearchScopeProvider != nil {
			qrSearchScopeProvider = *o.SearchScopeProvider
		}
		qSearchScopeProvider := qrSearchScopeProvider
		if qSearchScopeProvider != "" {

			if err := r.SetQueryParam("searchScope.provider", qSearchScopeProvider); err != nil {
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