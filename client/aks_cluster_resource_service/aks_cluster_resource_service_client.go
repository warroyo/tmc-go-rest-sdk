// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new aks cluster resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for aks cluster resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AksClusterResourceServiceCreate(params *AksClusterResourceServiceCreateParams, opts ...ClientOption) (*AksClusterResourceServiceCreateOK, error)

	AksClusterResourceServiceDelete(params *AksClusterResourceServiceDeleteParams, opts ...ClientOption) (*AksClusterResourceServiceDeleteOK, error)

	AksClusterResourceServiceGet(params *AksClusterResourceServiceGetParams, opts ...ClientOption) (*AksClusterResourceServiceGetOK, error)

	AksClusterResourceServiceList(params *AksClusterResourceServiceListParams, opts ...ClientOption) (*AksClusterResourceServiceListOK, error)

	AksClusterResourceServicePatch(params *AksClusterResourceServicePatchParams, opts ...ClientOption) (*AksClusterResourceServicePatchOK, error)

	AksClusterResourceServiceUpdate(params *AksClusterResourceServiceUpdateParams, opts ...ClientOption) (*AksClusterResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AksClusterResourceServiceCreate creates an aks cluster
*/
func (a *Client) AksClusterResourceServiceCreate(params *AksClusterResourceServiceCreateParams, opts ...ClientOption) (*AksClusterResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/aksclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*AksClusterResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AksClusterResourceServiceDelete deletes an aks cluster
*/
func (a *Client) AksClusterResourceServiceDelete(params *AksClusterResourceServiceDeleteParams, opts ...ClientOption) (*AksClusterResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/aksclusters/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*AksClusterResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AksClusterResourceServiceGet gets an aks cluster
*/
func (a *Client) AksClusterResourceServiceGet(params *AksClusterResourceServiceGetParams, opts ...ClientOption) (*AksClusterResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/aksclusters/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*AksClusterResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AksClusterResourceServiceList lists aks clusters
*/
func (a *Client) AksClusterResourceServiceList(params *AksClusterResourceServiceListParams, opts ...ClientOption) (*AksClusterResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/aksclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*AksClusterResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AksClusterResourceServicePatch patches partially update an aks cluster
*/
func (a *Client) AksClusterResourceServicePatch(params *AksClusterResourceServicePatchParams, opts ...ClientOption) (*AksClusterResourceServicePatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterResourceServicePatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterResourceService_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1alpha1/aksclusters/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterResourceServicePatchReader{formats: a.formats},
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
	success, ok := result.(*AksClusterResourceServicePatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterResourceServicePatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AksClusterResourceServiceUpdate updates overwrite an aks cluster
*/
func (a *Client) AksClusterResourceServiceUpdate(params *AksClusterResourceServiceUpdateParams, opts ...ClientOption) (*AksClusterResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/aksclusters/{aksCluster.fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterResourceServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*AksClusterResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}