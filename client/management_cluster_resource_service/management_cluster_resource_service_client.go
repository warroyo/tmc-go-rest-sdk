// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new management cluster resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management cluster resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ManagementClusterResourceServiceCreate(params *ManagementClusterResourceServiceCreateParams, opts ...ClientOption) (*ManagementClusterResourceServiceCreateOK, error)

	ManagementClusterResourceServiceDelete(params *ManagementClusterResourceServiceDeleteParams, opts ...ClientOption) (*ManagementClusterResourceServiceDeleteOK, error)

	ManagementClusterResourceServiceGet(params *ManagementClusterResourceServiceGetParams, opts ...ClientOption) (*ManagementClusterResourceServiceGetOK, error)

	ManagementClusterResourceServiceList(params *ManagementClusterResourceServiceListParams, opts ...ClientOption) (*ManagementClusterResourceServiceListOK, error)

	ManagementClusterResourceServiceUpdate(params *ManagementClusterResourceServiceUpdateParams, opts ...ClientOption) (*ManagementClusterResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ManagementClusterResourceServiceCreate creates a management cluster
*/
func (a *Client) ManagementClusterResourceServiceCreate(params *ManagementClusterResourceServiceCreateParams, opts ...ClientOption) (*ManagementClusterResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagementClusterResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ManagementClusterResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/managementclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagementClusterResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*ManagementClusterResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ManagementClusterResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ManagementClusterResourceServiceDelete deletes a management cluster
*/
func (a *Client) ManagementClusterResourceServiceDelete(params *ManagementClusterResourceServiceDeleteParams, opts ...ClientOption) (*ManagementClusterResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagementClusterResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ManagementClusterResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/managementclusters/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagementClusterResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*ManagementClusterResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ManagementClusterResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ManagementClusterResourceServiceGet gets a management cluster
*/
func (a *Client) ManagementClusterResourceServiceGet(params *ManagementClusterResourceServiceGetParams, opts ...ClientOption) (*ManagementClusterResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagementClusterResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ManagementClusterResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/managementclusters/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagementClusterResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*ManagementClusterResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ManagementClusterResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ManagementClusterResourceServiceList lists management clusters
*/
func (a *Client) ManagementClusterResourceServiceList(params *ManagementClusterResourceServiceListParams, opts ...ClientOption) (*ManagementClusterResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagementClusterResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ManagementClusterResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/managementclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagementClusterResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*ManagementClusterResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ManagementClusterResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ManagementClusterResourceServiceUpdate updates overwrite a management cluster
*/
func (a *Client) ManagementClusterResourceServiceUpdate(params *ManagementClusterResourceServiceUpdateParams, opts ...ClientOption) (*ManagementClusterResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagementClusterResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ManagementClusterResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/managementclusters/{managementCluster.fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagementClusterResourceServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*ManagementClusterResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ManagementClusterResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
