// Code generated by go-swagger; DO NOT EDIT.

package traceflow_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new traceflow resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for traceflow resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TraceflowResourceServiceCreate(params *TraceflowResourceServiceCreateParams, opts ...ClientOption) (*TraceflowResourceServiceCreateOK, error)

	TraceflowResourceServiceDelete(params *TraceflowResourceServiceDeleteParams, opts ...ClientOption) (*TraceflowResourceServiceDeleteOK, error)

	TraceflowResourceServiceGet(params *TraceflowResourceServiceGetParams, opts ...ClientOption) (*TraceflowResourceServiceGetOK, error)

	TraceflowResourceServiceList(params *TraceflowResourceServiceListParams, opts ...ClientOption) (*TraceflowResourceServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
TraceflowResourceServiceCreate creates a traceflow
*/
func (a *Client) TraceflowResourceServiceCreate(params *TraceflowResourceServiceCreateParams, opts ...ClientOption) (*TraceflowResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTraceflowResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TraceflowResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/clusters/{traceflow.fullName.clusterName}/network/antrea/traceflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TraceflowResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*TraceflowResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TraceflowResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TraceflowResourceServiceDelete deletes a traceflow
*/
func (a *Client) TraceflowResourceServiceDelete(params *TraceflowResourceServiceDeleteParams, opts ...ClientOption) (*TraceflowResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTraceflowResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TraceflowResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/network/antrea/traceflows/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TraceflowResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*TraceflowResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TraceflowResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TraceflowResourceServiceGet gets a traceflow
*/
func (a *Client) TraceflowResourceServiceGet(params *TraceflowResourceServiceGetParams, opts ...ClientOption) (*TraceflowResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTraceflowResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TraceflowResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/network/antrea/traceflows/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TraceflowResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*TraceflowResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TraceflowResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TraceflowResourceServiceList lists traceflows
*/
func (a *Client) TraceflowResourceServiceList(params *TraceflowResourceServiceListParams, opts ...ClientOption) (*TraceflowResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTraceflowResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TraceflowResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{searchScope.clusterName}/network/antrea/traceflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TraceflowResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*TraceflowResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TraceflowResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
