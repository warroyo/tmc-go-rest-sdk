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
)

// NewRecipeResourceServiceGetParams creates a new RecipeResourceServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecipeResourceServiceGetParams() *RecipeResourceServiceGetParams {
	return &RecipeResourceServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecipeResourceServiceGetParamsWithTimeout creates a new RecipeResourceServiceGetParams object
// with the ability to set a timeout on a request.
func NewRecipeResourceServiceGetParamsWithTimeout(timeout time.Duration) *RecipeResourceServiceGetParams {
	return &RecipeResourceServiceGetParams{
		timeout: timeout,
	}
}

// NewRecipeResourceServiceGetParamsWithContext creates a new RecipeResourceServiceGetParams object
// with the ability to set a context for a request.
func NewRecipeResourceServiceGetParamsWithContext(ctx context.Context) *RecipeResourceServiceGetParams {
	return &RecipeResourceServiceGetParams{
		Context: ctx,
	}
}

// NewRecipeResourceServiceGetParamsWithHTTPClient creates a new RecipeResourceServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecipeResourceServiceGetParamsWithHTTPClient(client *http.Client) *RecipeResourceServiceGetParams {
	return &RecipeResourceServiceGetParams{
		HTTPClient: client,
	}
}

/*
RecipeResourceServiceGetParams contains all the parameters to send to the API endpoint

	for the recipe resource service get operation.

	Typically these are written to a http.Request.
*/
type RecipeResourceServiceGetParams struct {

	/* FullNameName.

	   Name of policy recipe.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameTypeName.

	   Name of policy type.
	*/
	FullNameTypeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the recipe resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecipeResourceServiceGetParams) WithDefaults() *RecipeResourceServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the recipe resource service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecipeResourceServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) WithTimeout(timeout time.Duration) *RecipeResourceServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) WithContext(ctx context.Context) *RecipeResourceServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) WithHTTPClient(client *http.Client) *RecipeResourceServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameName adds the fullNameName to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) WithFullNameName(fullNameName string) *RecipeResourceServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *RecipeResourceServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameTypeName adds the fullNameTypeName to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) WithFullNameTypeName(fullNameTypeName string) *RecipeResourceServiceGetParams {
	o.SetFullNameTypeName(fullNameTypeName)
	return o
}

// SetFullNameTypeName adds the fullNameTypeName to the recipe resource service get params
func (o *RecipeResourceServiceGetParams) SetFullNameTypeName(fullNameTypeName string) {
	o.FullNameTypeName = fullNameTypeName
}

// WriteToRequest writes these params to a swagger request
func (o *RecipeResourceServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fullName.name
	if err := r.SetPathParam("fullName.name", o.FullNameName); err != nil {
		return err
	}

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

	// path param fullName.typeName
	if err := r.SetPathParam("fullName.typeName", o.FullNameTypeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
