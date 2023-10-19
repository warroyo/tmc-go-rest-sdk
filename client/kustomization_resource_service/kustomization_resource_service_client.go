// Code generated by go-swagger; DO NOT EDIT.

package kustomization_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new kustomization resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kustomization resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	KustomizationResourceServiceCreate(params *KustomizationResourceServiceCreateParams, opts ...ClientOption) (*KustomizationResourceServiceCreateOK, error)

	KustomizationResourceServiceDelete(params *KustomizationResourceServiceDeleteParams, opts ...ClientOption) (*KustomizationResourceServiceDeleteOK, error)

	KustomizationResourceServiceGet(params *KustomizationResourceServiceGetParams, opts ...ClientOption) (*KustomizationResourceServiceGetOK, error)

	KustomizationResourceServiceList(params *KustomizationResourceServiceListParams, opts ...ClientOption) (*KustomizationResourceServiceListOK, error)

	KustomizationResourceServiceUpdate(params *KustomizationResourceServiceUpdateParams, opts ...ClientOption) (*KustomizationResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
KustomizationResourceServiceCreate creates a kustomization
*/
func (a *Client) KustomizationResourceServiceCreate(params *KustomizationResourceServiceCreateParams, opts ...ClientOption) (*KustomizationResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKustomizationResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KustomizationResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/clusters/{kustomization.fullName.clusterName}/namespaces/{kustomization.fullName.namespaceName}/fluxcd/kustomizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KustomizationResourceServiceCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KustomizationResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KustomizationResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KustomizationResourceServiceDelete deletes a kustomization
*/
func (a *Client) KustomizationResourceServiceDelete(params *KustomizationResourceServiceDeleteParams, opts ...ClientOption) (*KustomizationResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKustomizationResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KustomizationResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/fluxcd/kustomizations/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KustomizationResourceServiceDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KustomizationResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KustomizationResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KustomizationResourceServiceGet gets a kustomization
*/
func (a *Client) KustomizationResourceServiceGet(params *KustomizationResourceServiceGetParams, opts ...ClientOption) (*KustomizationResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKustomizationResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KustomizationResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/fluxcd/kustomizations/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KustomizationResourceServiceGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KustomizationResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KustomizationResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KustomizationResourceServiceList lists kustomizations
*/
func (a *Client) KustomizationResourceServiceList(params *KustomizationResourceServiceListParams, opts ...ClientOption) (*KustomizationResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKustomizationResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KustomizationResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{searchScope.clusterName}/namespaces/{searchScope.namespaceName}/fluxcd/kustomizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KustomizationResourceServiceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KustomizationResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KustomizationResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KustomizationResourceServiceUpdate updates overwrite a kustomization
*/
func (a *Client) KustomizationResourceServiceUpdate(params *KustomizationResourceServiceUpdateParams, opts ...ClientOption) (*KustomizationResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKustomizationResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KustomizationResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/clusters/{kustomization.fullName.clusterName}/namespaces/{kustomization.fullName.namespaceName}/fluxcd/kustomizations/{kustomization.fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KustomizationResourceServiceUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KustomizationResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KustomizationResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
