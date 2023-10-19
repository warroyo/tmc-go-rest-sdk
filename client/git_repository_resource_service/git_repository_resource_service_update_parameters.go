// Code generated by go-swagger; DO NOT EDIT.

package git_repository_resource_service

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

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// NewGitRepositoryResourceServiceUpdateParams creates a new GitRepositoryResourceServiceUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGitRepositoryResourceServiceUpdateParams() *GitRepositoryResourceServiceUpdateParams {
	return &GitRepositoryResourceServiceUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGitRepositoryResourceServiceUpdateParamsWithTimeout creates a new GitRepositoryResourceServiceUpdateParams object
// with the ability to set a timeout on a request.
func NewGitRepositoryResourceServiceUpdateParamsWithTimeout(timeout time.Duration) *GitRepositoryResourceServiceUpdateParams {
	return &GitRepositoryResourceServiceUpdateParams{
		timeout: timeout,
	}
}

// NewGitRepositoryResourceServiceUpdateParamsWithContext creates a new GitRepositoryResourceServiceUpdateParams object
// with the ability to set a context for a request.
func NewGitRepositoryResourceServiceUpdateParamsWithContext(ctx context.Context) *GitRepositoryResourceServiceUpdateParams {
	return &GitRepositoryResourceServiceUpdateParams{
		Context: ctx,
	}
}

// NewGitRepositoryResourceServiceUpdateParamsWithHTTPClient creates a new GitRepositoryResourceServiceUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGitRepositoryResourceServiceUpdateParamsWithHTTPClient(client *http.Client) *GitRepositoryResourceServiceUpdateParams {
	return &GitRepositoryResourceServiceUpdateParams{
		HTTPClient: client,
	}
}

/*
GitRepositoryResourceServiceUpdateParams contains all the parameters to send to the API endpoint

	for the git repository resource service update operation.

	Typically these are written to a http.Request.
*/
type GitRepositoryResourceServiceUpdateParams struct {

	// Body.
	Body *models.ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest

	/* GitRepositoryFullNameClusterName.

	   Name of Cluster.
	*/
	GitRepositoryFullNameClusterName string

	/* GitRepositoryFullNameName.

	   Name of the Repository.
	*/
	GitRepositoryFullNameName string

	/* GitRepositoryFullNameNamespaceName.

	   Name of Namespace.
	*/
	GitRepositoryFullNameNamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the git repository resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GitRepositoryResourceServiceUpdateParams) WithDefaults() *GitRepositoryResourceServiceUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the git repository resource service update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GitRepositoryResourceServiceUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithTimeout(timeout time.Duration) *GitRepositoryResourceServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithContext(ctx context.Context) *GitRepositoryResourceServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithHTTPClient(client *http.Client) *GitRepositoryResourceServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithBody(body *models.ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) *GitRepositoryResourceServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetBody(body *models.ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryRequest) {
	o.Body = body
}

// WithGitRepositoryFullNameClusterName adds the gitRepositoryFullNameClusterName to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithGitRepositoryFullNameClusterName(gitRepositoryFullNameClusterName string) *GitRepositoryResourceServiceUpdateParams {
	o.SetGitRepositoryFullNameClusterName(gitRepositoryFullNameClusterName)
	return o
}

// SetGitRepositoryFullNameClusterName adds the gitRepositoryFullNameClusterName to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetGitRepositoryFullNameClusterName(gitRepositoryFullNameClusterName string) {
	o.GitRepositoryFullNameClusterName = gitRepositoryFullNameClusterName
}

// WithGitRepositoryFullNameName adds the gitRepositoryFullNameName to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithGitRepositoryFullNameName(gitRepositoryFullNameName string) *GitRepositoryResourceServiceUpdateParams {
	o.SetGitRepositoryFullNameName(gitRepositoryFullNameName)
	return o
}

// SetGitRepositoryFullNameName adds the gitRepositoryFullNameName to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetGitRepositoryFullNameName(gitRepositoryFullNameName string) {
	o.GitRepositoryFullNameName = gitRepositoryFullNameName
}

// WithGitRepositoryFullNameNamespaceName adds the gitRepositoryFullNameNamespaceName to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) WithGitRepositoryFullNameNamespaceName(gitRepositoryFullNameNamespaceName string) *GitRepositoryResourceServiceUpdateParams {
	o.SetGitRepositoryFullNameNamespaceName(gitRepositoryFullNameNamespaceName)
	return o
}

// SetGitRepositoryFullNameNamespaceName adds the gitRepositoryFullNameNamespaceName to the git repository resource service update params
func (o *GitRepositoryResourceServiceUpdateParams) SetGitRepositoryFullNameNamespaceName(gitRepositoryFullNameNamespaceName string) {
	o.GitRepositoryFullNameNamespaceName = gitRepositoryFullNameNamespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *GitRepositoryResourceServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param gitRepository.fullName.clusterName
	if err := r.SetPathParam("gitRepository.fullName.clusterName", o.GitRepositoryFullNameClusterName); err != nil {
		return err
	}

	// path param gitRepository.fullName.name
	if err := r.SetPathParam("gitRepository.fullName.name", o.GitRepositoryFullNameName); err != nil {
		return err
	}

	// path param gitRepository.fullName.namespaceName
	if err := r.SetPathParam("gitRepository.fullName.namespaceName", o.GitRepositoryFullNameNamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}