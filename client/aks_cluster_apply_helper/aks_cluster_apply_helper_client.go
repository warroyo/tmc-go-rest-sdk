// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_apply_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new aks cluster apply helper API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for aks cluster apply helper API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AksClusterApplyHelperApply(params *AksClusterApplyHelperApplyParams, opts ...ClientOption) (*AksClusterApplyHelperApplyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AksClusterApplyHelperApply applies a aks cluster
*/
func (a *Client) AksClusterApplyHelperApply(params *AksClusterApplyHelperApplyParams, opts ...ClientOption) (*AksClusterApplyHelperApplyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAksClusterApplyHelperApplyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AksClusterApplyHelper_Apply",
		Method:             "POST",
		PathPattern:        "/v1alpha1/aksclusters:apply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AksClusterApplyHelperApplyReader{formats: a.formats},
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
	success, ok := result.(*AksClusterApplyHelperApplyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AksClusterApplyHelperApplyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
