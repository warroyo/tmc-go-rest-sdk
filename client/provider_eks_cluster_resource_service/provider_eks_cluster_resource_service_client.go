// Code generated by go-swagger; DO NOT EDIT.

package provider_eks_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new provider eks cluster resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for provider eks cluster resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ProviderEksClusterResourceServiceGet(params *ProviderEksClusterResourceServiceGetParams, opts ...ClientOption) (*ProviderEksClusterResourceServiceGetOK, error)

	ProviderEksClusterResourceServiceList(params *ProviderEksClusterResourceServiceListParams, opts ...ClientOption) (*ProviderEksClusterResourceServiceListOK, error)

	ProviderEksClusterResourceServiceUpdate(params *ProviderEksClusterResourceServiceUpdateParams, opts ...ClientOption) (*ProviderEksClusterResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ProviderEksClusterResourceServiceGet gets a provider eks cluster
*/
func (a *Client) ProviderEksClusterResourceServiceGet(params *ProviderEksClusterResourceServiceGetParams, opts ...ClientOption) (*ProviderEksClusterResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProviderEksClusterResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProviderEksClusterResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/manage/providereksclusters/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProviderEksClusterResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*ProviderEksClusterResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProviderEksClusterResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProviderEksClusterResourceServiceList lists provider eks clusters
*/
func (a *Client) ProviderEksClusterResourceServiceList(params *ProviderEksClusterResourceServiceListParams, opts ...ClientOption) (*ProviderEksClusterResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProviderEksClusterResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProviderEksClusterResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/manage/providereksclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProviderEksClusterResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*ProviderEksClusterResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProviderEksClusterResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProviderEksClusterResourceServiceUpdate updates overwrite a provider eks cluster
*/
func (a *Client) ProviderEksClusterResourceServiceUpdate(params *ProviderEksClusterResourceServiceUpdateParams, opts ...ClientOption) (*ProviderEksClusterResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProviderEksClusterResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProviderEksClusterResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/manage/providereksclusters/{providerEksCluster.fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProviderEksClusterResourceServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*ProviderEksClusterResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProviderEksClusterResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}