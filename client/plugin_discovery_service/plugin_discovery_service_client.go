// Code generated by go-swagger; DO NOT EDIT.

package plugin_discovery_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plugin discovery service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugin discovery service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PluginDiscoveryServiceDescribe(params *PluginDiscoveryServiceDescribeParams, opts ...ClientOption) (*PluginDiscoveryServiceDescribeOK, error)

	PluginDiscoveryServiceList(params *PluginDiscoveryServiceListParams, opts ...ClientOption) (*PluginDiscoveryServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PluginDiscoveryServiceDescribe describes a requested tmc plugin
*/
func (a *Client) PluginDiscoveryServiceDescribe(params *PluginDiscoveryServiceDescribeParams, opts ...ClientOption) (*PluginDiscoveryServiceDescribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginDiscoveryServiceDescribeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PluginDiscoveryService_Describe",
		Method:             "GET",
		PathPattern:        "/v1alpha1/system/binaries/plugins/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PluginDiscoveryServiceDescribeReader{formats: a.formats},
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
	success, ok := result.(*PluginDiscoveryServiceDescribeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PluginDiscoveryServiceDescribeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PluginDiscoveryServiceList lists tmc plugins
*/
func (a *Client) PluginDiscoveryServiceList(params *PluginDiscoveryServiceListParams, opts ...ClientOption) (*PluginDiscoveryServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginDiscoveryServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PluginDiscoveryService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/system/binaries/plugins",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PluginDiscoveryServiceListReader{formats: a.formats},
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
	success, ok := result.(*PluginDiscoveryServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PluginDiscoveryServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}