// Code generated by go-swagger; DO NOT EDIT.

package usage_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new usage service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for usage service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UsageServiceGetClusterLevelUsage(params *UsageServiceGetClusterLevelUsageParams, opts ...ClientOption) (*UsageServiceGetClusterLevelUsageOK, error)

	UsageServiceGetOrgLevelUsage(params *UsageServiceGetOrgLevelUsageParams, opts ...ClientOption) (*UsageServiceGetOrgLevelUsageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
UsageServiceGetClusterLevelUsage methods to get the usage info at cluster level
*/
func (a *Client) UsageServiceGetClusterLevelUsage(params *UsageServiceGetClusterLevelUsageParams, opts ...ClientOption) (*UsageServiceGetClusterLevelUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsageServiceGetClusterLevelUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UsageService_GetClusterLevelUsage",
		Method:             "GET",
		PathPattern:        "/v1alpha1/metering/usage/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UsageServiceGetClusterLevelUsageReader{formats: a.formats},
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
	success, ok := result.(*UsageServiceGetClusterLevelUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsageServiceGetClusterLevelUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UsageServiceGetOrgLevelUsage methods to get the usage info at org level
*/
func (a *Client) UsageServiceGetOrgLevelUsage(params *UsageServiceGetOrgLevelUsageParams, opts ...ClientOption) (*UsageServiceGetOrgLevelUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsageServiceGetOrgLevelUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UsageService_GetOrgLevelUsage",
		Method:             "GET",
		PathPattern:        "/v1alpha1/metering/usage/org",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UsageServiceGetOrgLevelUsageReader{formats: a.formats},
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
	success, ok := result.(*UsageServiceGetOrgLevelUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UsageServiceGetOrgLevelUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}