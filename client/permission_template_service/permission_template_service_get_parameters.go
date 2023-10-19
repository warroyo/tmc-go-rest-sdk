// Code generated by go-swagger; DO NOT EDIT.

package permission_template_service

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

// NewPermissionTemplateServiceGetParams creates a new PermissionTemplateServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPermissionTemplateServiceGetParams() *PermissionTemplateServiceGetParams {
	return &PermissionTemplateServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPermissionTemplateServiceGetParamsWithTimeout creates a new PermissionTemplateServiceGetParams object
// with the ability to set a timeout on a request.
func NewPermissionTemplateServiceGetParamsWithTimeout(timeout time.Duration) *PermissionTemplateServiceGetParams {
	return &PermissionTemplateServiceGetParams{
		timeout: timeout,
	}
}

// NewPermissionTemplateServiceGetParamsWithContext creates a new PermissionTemplateServiceGetParams object
// with the ability to set a context for a request.
func NewPermissionTemplateServiceGetParamsWithContext(ctx context.Context) *PermissionTemplateServiceGetParams {
	return &PermissionTemplateServiceGetParams{
		Context: ctx,
	}
}

// NewPermissionTemplateServiceGetParamsWithHTTPClient creates a new PermissionTemplateServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPermissionTemplateServiceGetParamsWithHTTPClient(client *http.Client) *PermissionTemplateServiceGetParams {
	return &PermissionTemplateServiceGetParams{
		HTTPClient: client,
	}
}

/*
PermissionTemplateServiceGetParams contains all the parameters to send to the API endpoint

	for the permission template service get operation.

	Typically these are written to a http.Request.
*/
type PermissionTemplateServiceGetParams struct {

	/* FullNameManagementClusterName.

	   Name of the ManagementCluster.
	*/
	FullNameManagementClusterName *string

	/* FullNameName.

	   Name of this credential.
	*/
	FullNameName string

	/* FullNameOrgID.

	   ID of Organization.
	*/
	FullNameOrgID *string

	/* FullNameProvisionerName.

	   Name of the Provisioner.
	*/
	FullNameProvisionerName *string

	/* Provider.

	    The infrastructure provider that the permission template is for.

	- PROVIDER_UNSPECIFIED: Unspecified credential provider (default).
	- AWS_EC2: AmazonWeb Services EC2.
	- GENERIC_S3: Generic S3 provider.
	- AZURE_AD: Azure Active Directory.
	- AWS_EKS: AWS EKS.
	- AZURE_AKS: Azure AKS.
	- GENERIC_KEY_VALUE: Generic key value pair.

	    Default: "PROVIDER_UNSPECIFIED"
	*/
	Provider *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the permission template service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermissionTemplateServiceGetParams) WithDefaults() *PermissionTemplateServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the permission template service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermissionTemplateServiceGetParams) SetDefaults() {
	var (
		providerDefault = string("PROVIDER_UNSPECIFIED")
	)

	val := PermissionTemplateServiceGetParams{
		Provider: &providerDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithTimeout(timeout time.Duration) *PermissionTemplateServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithContext(ctx context.Context) *PermissionTemplateServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithHTTPClient(client *http.Client) *PermissionTemplateServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullNameManagementClusterName adds the fullNameManagementClusterName to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithFullNameManagementClusterName(fullNameManagementClusterName *string) *PermissionTemplateServiceGetParams {
	o.SetFullNameManagementClusterName(fullNameManagementClusterName)
	return o
}

// SetFullNameManagementClusterName adds the fullNameManagementClusterName to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetFullNameManagementClusterName(fullNameManagementClusterName *string) {
	o.FullNameManagementClusterName = fullNameManagementClusterName
}

// WithFullNameName adds the fullNameName to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithFullNameName(fullNameName string) *PermissionTemplateServiceGetParams {
	o.SetFullNameName(fullNameName)
	return o
}

// SetFullNameName adds the fullNameName to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetFullNameName(fullNameName string) {
	o.FullNameName = fullNameName
}

// WithFullNameOrgID adds the fullNameOrgID to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithFullNameOrgID(fullNameOrgID *string) *PermissionTemplateServiceGetParams {
	o.SetFullNameOrgID(fullNameOrgID)
	return o
}

// SetFullNameOrgID adds the fullNameOrgId to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetFullNameOrgID(fullNameOrgID *string) {
	o.FullNameOrgID = fullNameOrgID
}

// WithFullNameProvisionerName adds the fullNameProvisionerName to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithFullNameProvisionerName(fullNameProvisionerName *string) *PermissionTemplateServiceGetParams {
	o.SetFullNameProvisionerName(fullNameProvisionerName)
	return o
}

// SetFullNameProvisionerName adds the fullNameProvisionerName to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetFullNameProvisionerName(fullNameProvisionerName *string) {
	o.FullNameProvisionerName = fullNameProvisionerName
}

// WithProvider adds the provider to the permission template service get params
func (o *PermissionTemplateServiceGetParams) WithProvider(provider *string) *PermissionTemplateServiceGetParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the permission template service get params
func (o *PermissionTemplateServiceGetParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *PermissionTemplateServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FullNameManagementClusterName != nil {

		// query param fullName.managementClusterName
		var qrFullNameManagementClusterName string

		if o.FullNameManagementClusterName != nil {
			qrFullNameManagementClusterName = *o.FullNameManagementClusterName
		}
		qFullNameManagementClusterName := qrFullNameManagementClusterName
		if qFullNameManagementClusterName != "" {

			if err := r.SetQueryParam("fullName.managementClusterName", qFullNameManagementClusterName); err != nil {
				return err
			}
		}
	}

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

	if o.FullNameProvisionerName != nil {

		// query param fullName.provisionerName
		var qrFullNameProvisionerName string

		if o.FullNameProvisionerName != nil {
			qrFullNameProvisionerName = *o.FullNameProvisionerName
		}
		qFullNameProvisionerName := qrFullNameProvisionerName
		if qFullNameProvisionerName != "" {

			if err := r.SetQueryParam("fullName.provisionerName", qFullNameProvisionerName); err != nil {
				return err
			}
		}
	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}