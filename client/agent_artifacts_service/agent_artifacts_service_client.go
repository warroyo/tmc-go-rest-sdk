// Code generated by go-swagger; DO NOT EDIT.

package agent_artifacts_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new agent artifacts service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agent artifacts service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AgentArtifactsServiceList(params *AgentArtifactsServiceListParams, opts ...ClientOption) (*AgentArtifactsServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AgentArtifactsServiceList thes RPC to list all the image artifacts supported by t m c
*/
func (a *Client) AgentArtifactsServiceList(params *AgentArtifactsServiceListParams, opts ...ClientOption) (*AgentArtifactsServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAgentArtifactsServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AgentArtifactsService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/agentartifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AgentArtifactsServiceListReader{formats: a.formats},
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
	success, ok := result.(*AgentArtifactsServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AgentArtifactsServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
