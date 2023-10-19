// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_reregister_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new management cluster reregister helper API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management cluster reregister helper API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ManagementClusterReregisterHelperReregister(params *ManagementClusterReregisterHelperReregisterParams, opts ...ClientOption) (*ManagementClusterReregisterHelperReregisterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ManagementClusterReregisterHelperReregister reregisters a management cluster
*/
func (a *Client) ManagementClusterReregisterHelperReregister(params *ManagementClusterReregisterHelperReregisterParams, opts ...ClientOption) (*ManagementClusterReregisterHelperReregisterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewManagementClusterReregisterHelperReregisterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ManagementClusterReregisterHelper_Reregister",
		Method:             "POST",
		PathPattern:        "/v1alpha1/managementclusters:reregister/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ManagementClusterReregisterHelperReregisterReader{formats: a.formats},
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
	success, ok := result.(*ManagementClusterReregisterHelperReregisterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ManagementClusterReregisterHelperReregisterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
