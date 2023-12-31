// Code generated by go-swagger; DO NOT EDIT.

package nodepool_apply_helper

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

// NewNodepoolApplyHelperApplyParams creates a new NodepoolApplyHelperApplyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodepoolApplyHelperApplyParams() *NodepoolApplyHelperApplyParams {
	return &NodepoolApplyHelperApplyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodepoolApplyHelperApplyParamsWithTimeout creates a new NodepoolApplyHelperApplyParams object
// with the ability to set a timeout on a request.
func NewNodepoolApplyHelperApplyParamsWithTimeout(timeout time.Duration) *NodepoolApplyHelperApplyParams {
	return &NodepoolApplyHelperApplyParams{
		timeout: timeout,
	}
}

// NewNodepoolApplyHelperApplyParamsWithContext creates a new NodepoolApplyHelperApplyParams object
// with the ability to set a context for a request.
func NewNodepoolApplyHelperApplyParamsWithContext(ctx context.Context) *NodepoolApplyHelperApplyParams {
	return &NodepoolApplyHelperApplyParams{
		Context: ctx,
	}
}

// NewNodepoolApplyHelperApplyParamsWithHTTPClient creates a new NodepoolApplyHelperApplyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodepoolApplyHelperApplyParamsWithHTTPClient(client *http.Client) *NodepoolApplyHelperApplyParams {
	return &NodepoolApplyHelperApplyParams{
		HTTPClient: client,
	}
}

/*
NodepoolApplyHelperApplyParams contains all the parameters to send to the API endpoint

	for the nodepool apply helper apply operation.

	Typically these are written to a http.Request.
*/
type NodepoolApplyHelperApplyParams struct {

	// Body.
	Body *models.ManagementclusterProvisionerTanzukubernetesclusterNodepoolApplyNodepoolRequest

	/* NodepoolFullNameManagementClusterName.

	   Name of the management cluster.
	*/
	NodepoolFullNameManagementClusterName string

	/* NodepoolFullNameProvisionerName.

	   Provisioner of the cluster.
	*/
	NodepoolFullNameProvisionerName string

	/* NodepoolFullNameTanzuKubernetesClusterName.

	   Name of the cluster.
	*/
	NodepoolFullNameTanzuKubernetesClusterName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nodepool apply helper apply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolApplyHelperApplyParams) WithDefaults() *NodepoolApplyHelperApplyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodepool apply helper apply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodepoolApplyHelperApplyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithTimeout(timeout time.Duration) *NodepoolApplyHelperApplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithContext(ctx context.Context) *NodepoolApplyHelperApplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithHTTPClient(client *http.Client) *NodepoolApplyHelperApplyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithBody(body *models.ManagementclusterProvisionerTanzukubernetesclusterNodepoolApplyNodepoolRequest) *NodepoolApplyHelperApplyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetBody(body *models.ManagementclusterProvisionerTanzukubernetesclusterNodepoolApplyNodepoolRequest) {
	o.Body = body
}

// WithNodepoolFullNameManagementClusterName adds the nodepoolFullNameManagementClusterName to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithNodepoolFullNameManagementClusterName(nodepoolFullNameManagementClusterName string) *NodepoolApplyHelperApplyParams {
	o.SetNodepoolFullNameManagementClusterName(nodepoolFullNameManagementClusterName)
	return o
}

// SetNodepoolFullNameManagementClusterName adds the nodepoolFullNameManagementClusterName to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetNodepoolFullNameManagementClusterName(nodepoolFullNameManagementClusterName string) {
	o.NodepoolFullNameManagementClusterName = nodepoolFullNameManagementClusterName
}

// WithNodepoolFullNameProvisionerName adds the nodepoolFullNameProvisionerName to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithNodepoolFullNameProvisionerName(nodepoolFullNameProvisionerName string) *NodepoolApplyHelperApplyParams {
	o.SetNodepoolFullNameProvisionerName(nodepoolFullNameProvisionerName)
	return o
}

// SetNodepoolFullNameProvisionerName adds the nodepoolFullNameProvisionerName to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetNodepoolFullNameProvisionerName(nodepoolFullNameProvisionerName string) {
	o.NodepoolFullNameProvisionerName = nodepoolFullNameProvisionerName
}

// WithNodepoolFullNameTanzuKubernetesClusterName adds the nodepoolFullNameTanzuKubernetesClusterName to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) WithNodepoolFullNameTanzuKubernetesClusterName(nodepoolFullNameTanzuKubernetesClusterName string) *NodepoolApplyHelperApplyParams {
	o.SetNodepoolFullNameTanzuKubernetesClusterName(nodepoolFullNameTanzuKubernetesClusterName)
	return o
}

// SetNodepoolFullNameTanzuKubernetesClusterName adds the nodepoolFullNameTanzuKubernetesClusterName to the nodepool apply helper apply params
func (o *NodepoolApplyHelperApplyParams) SetNodepoolFullNameTanzuKubernetesClusterName(nodepoolFullNameTanzuKubernetesClusterName string) {
	o.NodepoolFullNameTanzuKubernetesClusterName = nodepoolFullNameTanzuKubernetesClusterName
}

// WriteToRequest writes these params to a swagger request
func (o *NodepoolApplyHelperApplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param nodepool.fullName.managementClusterName
	if err := r.SetPathParam("nodepool.fullName.managementClusterName", o.NodepoolFullNameManagementClusterName); err != nil {
		return err
	}

	// path param nodepool.fullName.provisionerName
	if err := r.SetPathParam("nodepool.fullName.provisionerName", o.NodepoolFullNameProvisionerName); err != nil {
		return err
	}

	// path param nodepool.fullName.tanzuKubernetesClusterName
	if err := r.SetPathParam("nodepool.fullName.tanzuKubernetesClusterName", o.NodepoolFullNameTanzuKubernetesClusterName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
