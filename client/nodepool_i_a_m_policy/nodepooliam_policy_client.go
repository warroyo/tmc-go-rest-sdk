// Code generated by go-swagger; DO NOT EDIT.

package nodepool_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nodepool i a m policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodepool i a m policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NodepoolIAMPolicyGet(params *NodepoolIAMPolicyGetParams, opts ...ClientOption) (*NodepoolIAMPolicyGetOK, error)

	NodepoolIAMPolicyPatch(params *NodepoolIAMPolicyPatchParams, opts ...ClientOption) (*NodepoolIAMPolicyPatchOK, error)

	NodepoolIAMPolicyTestPermissions(params *NodepoolIAMPolicyTestPermissionsParams, opts ...ClientOption) (*NodepoolIAMPolicyTestPermissionsOK, error)

	NodepoolIAMPolicyUpdate(params *NodepoolIAMPolicyUpdateParams, opts ...ClientOption) (*NodepoolIAMPolicyUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
NodepoolIAMPolicyGet gets policy for a nodepool
*/
func (a *Client) NodepoolIAMPolicyGet(params *NodepoolIAMPolicyGetParams, opts ...ClientOption) (*NodepoolIAMPolicyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodepoolIAMPolicyGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodepoolIAMPolicy_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodepoolIAMPolicyGetReader{formats: a.formats},
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
	success, ok := result.(*NodepoolIAMPolicyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodepoolIAMPolicyGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NodepoolIAMPolicyPatch patches a nodepool policy
*/
func (a *Client) NodepoolIAMPolicyPatch(params *NodepoolIAMPolicyPatchParams, opts ...ClientOption) (*NodepoolIAMPolicyPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodepoolIAMPolicyPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodepoolIAMPolicy_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodepoolIAMPolicyPatchReader{formats: a.formats},
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
	success, ok := result.(*NodepoolIAMPolicyPatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodepoolIAMPolicyPatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NodepoolIAMPolicyTestPermissions tests permissions for a nodepool
*/
func (a *Client) NodepoolIAMPolicyTestPermissions(params *NodepoolIAMPolicyTestPermissionsParams, opts ...ClientOption) (*NodepoolIAMPolicyTestPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodepoolIAMPolicyTestPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodepoolIAMPolicy_TestPermissions",
		Method:             "POST",
		PathPattern:        "/v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}/testPermissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodepoolIAMPolicyTestPermissionsReader{formats: a.formats},
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
	success, ok := result.(*NodepoolIAMPolicyTestPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodepoolIAMPolicyTestPermissionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NodepoolIAMPolicyUpdate updates overwrite policy for a nodepool deleted if body is empty
*/
func (a *Client) NodepoolIAMPolicyUpdate(params *NodepoolIAMPolicyUpdateParams, opts ...ClientOption) (*NodepoolIAMPolicyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodepoolIAMPolicyUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodepoolIAMPolicy_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodepoolIAMPolicyUpdateReader{formats: a.formats},
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
	success, ok := result.(*NodepoolIAMPolicyUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodepoolIAMPolicyUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
