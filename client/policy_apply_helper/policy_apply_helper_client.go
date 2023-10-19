// Code generated by go-swagger; DO NOT EDIT.

package policy_apply_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new policy apply helper API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policy apply helper API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PolicyApplyHelperApply(params *PolicyApplyHelperApplyParams, opts ...ClientOption) (*PolicyApplyHelperApplyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PolicyApplyHelperApply applies a policy
*/
func (a *Client) PolicyApplyHelperApply(params *PolicyApplyHelperApplyParams, opts ...ClientOption) (*PolicyApplyHelperApplyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPolicyApplyHelperApplyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PolicyApplyHelper_Apply",
		Method:             "POST",
		PathPattern:        "/v1alpha1/workspaces/{policy.fullName.workspaceName}/policies:apply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PolicyApplyHelperApplyReader{formats: a.formats},
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
	success, ok := result.(*PolicyApplyHelperApplyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PolicyApplyHelperApplyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
